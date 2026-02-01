# Normalization Analysis (3NF/BCNF)

> **Date:** 2026-02-01
> **Subject:** Database Schema Compliance Check
> **Standard:** 3NF (Third Normal Form) & BCNF (Boyce-Codd Normal Form)

## 1. Executive Summary

The schema is largely **3NF/BCNF compliant**, with specific, intentional deviations for **Performance (Caching)** and **Architecture (CQRS)**.

- **Compliant:** Core Identity, Tier, Subscription, Admin Action.
- **Intentional Denormalization:** Membership (LTV), Wallet (Balance Projections).
- **Architectural Pattern:** Ledger relies on append-only immutability rather than update-in-place normalization.

---

## 2. Detailed Entity Analysis

### 2.1 Identity & Profile (Compliant)

| Entity | Form | Analysis |
| :--- | :--- | :--- |
| **ACCOUNT** | **BCNF** | `account_id` determines all attributes. Email is a candidate key. No transitive dependencies. |
| **CREATOR_PROFILE** | **BCNF** | `creator_id` determines all attributes. `slug` is a candidate key. `verification_status` is dependent solely on `creator_id`. |

### 2.2 Product & Policy (Compliant)

| Entity | Form | Analysis |
| :--- | :--- | :--- |
| **TIER** | **3NF** | `tier_id` determines price, rank, etc. <br>**Note:** `rank` might have a composite constraint `(creator_id, rank)` but structurally it depends on the PK. |
| **DIGITAL_PRODUCT** | **3NF** | Unique product attributes dependent on `product_id`. |

### 2.3 Relationships (Mixed)

| Entity | Form | Analysis |
| :--- | :--- | :--- |
| **MEMBERSHIP** | **2NF** (Strictly) | `(supporter_id, creator_id)` could show functional dependency. <br>🔴 **Deviation:** `lifetime_value_cents` is an **Aggregate**. Ideally, LTV is `SUM(LedgerEntries WHERE category=earning)`. Storing it here is a **Cache Denormalization**. |
| **SUBSCRIPTION** | **3NF** | `subscription_id` is the key. `current_period_*` fields are relevant to the subscription instance. |
| **ACCESS_ENTITLEMENT**| **3NF** | `entitlement_id` determines access. `source_subscription_id` is a removable dependency (nullable), effectively separating the *grant* from the *source*. |

### 2.4 Financials (CQRS Pattern)

| Entity | Form | Analysis |
| :--- | :--- | :--- |
| **LEDGER_ENTRY** | **BCNF** | Immutable facts. `entry_id` is PK. No updates mean no update anomalies. `policy_version_id` correctly normals out rule attributes. |
| **WALLET** | **Denormalized** | 🔴 **Deviation:** `balance_settled` and `balance_pending` are strictly derived from `SUM(LEDGER_ENTRY)`. <br>**Justification:** This is a **Read Model (Projection)** in a CQRS system. It is NOT the source of truth, but a cached view for performance and locking. |

### 2.5 Governance (Compliant)

| Entity | Form | Analysis |
| :--- | :--- | :--- |
| **ADMIN_ACTION** | **3NF** | `action_id` determines scope and duration. `target_account_id` does not transitively allow determining the action properties. |
| **POLICY_VERSION** | **BCNF** | `policy_version_id` determines the fee structure snapshot. This is a crucial normalization to prevent historical data corruption if fees change. |

---

## 3. Findings & Recommendations

### Finding 1: Membership LTV Cache
- **Observation:** `MEMBERSHIP.lifetime_value_cents` is a computed column.
- **Risk:** Update anomaly. If a Ledger Entry is added but Membership is not updated, data drifts.
- **Mitigation:** Ensure the "Payment Saga" updates both OR rebuild Membership LTV from Ledger asynchronously.
- **Status:** **Acceptable Deviation** (if acknowledged as cache).

### Finding 2: Wallet Balances
- **Observation:** `WALLET` table is entirely redundant data (can be calculated from Ledger).
- **Risk:** "Invariant Violation" (Wallet != Sum(Ledger)).
- **Mitigation:** The architecture explicitly includes an **Invariant Checker** to monitor this risk.
- **Status:** **Required Architectural Pattern** (CQRS).

### Finding 3: Tier Currency
- **Observation:** `TIER` has `price_currency`.
- **Question:** Can a Creator have Tiers in EUR and USD mixed?
- **Normalization Impact:** If `Creator` forces a single currency, then `price_currency` is transitively dependent on `creator_id` (Violates 3NF).
- **Recommendation:** Check business rule. If Creators are multi-currency capable, 3NF holds. If Creators are single-currency, move `currency` to `CREATOR_PROFILE` (or `CREATOR_SETTINGS`) to enforce 3NF.
- **Current Assumption:** HyperX uses USD as base, but this suggests Tiers *could* vary? **Action:** Clarify currency scope.

## 4. Conclusion
The database schema is robust. The only deviations from strict normalization are well-justified by the "Ledger-First" and strict "Auditability" requirements which necessitate immutable logs and derived read models over simple normalized state.
