# Phase 0 Implementation Plan: Foundation & Truth

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Establish the technical foundation (Go skeletal structure, Docker infrastructure) and the "Source of Truth" database tables defined in the audit.

**Architecture:** Hexagonal Architecture (Ports & Adapters) with a "Ledger-First" data model.
**Tech Stack:** Go 1.24, PostgreSQL 16 (Local), Redis 7 (Local), `golang-migrate`.

---

### Task 1: Project Skeleton & Migration Scripts

**Files:**
- Create: `go.mod`
- Create: `Makefile`
- Create: `scripts/migrate_up.ps1` (Windows)
- Create: `scripts/migrate_down.ps1` (Windows)
- Create: `cmd/server/main.go`
- Create: `internal/common/server/http.go`

**Step 1: Initialize Module**
Initialize `go.mod` (e.g., `github.com/hyperx/backend`).

**Step 2: Migration Scripts**
Create PowerShell scripts to run `migrate`:
- `migrate_up.ps1`: Reads DB connection from env (or default localhost) and runs `migrate -path migrations -database ... up`
- `migrate_down.ps1`: Runs `migrate ... down`
*Assumption:* User has `migrate` CLI installed or we provide instruction to install it.

**Step 3: Server Entrypoint**
Create a minimal HTTP server in `cmd/server/main.go` that connects to `localhost` Postgres/Redis (via config/flags) to verify connectivity.

**Step 4: Commit**
```bash
git add .
git commit -m "chore: init project skeleton and infrastructure"
```

### Task 2: Database Migration System

**Files:**
- Create: `migrations/000001_init_schema.up.sql`
- Create: `migrations/000001_init_schema.down.sql`

**Step 1: Identity & Governance Schema**
Implement tables based on `docs/audit/2026-02-01-entity-definitions.md`:
- `accounts` (UUID PK, email, system_account flag)
- `admin_actions` (UUID PK, action_type ENUM)
- `creator_profiles` (UUID PK, verification status)

**Step 2: Financial Core Schema**
Implement "Truth" tables:
- `policy_versions` (Snapshot of fee rules)
- `ledger_entries` (Immutable log, sequence PK, policy_version FK)

**Step 2a: DB Optimization (Triggers & Constraints)**
Implement the optimizations defined in `docs/audit/2026-02-01-db-logic-optimization.md`:
- **Ledger Immutability:** Trigger `BEFORE UPDATE/DELETE` on `ledger_entries` (Raise Exception).
- **Wallet Sync:** Trigger `AFTER INSERT` on `ledger_entries` -> Update `wallets`.
- **LTV Cache:** Trigger `AFTER INSERT` on `ledger_entries` -> Update `memberships.lifetime_value_cents`.
- **Unique Sub:** `CREATE UNIQUE INDEX idx_one_active_sub` (Partial Index).


**Step 3: Verification**
Run `scripts/migrate_up.ps1`. Connect to local Postgres to verify tables exist.

**Step 4: Commit**
```bash
git add migrations/
git commit -m "feat: initial database schema (identity, sub, ledger)"
```

### Task 3: Domain Types (Go Structs)

**Files:**
- Create: `internal/domain/identity/entity.go`
- Create: `internal/domain/finance/entity.go`

**Step 1: Identity Entities**
Define `Account` and `CreatorProfile` structs with JSON tags and validation methods.

**Step 2: Financial Entities**
Define `LedgerEntry` and `PolicyVersion` structs. Ensure `LedgerEntry` has no setters for immutable fields.

**Step 3: Commit**
```bash
git add internal/
git commit -m "feat: core domain entities"
```

---

## Verification Plan

### Automated Tests
- **Integration Test:** `TestInfrastructureConnection`: Starts the server, hits `/health`, confirms DB/Redis connectivity.
- **Migration Test:** Run `up` and `down` migrations to ensure idempotency.

### Manual Verification
- Run `scripts/migrate_up.ps1` -> Success.
- Connect to DB -> Tables exist.
