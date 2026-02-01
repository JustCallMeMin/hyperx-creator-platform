-- Migration: Rollback Database Optimizations
-- Version: 000002

DROP TRIGGER IF EXISTS trg_accounts_updated_at ON accounts;
DROP TRIGGER IF EXISTS trg_update_ltv ON ledger_entries;
DROP TRIGGER IF EXISTS trg_sync_wallet ON ledger_entries;
DROP TRIGGER IF EXISTS trg_prevent_ledger_delete ON ledger_entries;
DROP TRIGGER IF EXISTS trg_prevent_ledger_update ON ledger_entries;

DROP FUNCTION IF EXISTS update_updated_at_column();
DROP FUNCTION IF EXISTS update_membership_ltv();
DROP FUNCTION IF EXISTS sync_wallet_balance();
DROP FUNCTION IF EXISTS prevent_ledger_modification();
