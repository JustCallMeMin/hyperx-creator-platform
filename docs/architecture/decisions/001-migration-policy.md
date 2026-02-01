# HyperX Migration & Data Policies
**Status:** BINDING
**Applies to:** All Engineers

## 1. Migration Standards
- **Numbering:** Sequential 6-digit integers (e.g., `000001`, `000002`). DO NOT use timestamps (prevents merge conflicts in our specific workflow).
- **Forward-Only:**
  - We generate `down` migrations for local dev convenience ONLY.
  - **Production Rollback = Forbidden.** If a migration is bad, apply a new `fix` migration forward.
  - REASON: Ledger immutability means we cannot "undo" schema changes that might have captured financial facts.
- **Immutability:** Once a migration is merged to `main`, it is FROZEN. Never edit an existing migration file.

## 2. Repository Contract
- **Input/Output:** Repositories MUST accept/return **Domain Entities** (`internal/identity/domain`), NOT Database Models.
- **Mapping:** The Repo layer is responsible for mapping `DB Model <-> Domain Entity`.
- **Tx Scope:** Repositories methods should accept `context.Context` and be transaction-agnostic (the Usecase layer manages Tx boundaries via the `UnitOfWork` pattern if needed).
