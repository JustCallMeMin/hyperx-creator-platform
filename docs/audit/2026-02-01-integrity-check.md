# Integrity & Constraints Verification

> **Date:** 2026-02-01
> **Subject:** Business Logic Alignment & Referential Integrity

## 1. Business Rule Alignment

### 1.1 Strict Authority Model
The schema supports the hierarchy of truth:
- **Admin Action (Governance):** `ADMIN_ACTION` table allows overriding business state (e.g., locking payouts despite positive ledger balance). `action_type` enum covers required interventions.
- **Entitlement (Access):** `ACCESS_ENTITLEMENT` is separated from `SUBSCRIPTION`. This correctly allows "Grace Periods" (Subscription: past_due, Entitlement: active) and "Bans" (Subscription: active, Entitlement: revoked).
- **Ledger (Finance):** `LEDGER_ENTRY` is the bottom-turtle truth. `occurred_at` ensures we respect Provider timestamps, not server insertion time.

### 1.1.1 Safety Nets (Invariant Violation)
- **Missing Link Fixed:** Added `INVARIANT_VIOLATION` table definition.
- **Rule:** This table is the "Safe Mode" trigger. Database Architect confirms this table must exist to persist the "State of Emergency" even if the application crashes.


### 1.2 "Ledger-First" & Double-Entry
The schema requires:
- `LEDGER_ENTRY` to have a `category`.
- A check (likely application-side or trigger) that Sum(Inflow) == Sum(Outflow + Fee + Tax).
- **Schema Support:** The structure supports this, but doesn't strictly *enforce* double-entry in the path (e.g., no single `transaction_group_id` forcing paired rows).
- **System Accounts:** To satisfy Double-Entry (Sum=0), we need "Counter-Party" accounts for Platform Fees and Tax.
  - *Mitigation:* `ACCOUNT` table includes `is_system_account` flag. `LEDGER_ENTRY` references these distinct IDs for fee credits.
- **Recommendation:** Ensure the application wraps multiple Ledger Entries in a single DB Transaction with a shared `reference_id`.

### 1.3 Policy Versioning (Immutable History)
- `LEDGER_ENTRY` links to `POLICY_VERSION`.
- **Verified:** This prevents "History Rewriting". If the Platform Fee changes from 10% to 15% tomorrow, historical entries remain linked to the 10% version. Excellent for audit audits.

---

## 2. Referential Integrity (Foreign Keys)

### 2.1 Critical Links
- `CREATOR_PROFILE` -> `ACCOUNT`: **Mandatory 1:1**.
- `SUBSCRIPTION` -> `MEMBERSHIP`: **Mandatory**. Prevents "Orphaned Subscriptions".
- `MEMBERSHIP` -> `ACCOUNT`: **Mandatory**.
- `WALLET` -> `ACCOUNT`: **Mandatory 1:1**.

### 2.2 Potential Weak Points
- `ACCESS_ENTITLEMENT.source_subscription_id` is **Nullable**.
  - **Reasoning:** Valid. Examples: Admin Grants, Gifts, One-time Purchases.
  - **Risk:** "Zombie Entitlements" with no clear source.
  - **Mitigation:** Application logic must enforce `granted_via` corresponds to a valid source ID (e.g., if `granted_via=subscription`, FK must be present). The DB cannot enforce this conditional FK easily.

---

## 3. Final Verdict
The schema supports the high-integrity requirements of the HyperX platform. It prioritizes **Auditability** (Ledger, Policy Versions) and **Governance** (Admin Action) over simplistic CRUD convenience. The identified structural choices regarding Denormalization (Wallet, LTV) are standard and safe given the architecture's CQRS pattern.
