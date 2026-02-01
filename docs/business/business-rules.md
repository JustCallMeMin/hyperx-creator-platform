# Business Rules & Domain Logic: HyperX (Patreon-like Platform)
**Version:** 1.1.0
**Author:** Business Analyst Agent
**Status:** Approved (Principal Engineer Hardened)

## 1. Glossary (Ubiquitous Language)
- **Account:** The central identity and business container.
- **Creator Context:** Role/Profile for individuals earning money by sharing content.
- **Supporter Context:** Role/Context for individuals spending money to support creators.
- **Tier (Policy):** A definition of pricing, billing intervals, and associated benefits.
- **Membership:** The long-term relational link between a Supporter and a Creator.
- **Subscription:** The active financial contract tied to a specific Tier and payment provider.
- **Entitlement:** The actual permission granted to access a resource (derived from Subscription or Admin Action).
- **Ledger:** The immutable source of truth for all financial movements.
- **Dispute:** A formal conflict over a transaction (Chargeback/Refund).
- **Payout:** The process of moving "Settled" earnings from the platform to the Creator's bank/PayPal.
- **Policy Version:** A snapshot of platform rules (fees, taxes) active at a specific time.
- **Digital Product:** One-time purchase item (e.g., file, video) as opposed to recurring Subscriptions.
- **Tax Payable:** A liability entry representing taxes collected on behalf of authorities.
- **Verification Status:** The legal/KYC state of a Creator (`unverified`, `pending`, `verified`, `rejected`).
- **Escrow/Hold:** A state where funds are received but not yet eligible for Payout.

## 2. Domain Authority & Precedence
### 2.1 Authority Resolver (Central Enforcement)
All access, payout, and enforcement checks MUST use the central **Authority Resolver** component. This is the single implementation of the Precedence Order to prevent logic divergence.
- **Evaluation Order:** `AdminAction` -> `Entitlement` -> `Ledger` -> `Subscription`.

### 2.2 Conflict Resolution Principles
- If Ledger vs. Subscription conflict: **Ledger Wins** (Financial Truth).
- If Subscription vs. Entitlement conflict: **Entitlement Wins** (Access Truth).

## 3. Global Business Principles
- **Multi-Role Accounts:** One Account CAN be both a Creator and a Supporter simultaneously.
- **Context Separation:** "Earning" (Creator) and "Spending" (Supporter) are separate logic contexts.
- **Auditability First:** Every state change must be recorded as an immutable fact (Event).
- **Base Currency Principle:** Internally, the Ledger uses **USD** as the single base currency for all accounting. Payment providers handle conversion at the time of the transaction; the `occurred_at` rate is final.

## 4. Financial Rules (Ledger-First)
### 4.1 Fee Structure
Every payment follows this calculation:
- `GrossAmount`: The total paid by the Supporter.
- `ProcessingFee`: Deducted by the payment provider (e.g., PayPal/Stripe).
- `PlatformFee`: System commission (e.g., 10% of Gross).
- `NetEarning`: `Gross - ProcessingFee - PlatformFee`.

### 4.2 Double-Entry Integrity
- The sum of all Creator Net Earnings + Platform Earnings + Processing Fees + Tax Payable must EQUAL the total Inflow from Supporters.
- **Wallets** are projections of the Ledger. `WalletBalance == Sum(LedgerEntries)`.

## 5. Subscription & Access Logic
### 5.1 Tier Lifecycle
- **Tier Deletion (Deprecation):** Subscriptions remain active until the end of the current billing cycle; new subscriptions are blocked.
- **Rank Ordering:** Enables Upgrades/Downgrades logic.

### 5.2 Entitlement Enforcement
- Access requested MUST be checked against the **Entitlement Record**, NOT the Subscription table.
- **Granting:** Automated via `PaymentConfirmed` or manual via `AdminAction:GrantEntitlement`.
- **Revocation:** Automated via `SubscriptionExpired`/`PaymentFailed` or manual via `AdminAction:revoke_access`.

## 6. Governance & Enforcement
### 6.1 AdminAction Scopes
- **Account-level:** Affects login and all role activities (e.g., `suspend_account`).
- **Creator-level:** Affects all tiers/payouts for a specific profile (e.g., `freeze_payout`).
- **Membership-level:** Affects specific supporter-creator relationship.
- **Entitlement-level:** Affects a single access grant.

### 6.1.1 Temporal Validity (AdminAction)
All AdminActions MUST have a defined temporal scope:
- **`effective_at`**: When the action begins (defaults to "now").
- **`expires_at`**: When the enforcement naturally ends (nullable for indefinite actions).
- **Reversibility:** Reversing an action simply sets a new `expires_at` or `invalidated_at` fact; it does not delete the original action record.

### 6.2 Admin Action Effect Matrix (Contract)
| Action Type | Access Flow | Payout Flow | Login/Session | Billing/Webhooks |
| :--- | :--- | :--- | :--- | :--- |
| `suspend_account` | Blocked (All) | Blocked (All) | Terminated | No Future Intents |
| `freeze_payout` | Allowed | Blocked | Allowed | Allowed |
| `revoke_access` | Blocked (Target) | Allowed | Allowed | Allowed |

### 6.3 AdminAction Reversibility Rules
- **Non-Restorative:** Reversing an action (e.g., `un-suspend`) re-enables the *path* to access/payout but does NOT automatically restore temporal rights (Entitlements) that expired during the suspension period.
- **Manual Intervention:** If an Admin mistakenly revoked access, they must issue a `GrantEntitlement` action after reversing the `revoke_access` to restore original state.

### 6.4 Invariant Safe Mode (Operational Boundaries)
When a **Critical Invariant Violation** is detected, the affected scope enters **Safe Mode**:
- **ALLOWED Actions:**
  - READ-ONLY operations (Viewing balance, reading public content).
  - Admin Actions (Resolving, freezing, auditing).
  - Support ticket creation.
- **BLOCKED Actions:**
  - Payout processing (Hard Block).
  - New Subscription/Digital Product purchases.
  - Creating new Tiers or changing prices.
- **SCOPE:** Safe Mode is restricted to the affected Account/Creator unless the violation is System-wide (Ledger-total mismatch).

### 6.5 Invariant Severity & Resolution
- **Critical (Hard Freeze):** Financial mismatches. Resolution requires a manual **InvariantResolved** event after verification.
- **Warning (Soft Alert):** Minor discrepancies (e.g., Cache out of sync). Automated background reconciliation.

## 7. Operational Constraints
- **Time:** All business logic uses `UTC`. Authority time is `occurred_at` from the provider.
- **Idempotency:** Re-processing a payment event with the same `(provider, provider_event_id)` is forbidden.
- **Replay Safety:** Replaying events MUST NOT trigger side-effects (e.g., no double-charging).

## 8. Dispute & Chargeback Management
- **Immediate Hold:** A dispute notification creates an immediate **Hold** entry in the Ledger.
### 8.1 Retroactive Dispute Semantics
- **Backdating:** Entitlement expiry is backdated to the transaction's `occurred_at`.
- **User Impact of Revocation:**
  - **Notification:** Users receive an immediate "Access Terminated: Payment Dispute" notification.
  - **Tolerated Access:** Historical access logs are kept but future access is immediately denied.
  - **Support Path:** User must resolve the dispute or pay via a new method to restore access.
- **Async Execution:** Dispute processing MUST use an **Async Saga** to prevent blocking webhooks.
- **Loss Resolution:** If dispute is lost, `NetEarning` is reversed. Platform also reverses `PlatformFee` (Shared Risk).

## 9. Payout & Settlement Workflow
- **Settlement Delay (Hold Period):** Earnings move from `Pending` to `Settled` after a configured delay (e.g., 7 days).
- **Eligibility:** `Settled Balance >= Threshold AND balance_settled > 0`.
- **Negative Balance (Debt):** A Creator with a negative balance CAN continue selling; balance recovers from future `NetEarning`. Payouts are blocked by the balance rule.

## 10. Policy Versioning & Retroactive Changes
- **Snapshot Principle:** Every `PaymentConfirmed` records the `PolicyVersionID` active at that moment.
- **Legacy Protection:** Existing subscriptions retain original pricing unless a "Protocol Migration" is authorized.
- **Auditability:** Ledger entries must be reconstructible using ONLY the Event + linked Policy Version.

## 11. Digital Products & Refund Policy
- **Anti-Fraud:** No auto-refunds for digital products. All refunds require manual **AdminAction**.
- **Settlement Window:** Digital product sales have a separate, often longer window (e.g., 14 days) before being `Settled`.

## 12. Onboarding & KYC (Verification)
- **Revenue vs. Payout:** Unverified Creators CAN receive funds but CANNOT request Payouts.
- **Verification Gate:** `CreatorVerificationStatus == verified` is a hard requirement for the Payout module.

## 13. Tax & Regional Compliance
- **Tax-Exclusive Pricing:** VAT/GST is calculated based on **Supporter Region** and added on top of the price.
- **Liability Accounting:** Tax is recorded as a `Credit` to **Tax Payable**. It is NOT platform revenue or creator earning.

---
**Final Doctrine:** The system design prioritizes auditability, recovery, and consistency. It follows a **Strict Authority Model** where the Ledger and Entitlements are the only sources of truth, enforced by a central Authority Resolver and declarative Governance.
