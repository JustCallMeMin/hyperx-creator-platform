# Entity Definitions for Audit

> **Source:** `docs/architecture/architecture-blueprint.md` (ERD) & `docs/business/business-design.md` (Attributes)
> **Purpose:** Baseline for normalization and integrity analysis.

## 1. Identity & Profile Scope

### ACCOUNT
**Definition:** The central identity and business container.
- **PK:** `account_id` (UUID)
- **Attributes:**
  - `email` (Unique, Indexed)
  - `password_hash`
  - `mfa_enabled` (Boolean)
  - `is_system_account` (Boolean) - *For Platform/Tax Authority placeholders*
  - `created_at` (UTC Timestamp)
  - `updated_at` (UTC Timestamp)
- **Relationships:**
  - 1:0..1 -> `CREATOR_PROFILE`
  - 1:N -> `MEMBERSHIP`
  - 1:N -> `LEDGER_ENTRY`
  - 1:N -> `ACCESS_ENTITLEMENT`

### CREATOR_PROFILE
**Definition:** Role/Profile for individuals earning money.
- **PK:** `creator_id` (UUID) - *Typically same as or FK to account_id 1:1*
- **FK:** `account_id`
- **Attributes:**
  - `slug` (Unique, URL-friendly)
  - `display_name`
  - `verification_status` (Enum: unverified, pending, verified, rejected)
  - `short_description`
  - `created_at` (UTC Timestamp)
- **Relationships:**
  - 1:N -> `TIER`
  - 1:N -> `DIGITAL_PRODUCT`

---

## 2. Product & Policy Scope

### TIER
**Definition:** A subscription policy defining price and benefits.
- **PK:** `tier_id` (UUID)
- **FK:** `creator_id`
- **Attributes:**
  - `title`
  - `price_amount_cents` (Integer)
  - `price_currency` (ISO Code, e.g., 'USD')
  - `billing_interval` (Enum: month, year)
  - `status` (Enum: active, deprecated, archived)
  - `rank` (Integer) - *For upgrade/downgrade logic*
  - `created_at`
- **Relationships:**
  - 1:N -> `SUBSCRIPTION`
  - Unique Constraint: `(creator_id, rank)`

### DIGITAL_PRODUCT
**Definition:** One-time purchase item.
- **PK:** `product_id` (UUID)
- **FK:** `creator_id`
- **Attributes:**
  - `title`
  - `price_amount_cents`
  - `price_currency`
  - `is_active` (Boolean)
- **Relationships:**
  - 1:N -> `LEDGER_ENTRY` (Sales)

---

## 3. Relationship & Billing Scope

### MEMBERSHIP
**Definition:** Long-term relationship between Supporter and Creator.
- **PK:** `membership_id` (UUID)
- **FK:** `supporter_account_id`
- **FK:** `creator_id`
- **Attributes:**
  - `status` (Enum: active, suspended, terminated)
  - `created_at`
  - `lifetime_value_cents` (Aggregated)
- **Relationships:**
  - 1:N -> `SUBSCRIPTION` (History of subs)

### SUBSCRIPTION
**Definition:** Active financial commitment.
- **PK:** `subscription_id` (UUID)
- **FK:** `membership_id`
- **FK:** `tier_id`
- **Attributes:**
  - `provider_sub_id` (External Reference)
  - `status` (Enum: active, past_due, cancelled, expired)
  - `current_period_start` (UTC)
  - `current_period_end` (UTC)
  - `payment_provider` (Enum: paypal, stripe)
- **Notes:**
  - *Conflict Rule:* If Ledger says paid but Sub says expired -> Ledger wins for financials, Sub wins for future billing.

### ACCESS_ENTITLEMENT
**Definition:** Permission grant (Source of Access Truth).
- **PK:** `entitlement_id` (UUID)
- **FK:** `account_id` (Recipient)
- **FK:** `resource_id` (Target: Tier or Product)
- **FK:** `source_subscription_id` (Nullable)
- **Attributes:**
  - `status` (Enum: active, revoked, expired)
  - `starts_at` (UTC)
  - `expires_at` (UTC - Crucial for access)
  - `granted_via` (Enum: subscription, purchase, admin_grant, gift)

---

## 4. Financial Scope (Ledger-First)

### LEDGER_ENTRY
**Definition:** Immutable financial fact.
- **PK:** `entry_id` (UUID)
- **seq:** `sequence` (BigSerial, Global Order)
- **FK:** `account_id` (Owner of the wallet affected)
- **Attributes:**
  - `type` (Enum: credit, debit)
  - `amount_minor` (Integer, Positive)
  - `currency` (CHAR(3))
  - `category` (Enum: earning, purchase, platform_fee, processing_fee, tax, payout, adjustment)
  - `reference_id` (FK to Payment/Payout)
  - `occurred_at` (UTC - **Authoritative**)
  - `policy_version_id` (Snapshot of rules applied)
- **Relationships:**
  - N:1 -> `POLICY_VERSION`

### PAYOUT
**Definition:** Mechanism to move settled funds to external world.
- **PK:** `payout_id` (UUID)
- **FK:** `creator_account_id`
- **Attributes:**
  - `amount_minor` (Integer)
  - `currency`
  - `status` (Enum: requested, processing, completed, failed, cancelled)
  - `provider_payout_id` (External Ref)
  - `requested_at` (UTC)
  - `processed_at` (UTC)
- **Integrity Rule:** MUST define a corresponding `LEDGER_ENTRY` (Debit).

### WALLET
**Definition:** Projected balance container (Read Model).
- **PK:** `wallet_id` (UUID)
- **FK:** `account_id`
- **Attributes:**
  - `balance_settled` (Integer)
  - `balance_pending` (Integer)
  - `currency`
  - `last_updated_at`
  - `last_entry_sequence` (Watermark)

### POLICY_VERSION
**Definition:** Snapshot of platform rules at time of transaction.
- **PK:** `policy_version_id` (UUID)
- **Attributes:**
  - `platform_fee_percent` (Decimal)
  - `effective_from` (UTC)
  - `effective_to` (UTC/Null)

---

## 5. Governance Scope

### ADMIN_ACTION
**Definition:** Governance event overriding standard rules.
- **PK:** `action_id` (UUID)
- **FK:** `target_account_id`
- **FK:** `admin_user_id`
- **Attributes:**
  - `action_type` (Enum: suspend_account, freeze_payout, revoke_access, grant_entitlement)
  - `reason` (Text)
  - `effective_at` (UTC)
  - `expires_at` (UTC, Nullable)
  - `invalidated_at` (UTC, Nullable - for Reversal)

### INVARIANT_VIOLATION
**Definition:** Safety record for detected system anomalies.
- **PK:** `violation_id` (UUID)
- **FK:** `account_id` (Scope of violation)
- **Attributes:**
  - `violation_type` (Enum: ledger_mismatch, subscription_entitlement_mismatch)
  - `severity` (Enum: critical, warning)
  - `details_json` (Snapshot of mismatch)
  - `detected_at` (UTC)
  - `resolved_at` (UTC, Nullable)
  - `resolution_note` (Text)
