# Architecture Blueprint & Diagrams: HyperX
**Version:** 1.1.0
**Role:** Senior Solutions Architect
**Status:** Approved (Principal Engineer Hardened)

This document provides a comprehensive visual and structural guide for building, operating, and scaling the HyperX platform.

---

## 1. System Context Diagram
**Purpose:** Defines the boundaries of the system and its interactions with external actors.

```mermaid
graph TD
    Supporter((Supporter)) -->|Subscribes/Pays| HyperX[HyperX Platform]
    Creator((Creator)) -->|Manages Tiers/Content| HyperX
    Admin((Admin)) -->|Governs/Supports| HyperX
    
    HyperX -->|Processes Payments| PayPal[PayPal/Stripe API]
    PayPal -.->|Webhook Notifications| HyperX
    
    subgraph "External Authority (Payment Provider)"
        PayPal
        note1[At-least-once delivery]
        note2[occurred_at is authoritative]
    end
    
    HyperX -->|Sends Notifications| MailEngine[Email/Push Provider]
```

**Notes:** 
- The system is centrally controlled but relies on External Authority for financial facts.
- **Constraint:** `occurred_at` from provider is the only authoritative time for financial logic.

---

## 2. High-Level Architecture (Modular Monolith)
**Purpose:** Shows internal organization and mandatory enforcement points.

```mermaid
graph LR
    subgraph "Application Layer"
        API[REST API / Chi Router]
        Workers[Background Workers]
    end

    subgraph "Modular Core (Hexagonal)"
        Identity[Identity Module]
        Subscription[Subscription Module]
        Wallet[Wallet/Ledger Module]
        Governance[Governance Module]
    end

    subgraph "Shared Infrastructure"
        AuthResolver[Authority Resolver - Shared Lib]
        Outbox[Transactional Outbox]
        Note[❗All decisions MUST pass through Resolver]
    end

    subgraph "Data Services"
        Postgres[(PostgreSQL - Truth)]
        Redis[(Redis - Access Cache)]
    end

    API --> Identity & Subscription & Wallet & Governance
    Workers --> Subscription & Wallet & Governance
    
    Identity & Subscription & Wallet & Governance --> AuthResolver
    Identity & Subscription & Wallet & Governance --> Postgres
    Subscription & Wallet --> Redis
```

---

## 3. Data Model (ERD)
**Purpose:** Defines the schema and relationships, emphasizing temporal truth.

```mermaid
erDiagram
    ACCOUNT ||--o| CREATOR_PROFILE : "has"
    CREATOR_PROFILE ||--o{ TIER : "defines"
    CREATOR_PROFILE ||--o{ DIGITAL_PRODUCT : "sells"
    ACCOUNT ||--o{ MEMBERSHIP : "participates"
    MEMBERSHIP ||--o{ SUBSCRIPTION : "has one active"
    
    ACCOUNT ||--o{ LEDGER_ENTRY : "owns"
    LEDGER_ENTRY }|--|| POLICY_VERSION : "calculated by"
    LEDGER_ENTRY ||--o| WALLET : "updates projection"
    
    ACCOUNT ||--o{ ACCESS_ENTITLEMENT : "granted to"
    CREATOR_PROFILE ||--o{ ACCESS_ENTITLEMENT : "grants for"

    ADMIN_ACTION }|--|| ACCOUNT : "targets"
```

**Design Notes:**
- **Temporal Truth:** All truth tables (Ledger, Entitlements, AdminActions) MUST carry authoritative `occurred_at_utc`.
- **Partitioning:** `ledger_entries` is monthly partitioned. `ledger_seq` (BigSerial) ensures total order.

---

## 4. AdminAction Lifecycle
**Purpose:** Visualizes the journey of a governance event from issuance to expiration/reversal.

```mermaid
stateDiagram-v2
    [*] --> Issued: Admin creates Action
    Issued --> Effective: effective_at reached
    Effective --> Active: Authority Resolver enforces
    Active --> Expired: expires_at reached
    Active --> Reversed: manual invalidated_at set
    Expired --> [*]
    Reversed --> [*]
    
    note right of Active
        Authority Resolver checks 
        Effective <= Now < Expires
    end
```

---

## 5. Sequence: Payment & Entitlement Flow (with Failure Path)
**Purpose:** Explains the journey from webhook to access, including retry logic.

```mermaid
sequenceDiagram
    participant P as PayPal
    participant W as Webhook Handler
    participant S as Payment Saga (Saga State)
    participant L as Ledger (Wallet Module)
    participant E as Entitlement Module
    participant R as Redis Cache

    P->>W: Webhook (Payment.Capture.Completed)
    W->>W: Validate Signature & Idempotency
    W->>S: Start Payment Saga
    
    S->>L: Append Ledger Entry
    L-->>S: Success
    
    alt Grant Success
        S->>E: Create Access Entitlement
        E->>R: Invalidate Cache
        E-->>S: Success
    else Partial Failure (Entitlement Error)
        S->>S: Retry using same internal_command_id
        Note right of S: Idempotency prevents double Ledger writes
    end
    
    S->>S: Mark Saga Completed
```

---

## 6. Access Validation (Fast Path & Fail-Closed)
**Purpose:** Shows security enforcement with performance and safety fallbacks.

```mermaid
sequenceDiagram
    participant U as Supporter
    participant G as API Gateway / Middleware
    participant AR as Authority Resolver
    participant R as Redis Cache
    participant DB as Postgres (Truth)

    U->>G: Request Content (GET /post/123)
    G->>AR: EvaluateAccess(acc_id, cre_id, view_post)
    
    AR->>R: Get Cached Entitlement
    alt Cache Hit
        R-->>AR: active_ranks: [1, 2]
    else Cache Miss / Inval
        AR->>DB: Query Access Entitlements
        DB-->>AR: active_ranks: [1, 2]
        AR->>R: Warm up Cache
    end
    
    AR->>DB: Check Admin Actions (suspend_account?)
    
    alt Authorized
        AR-->>G: Result (Allowed: true)
        G->>U: Content JSON
    else Fail-Closed / Denied
        AR-->>G: Result (Allowed: false, Reason: "infra_failure/suspended")
        G->>U: Deny Access (403/500)
    end
```

---

## 7. Failure & Recovery Flows (Saga & Replay)
**Purpose:** Shows how the system handles state corruption and recovery.

```mermaid
graph TD
    subgraph "Failure Handling"
        WH[Webhook Failure] -->|301/500| PP[Provider Retries]
        SagaL[Saga Step Fail] -->|Exponential Backoff| SagaR[Saga Retry]
    end

    subgraph "Recovery (Cold Replay)"
        Proj[(Projections)] -->|Truncate| Empty[(Empty State)]
        Events[Outbox/Ledger] -->|Replay Loop| Proj
        Note[❗Side-effects: DISABLED]
    end

    SagaR -->|Max Retries Reached| DLQ[Dead Letter / Manual Review]
```

---

## 8. Security Architecture Diagram
**Purpose:** Defines encryption, secrets, and trust zones.

```mermaid
graph TD
    subgraph "Public Internet"
        Client[Web/Mobile Client]
    end

    subgraph "Trust Zone: DMZ"
        LB[Load Balancer - TLS Termination]
    end

    subgraph "Trust Zone: Application"
        App[HyperX API Nodes]
    end

    subgraph "Trust Zone: Data (Hardened)"
        Vault[Managed Secrets Store]
        DB_Master[(Postgres Master - Read/Write)]
        DB_Audit[(Postgres Replica - Read-Only)]
    end

    Client -->|HTTPS| LB
    LB -->|Internal HTTP| App
    App -.->|Fetch Keys| Vault
    App -->|RW| DB_Master
    App -->|RO| DB_Audit
```

---

## 9. Operational & Observability Diagram
**Purpose:** Shows monitoring and the invariant resolution loop.

```mermaid
graph BT
    App[HyperX Nodes] -->|Structured Logs| ELK[ELK / Loki]
    App -->|Metrics| Grafana[Grafana]
    
    DB[(Postgres)] -->|Continuous Check| InvariantWorker[Invariant Checker]
    
    InvariantWorker -->|CRITICAL| SafeMode[Enforce Safe Mode]
    SafeMode -->|Alert| Slack[#alerts-critical]
    Admin -->|Manual Fix| Resolve[InvariantResolved Event]
    Resolve -->|Restore| App
```

---

## 10. Deployment Diagram (Scaling)
**Purpose:** Physical infrastructure layout.

```mermaid
graph TD
    subgraph "Regional VPC"
        subgraph "Public Subnet"
            LB[NGINX / Load Balancer]
        end
        
        subgraph "Private App Subnet"
            Node1[API Node 1]
            Node2[API Node 2]
        end
        
        subgraph "Private Data Subnet"
            PG_M[(Postgres Master)]
            PG_R[(Postgres Replica)]
            Redis[(Redis Cluster)]
        end
    end

    LB --> Node1 & Node2
    Node1 & Node2 --> PG_M
    PG_M -->|Streaming| PG_R
    Node1 & Node2 --> Redis
    Reporting[Reporting Tool] -->|Read-Only| PG_R
```
