# Database Logic Optimization Analysis

> **Context:** Reducing Backend Load via Database Constraints, Triggers, and Generated Columns.
> **Agents:** Database Architect & Business Analyst.

## 1. High-Value Optimization Candidates

### 1.1 Wallet Balance Projection (CQRS in DB)
- **Business Rule:** `WalletBalance = Sum(LedgerEntries)`.
- **Current Load:** Backend usually queries Ledger or manages a separate Wallet update transaction.
- **DB Optimization:**
  - **Mechanism:** `AFTER INSERT` Trigger on `LEDGER_ENTRY`.
  - **Action:** Atomically update `WALLET.balance_settled` (and `pending`).
  - **Benefit:** Guarantees `Wallet` is *always* perfectly in sync with Ledger. Eliminates app-level race conditions.
  - **Cost:** Slower Insert on Ledger (Locking Wallet row). *Acceptable for financial integrity.*

### 1.2 Immutable Ledger Enforcement
- **Business Rule:** Ledger Entries cannot be modified or deleted.
- **DB Optimization:**
  - **Mechanism:** `BEFORE UPDATE/DELETE` Trigger on `LEDGER_ENTRY` that raises `EXCEPTION`.
  - **Benefit:** Absolute security against developer error or SQL injection modifying history.

### 1.3 Subscription Uniqueness (Partial Index)
- **Business Rule:** A supporter can only have **one** `Active` subscription per Creator.
- **DB Optimization:**
  - **Mechanism:** `CREATE UNIQUE INDEX idx_one_active_sub ON subscriptions (membership_id) WHERE status IN ('active', 'past_due');`
  - **Benefit:** Eliminates "Race Condition" where two requests roughly simultaneous could create double subs. Zero backend check needed.

### 1.4 Tier Rank Integrity
- **Business Rule:** Tiers must have unique ranks within a Creator profile.
- **DB Optimization:**
  - **Mechanism:** `UNIQUE (creator_id, rank) DEFERRABLE`?
  - **Constraint:** Reordering ranks (Swap 1 and 2) is hard with strict unique constraints.
  - **Recommendation:** Use Backend logic for reordering, but enforce `UNIQUE` to prevent corruption.

### 1.5 Membership LTV Calculation
- **Business Rule:** `LifetimeValue` is strict sum of earnings.
- **DB Optimization:**
  - **Mechanism:** Trigger on `LEDGER_ENTRY` (similar to Wallet).
  - **Benefit:** Membership list sorting by "Top Supporters" becomes instant (Indexed Column) vs massive aggregation query.

---

## 2. Advanced: Generated Columns & Constraints

### 2.1 Email Case Insensitivity
- **Rule:** Emails are case-insensitive unique.
- **DB Optimization:** `CREATE UNIQUE INDEX idx_account_email ON accounts (LOWER(email));` OR Generated Column `email_normalized AS (LOWER(email)) STORED`.

### 2.2 Status Transitions (Check Constraints)
- **Rule:** `processed_at` must be > `requested_at`.
- **DB Optimization:** `CONSTRAINT chk_timeline CHECK (processed_at >= requested_at)`.

---

## 3. Rejected Candidates (Keep in Backend)

### 3.1 Complex Fee Calculation
- **Reasoning:** Computing "Tenancy Tax + Platform Fee + VAT" involves changing `PolicyVersions`.
- **Decision:** Keep in Go. Logic is versioned and complex. Hard to debug in PL/pgSQL.

### 3.2 Tier Eligibility (Entitlement Check)
- **Reasoning:** "Can User X access Post Y?" involves arbitrary rules (Tiers, One-offs, Gifts).
- **Decision:** Keep in Go (Authority Resolver). The query is simple (`SELECT count(*) FROM entitlements`), logic is dynamic.

---

## 4. Implementation Plan (Phase 0 Updates)

We should include these optimizations in the **Initial Schema (`up.sql`)** to ensure integrity from Day 1.

| Logic | Implementation | Impact |
| :--- | :--- | :--- |
| **Ledger Immutability** | Trigger (Raise Exception) | **Safety Critical** |
| **Wallet Sync** | Trigger (Update Row) | **Perf & Consistency** |
| **Single Active Sub** | Partial Unique Index | **Data Integrity** |
| **LTV Cache** | Trigger (Update Row) | **Read Perf** |
