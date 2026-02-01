# Technical Design: HyperX (Patreon-like Platform)
**Version:** 1.1.0
**Status:** Approved (Principal Engineer Hardened)

## 1. Architectural Patterns
- **Pattern:** Modular Monolith with Hexagonal Architecture.
- **Consistency Model:** Event-Driven with Transactional Outbox.
- **Accounting Pattern:** Event-Sourcing for Ledger; Projections for Wallets.
- **Access Pattern:** Fast Path (Redis Cache) + Authoritative Path (Postgres).

## 2. Component Diagram (Logical)
- `Identity`: Auth (JWT), Account Management.
- `Subscription`: Tier Policy, Membership Management, Billing Engine.
- `Payment`: Provider Integration (PayPal SDK), Webhook Ingestion, Dispute/Payment Sagas.
- `Wallet`: Ledger (Append-only), Balance Projections, Payout Logic, Dispute Logic.
- `Content`: Posts, Benefits, Entitlement Enforcement.
- `Governance`: Admin Actions, Invariant Checking, Policy Versioning, **Authority Resolver (Shared Lib)**.

## 3. Database Schema (Postgres)
### 3.1 Core Tables
- **accounts:** `id, email, password_hash, created_at`
- **creator_profiles:** `id, account_id, page_slug, bio, verification_status, created_at`
- **tiers:** `id, creator_id, name, price_cents, currency, rank, status, created_at`
- **digital_products:** `id, creator_id, name, price_cents, currency, status, created_at`
- **memberships:** `id, account_id, creator_id, status, created_at`
- **subscriptions:** `id, membership_id, tier_id, state, external_ref, expires_at`
- **access_entitlements:** `id, account_id, creator_id, scope, target_id, occurred_at_utc, expires_at` (UNIQUE on active scope)
- **ledger_entries (Partitioned by Month):** `id, ledger_seq (BigSerial), account_id, amount_cents, currency_code (ISO), direction (DR/CR), type, ref_id, policy_version_id, occurred_at_utc, created_at`
- **wallets (Projection):** `id, account_id, balance_pending, balance_settled, last_ledger_seq, currency_code`
- **tax_liabilities (Projection):** `id, region, total_tax_cents, last_ledger_seq`
- **disputes:** `id, external_ref, amount_cents, status, resolution, occurred_at_utc, created_at`
- **payouts:** `id, creator_id, amount_cents, currency_code, status, provider_ref, processed_at`
- **policy_versions:** `id, platform_fee_bps, processing_fee_bps, tax_rules, active_from`
- **admin_actions:** `id, action_type, target_type, target_id, reason, issued_by, effective_at, expires_at`

### 3.2 Reliability & State Tables
- **outbox_events:** `id, event_type, payload (jsonb), status, created_at`
- **idempotency_keys:** `provider_id, external_evt_id, processed_at` (PK: `provider_id, external_evt_id`)
- **saga_states:** `saga_id, saga_type, state (jsonb), current_step, updated_at`

## 4. Cache & Idempotency Strategy
### 4.1 Internal & External Idempotency
- **Boundary:** Idempotency is enforced at Webhook Ingress (`external_ref`) AND internally between modules using an `internal_command_id`.
- **Logic:** Consumers must check the `internal_command_id` before committing side-effects (Ledger, Entitlement).
- **Partial Failure:** Saga orchestrator retries specific failed steps using the same command ID.

### 4.2 Redis Entitlement Cache
- **Model:** Read-through cache.
- **Fail-Closed Policy:** Fallback to DB if Redis fails. If both fail, Access is **DENIED** (Fail-Closed).
- **TTL Safety:** Set to `min(Entitlement.expires_at, 1hr)`.
- **Invalidation:** Triggered by `EntitlementGranted`, `EntitlementRevoked`, and `AdminActionTaken`.

## 5. Event Catalog (Internal Bus)
| Event | Origin | Effect |
| :--- | :--- | :--- |
| **PaymentConfirmed** | Payment | **Financial Only:** Credits Ledger, Updates Wallet Pending balance. |
| **EntitlementGranted** | Sub / Admin | **Access Only:** Grants Entitlement, Invalidates Redis Cache. |
| **DisputeOpened** | Payment | Triggers **Dispute Saga**: Hold Ledger, Revoke Entitlement (Retroactive). |
| **AdminActionTaken** | Gov | Enforces declarative state changes (Declarative Precedence). |
| **InvariantViolated** | Gov | **CRITICAL:** Hard Payout Freeze; **WARN:** Background reconciler notification. |

## 6. Development Requirements
> **❗ AUTHORITY ENFORCEMENT:** All access, payout, and write decisions MUST go through the Authority Resolver. Direct database/status checks for logic enforcement are non-compliant and strictly forbidden.

### 6.1 Authority Resolver (Contract)
The centralized `internal/pkg/authority` package MUST be the only place enforcing the business precedence logic.
```go
// Authority Resolver Contract
package authority

type Result struct {
    Allowed bool
    Reason  string
    Scope   string // e.g., "account", "membership"
}

func EvaluateAccess(ctx context.Context, accountID, creatorID string, action ActionType) (Result, error)
```
- **Inputs:** `accountID`, `creatorID` (optional for global actions), `action` (e.g., `view_post`, `payout_request`).
- **Enforcement:** Used in Middlewares, Handlers, and Background workers.

### 6.2 Safe Mode & Fail-Closed
- **Fail-Closed:** Entitlement checks MUST deny access if Redis and DB are cả hai inaccessible.
- **Operational Freeze:** When `CriticalInvariant` is triggered, the Authority Resolver MUST inject a `Result{Allowed: false, Reason: "system_safe_mode"}` for all Payout and Write actions in the affected scope.

### 6.3 Technical Constraints
- **Transaction Isolation:** Financial writes (Webhook → Ledger) MUST use `REPEATABLE READ` or `SERIALIZABLE` isolation.
- **Replay Determinism:** The system MUST pass "Cold Replay" tests (zero side-effects during state reconstruction).
- **Stack:** Go 1.22+, PostgreSQL 15+ (Month-based Partitioning for Ledger), Redis 7+, `chi` router, `pgx` driver.
- **Monotonic Sequence:** All truth tables (Ledger, Entitlements) must use monotonic sequence IDs (BigSerial) to ensure strict ordering and easy drift detection.

## 7. Operational Recovery & Policy Migration
- **Policy Correction:** `PolicyCorrectionEvent` delta calculation + Ledger adjustment.
- **Replay Playbook:** Truncate projections -> Reset Outbox -> Replay (Side-effects disabled) -> Verify Invariants.

## 8. Operational Playbooks (Stubs)
- **[Safe Mode Handling]:** Procedures for investigating Critical Invariants while Payouts are frozen.
- **[Invariant Resolution]:** Steps for manual verification and committing `InvariantResolved` factual events.
- **[Projection Rebuild]:** Verified procedure for zero-downtime (or minimal impact) projection reconstruction from the Ledger.
