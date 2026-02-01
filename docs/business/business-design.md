# Business Design: Patreon-like Creator Platform (HyperX)

**Date:** 2026-02-01
**Status:** Validated (Business Core)

## 1. Overview & Objective

Build a high-integrity, audit-friendly creator platform (similar to Patreon/Ko-fi) that supports recurring billing, tiered access, and financial-grade accounting. The project serves as a showcase for Junior/Senior-level engineering skills in Go, focusing on architecture over simple CRUD.

## 2. Core Business Doctrine

### 2.1 Identity & Account Model

- **Relationship:** `Identity (Auth) -> Account (Business Container) -> Roles (Contexts)`.
- **Multi-Persona:** A single Account can act as both a **Creator** (earning money) and a **Supporter** (spending money) simultaneously.
- **Rationale:** Prevents "Identity Fragmentation," shares Trust/KYC, and enables network growth effects.

### 2.2 Domain Entities

| Entity                | Purpose                                                                                                 |
| :-------------------- | :------------------------------------------------------------------------------------------------------ |
| **Account**           | The primary holder of identity and legal status.                                                        |
| **Creator Context**   | Profiles and settings for selling content.                                                              |
| **Supporter Context** | Activity and preferences for buying content.                                                            |
| **Tier**              | A **Policy** defining price, billing interval, and a set of Benefits.                                   |
| **Membership**        | A long-term **Relationship** between a Supporter and a Creator (Status: Active, Suspended, Terminated). |
| **Subscription**      | A **Financial Commitment** tied to a specific Tier with a billing lifecycle (PayPal-backed).            |
| **Entitlement**       | A **permission record** granted via Subscription or one-time purchase.                                  |
| **AdminAction**       | **Governance Entity** used to override, freeze, or revoke business states (Action types: `suspend_account`, `freeze_payout`, `revoke_access`). |
| **InvariantViolation** | A registered **Safety Entity** capturing inconsistencies between state and events (e.g., Ledger vs Wallet mismatch). |

## 3. Financial Integrity (Ledger-First)

- **Principle:** No direct `balance` column mutation. All movements are immutable Ledger entries.
- **Triple-Fee Structure (Net-Earning Model):**
  - `Gross Amount` (Supporter pays)
  - `- Processing Fee` (e.g., PayPal fee)
  - `- Platform Fee` (e.g., 10%)
  - `= Net Earning` (Credited to Creator's Ledger).
- **Double-Entry:** Every transaction must have a balancing entry (Inflow to Platform clearing account vs Outflow to Creator + Platform earnings).

## 4. Key Business Rules

### 4.1 Tier Lifecycle

- **Tier Deletion Strategy:** **Option A (Deprecation)**.
  - When a Creator "deletes" a tier, it is marked as `deprecated`.
  - Existing subscribers continue until they cancel or the term ends.
  - No new subscribers allowed.
- **Hierarchy:** Tiers have a `rank/level` to allow "Upgrade/Downgrade" logic.

### 4.2 Membership Access (Access Control)

- **Rule:** Access to Content is determined by **Entitlement**, not directly by Subscription.
- **Path:** `Payment -> Subscription Valid -> Entitlement Granted -> Access Allowed`.
- This decoupling allows for Trials, Gifting, and Manual Overrides without touching the billing logic.

## 5. Workflow: Subscription Fulfillment (Saga)

1. **Supporter Intent:** Select Tier -> Create `SubscriptionIntent` (Status: `Pending`).
2. **Payment Hub:** Redirect to PayPal for user approval.
3. **Authoritative Event:** Receive PayPal Webhook (`PAYMENT.SALE.COMPLETED`).
4. **Verification:** System verifies amount/currency against the Tier policy.
5. **Fulfillment (Async):**
   - Emit `PaymentConfirmed`.
   - Create Ledger Entries.
   - Activate/Update `Subscription` record.
   - Grant `AccessEntitlement`.
   - Unlock Content.

## 6. Governance, Enforcement & Recovery
### 6.1 Admin Authority (Audit Gold)
- **Mandatory Logic:** Admin interventions (suspensions, payout locks) are NOT code hacks or direct DB flags. They are **Business Facts** recorded as `AdminAction` records.
- **Auditability:** Every intervention must have an `issued_by`, `reason`, and a `reversible` flag.
- **Enforcement:** If an `AdminAction:freeze_payout` exists for a Creator, the Wallet module MUST block all payout requests regardless of Ledger balance.

### 6.2 Invariant Violation Handling
- **Invariant Checks:** The system must periodically (or on-event) verify that `Ledger Sum == Wallet Balance` and `Active Subscription == Valid Entitlement`.
- **Conflict Resolution:** If a mismatch is detected, an `InvariantViolated` event is emitted, and a record is created. The system enters a "Safe Mode" for the affected account until resolved.

### 6.3 Operational Metadata (Idempotency & Time)
- **Time Authority:** All timestamps are `UTC`. Business authority follows `occurred_at` (from PayPal/Provider) rather than `received_at`.
- **Idempotency Contract:** Every external event must be processed via a unique `(provider, provider_event_id)` key. Double-processing is a system failure.
- **Replay-Safe Projections:** 
  - Projections (Wallets, Subscription status, Feed) are derived from the event stream.
  - **Contract:** Truncating a projection and replaying the event stream MUST result in the exact same state. Replay must never trigger external side-effects (e.g., re-sending emails).

---

**Final Doctrine:** The system design prioritizes auditability, recovery, and consistency over easy implementation. It follows a **Strict Authority Model** where the Ledger is the source of financial truth and Entitlements are the source of access truth, with declarative Governance (AdminActions) taking precedence.
