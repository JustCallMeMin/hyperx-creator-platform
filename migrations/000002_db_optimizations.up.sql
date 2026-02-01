-- Migration: Database Optimizations (Triggers & Functions)
-- Version: 000002
-- Description: Implements DB-side logic for Ledger Immutability, Wallet Sync, and LTV Cache

-- ============================================================================
-- LEDGER IMMUTABILITY TRIGGER
-- ============================================================================

CREATE OR REPLACE FUNCTION prevent_ledger_modification()
RETURNS TRIGGER
LANGUAGE plpgsql
VOLATILE
AS $$
BEGIN
    RAISE EXCEPTION 'Ledger entries are immutable and cannot be modified or deleted';
END;
$$;

CREATE TRIGGER trg_prevent_ledger_update
    BEFORE UPDATE ON ledger_entries
    FOR EACH ROW
    EXECUTE FUNCTION prevent_ledger_modification();

CREATE TRIGGER trg_prevent_ledger_delete
    BEFORE DELETE ON ledger_entries
    FOR EACH ROW
    EXECUTE FUNCTION prevent_ledger_modification();

-- ============================================================================
-- WALLET SYNC TRIGGER
-- ============================================================================

CREATE OR REPLACE FUNCTION sync_wallet_balance()
RETURNS TRIGGER
LANGUAGE plpgsql
VOLATILE
AS $$
DECLARE
    v_amount INTEGER;
BEGIN
    -- Determine the amount to add/subtract based on entry type
    IF NEW.entry_type = 'credit' THEN
        v_amount := NEW.amount_minor;
    ELSE
        v_amount := -NEW.amount_minor;
    END IF;

    -- Update or insert wallet balance
    INSERT INTO wallets (account_id, balance_settled, balance_pending, currency, last_entry_sequence, last_updated_at)
    VALUES (NEW.account_id, v_amount, 0, NEW.currency, NEW.sequence, NOW())
    ON CONFLICT (account_id) DO UPDATE
    SET 
        balance_settled = wallets.balance_settled + v_amount,
        last_entry_sequence = NEW.sequence,
        last_updated_at = NOW();

    RETURN NEW;
END;
$$;

CREATE TRIGGER trg_sync_wallet
    AFTER INSERT ON ledger_entries
    FOR EACH ROW
    EXECUTE FUNCTION sync_wallet_balance();

-- ============================================================================
-- MEMBERSHIP LTV CACHE TRIGGER
-- ============================================================================

CREATE OR REPLACE FUNCTION update_membership_ltv()
RETURNS TRIGGER
LANGUAGE plpgsql
VOLATILE
AS $$
BEGIN
    -- Only update LTV for earning-type entries
    IF NEW.category = 'earning' AND NEW.entry_type = 'credit' THEN
        UPDATE memberships
        SET lifetime_value_cents = lifetime_value_cents + NEW.amount_minor
        WHERE supporter_account_id = NEW.account_id;
    END IF;

    RETURN NEW;
END;
$$;

CREATE TRIGGER trg_update_ltv
    AFTER INSERT ON ledger_entries
    FOR EACH ROW
    EXECUTE FUNCTION update_membership_ltv();

-- ============================================================================
-- UPDATED_AT TRIGGER (Generic)
-- ============================================================================

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER
LANGUAGE plpgsql
VOLATILE
AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$;

CREATE TRIGGER trg_accounts_updated_at
    BEFORE UPDATE ON accounts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
