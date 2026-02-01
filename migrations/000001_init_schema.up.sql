-- Migration: Initial Schema
-- Version: 000001
-- Description: Core tables for Identity, Governance, Subscription, and Financial domains

-- ============================================================================
-- ENUMS
-- ============================================================================

CREATE TYPE verification_status AS ENUM ('unverified', 'pending', 'verified', 'rejected');
CREATE TYPE admin_action_type AS ENUM ('suspend_account', 'freeze_payout', 'revoke_access', 'grant_entitlement');
CREATE TYPE tier_status AS ENUM ('active', 'deprecated', 'archived');
CREATE TYPE billing_interval AS ENUM ('month', 'year');
CREATE TYPE membership_status AS ENUM ('active', 'suspended', 'terminated');
CREATE TYPE subscription_status AS ENUM ('active', 'past_due', 'cancelled', 'expired');
CREATE TYPE payment_provider AS ENUM ('paypal', 'stripe');
CREATE TYPE entitlement_status AS ENUM ('active', 'revoked', 'expired');
CREATE TYPE entitlement_source AS ENUM ('subscription', 'purchase', 'admin_grant', 'gift');
CREATE TYPE ledger_entry_type AS ENUM ('credit', 'debit');
CREATE TYPE ledger_category AS ENUM ('earning', 'purchase', 'platform_fee', 'processing_fee', 'tax', 'payout', 'adjustment');
CREATE TYPE payout_status AS ENUM ('requested', 'processing', 'completed', 'failed', 'cancelled');
CREATE TYPE violation_type AS ENUM ('ledger_mismatch', 'subscription_entitlement_mismatch');
CREATE TYPE violation_severity AS ENUM ('critical', 'warning');

-- ============================================================================
-- IDENTITY & PROFILE
-- ============================================================================

CREATE TABLE accounts (
    account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    mfa_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    is_system_account BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_account_email ON accounts (LOWER(email));
CREATE INDEX idx_system_accounts ON accounts (is_system_account) WHERE is_system_account = TRUE;

CREATE TABLE creator_profiles (
    creator_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts(account_id) ON DELETE CASCADE,
    slug VARCHAR(100) NOT NULL UNIQUE,
    display_name VARCHAR(255) NOT NULL,
    verification_status verification_status NOT NULL DEFAULT 'unverified',
    short_description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_creator_account ON creator_profiles(account_id);

-- ============================================================================
-- GOVERNANCE
-- ============================================================================

CREATE TABLE admin_actions (
    action_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    target_account_id UUID NOT NULL REFERENCES accounts(account_id),
    admin_user_id UUID NOT NULL REFERENCES accounts(account_id),
    action_type admin_action_type NOT NULL,
    reason TEXT NOT NULL,
    effective_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ,
    invalidated_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_admin_action_target ON admin_actions(target_account_id, effective_at);
CREATE INDEX idx_admin_action_active ON admin_actions(target_account_id) 
    WHERE invalidated_at IS NULL;

-- ============================================================================
-- PRODUCT & POLICY
-- ============================================================================

CREATE TABLE tiers (
    tier_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    creator_id UUID NOT NULL REFERENCES creator_profiles(creator_id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    price_amount_cents INTEGER NOT NULL CHECK (price_amount_cents >= 0),
    price_currency CHAR(3) NOT NULL DEFAULT 'USD',
    billing_interval billing_interval NOT NULL,
    status tier_status NOT NULL DEFAULT 'active',
    rank INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_creator_rank UNIQUE (creator_id, rank)
);

CREATE INDEX idx_tier_creator ON tiers(creator_id);

CREATE TABLE digital_products (
    product_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    creator_id UUID NOT NULL REFERENCES creator_profiles(creator_id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    price_amount_cents INTEGER NOT NULL CHECK (price_amount_cents >= 0),
    price_currency CHAR(3) NOT NULL DEFAULT 'USD',
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================================
-- RELATIONSHIPS & BILLING
-- ============================================================================

CREATE TABLE memberships (
    membership_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    supporter_account_id UUID NOT NULL REFERENCES accounts(account_id),
    creator_id UUID NOT NULL REFERENCES creator_profiles(creator_id),
    status membership_status NOT NULL DEFAULT 'active',
    lifetime_value_cents INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_supporter_creator UNIQUE (supporter_account_id, creator_id)
);

CREATE INDEX idx_membership_supporter ON memberships(supporter_account_id);
CREATE INDEX idx_membership_creator ON memberships(creator_id);

CREATE TABLE subscriptions (
    subscription_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    membership_id UUID NOT NULL REFERENCES memberships(membership_id),
    tier_id UUID NOT NULL REFERENCES tiers(tier_id),
    provider_sub_id VARCHAR(255),
    status subscription_status NOT NULL DEFAULT 'active',
    current_period_start TIMESTAMPTZ NOT NULL,
    current_period_end TIMESTAMPTZ NOT NULL,
    payment_provider payment_provider NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_subscription_membership ON subscriptions(membership_id);
CREATE UNIQUE INDEX idx_one_active_sub ON subscriptions(membership_id) 
    WHERE status IN ('active', 'past_due');

CREATE TABLE access_entitlements (
    entitlement_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts(account_id),
    resource_id UUID NOT NULL,
    source_subscription_id UUID REFERENCES subscriptions(subscription_id),
    status entitlement_status NOT NULL DEFAULT 'active',
    starts_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ,
    granted_via entitlement_source NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_entitlement_account ON access_entitlements(account_id);
CREATE INDEX idx_entitlement_active ON access_entitlements(account_id, status) 
    WHERE status = 'active';

-- ============================================================================
-- FINANCIAL CORE
-- ============================================================================

CREATE TABLE policy_versions (
    policy_version_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    platform_fee_percent DECIMAL(5,2) NOT NULL CHECK (platform_fee_percent >= 0 AND platform_fee_percent <= 100),
    effective_from TIMESTAMPTZ NOT NULL,
    effective_to TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_policy_effective ON policy_versions(effective_from, effective_to);

CREATE TABLE ledger_entries (
    entry_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sequence BIGSERIAL UNIQUE NOT NULL,
    account_id UUID NOT NULL REFERENCES accounts(account_id),
    entry_type ledger_entry_type NOT NULL,
    amount_minor INTEGER NOT NULL CHECK (amount_minor > 0),
    currency CHAR(3) NOT NULL DEFAULT 'USD',
    category ledger_category NOT NULL,
    reference_id UUID,
    occurred_at TIMESTAMPTZ NOT NULL,
    policy_version_id UUID REFERENCES policy_versions(policy_version_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_ledger_account ON ledger_entries(account_id, occurred_at DESC);
CREATE INDEX idx_ledger_sequence ON ledger_entries(sequence);
CREATE INDEX idx_ledger_reference ON ledger_entries(reference_id);

CREATE TABLE wallets (
    wallet_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL UNIQUE REFERENCES accounts(account_id),
    balance_settled INTEGER NOT NULL DEFAULT 0,
    balance_pending INTEGER NOT NULL DEFAULT 0,
    currency CHAR(3) NOT NULL DEFAULT 'USD',
    last_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_entry_sequence BIGINT
);

CREATE TABLE payouts (
    payout_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    creator_account_id UUID NOT NULL REFERENCES accounts(account_id),
    amount_minor INTEGER NOT NULL CHECK (amount_minor > 0),
    currency CHAR(3) NOT NULL DEFAULT 'USD',
    status payout_status NOT NULL DEFAULT 'requested',
    provider_payout_id VARCHAR(255),
    requested_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMPTZ
);

CREATE INDEX idx_payout_creator ON payouts(creator_account_id);

-- ============================================================================
-- SAFETY & MONITORING
-- ============================================================================

CREATE TABLE invariant_violations (
    violation_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID REFERENCES accounts(account_id),
    violation_type violation_type NOT NULL,
    severity violation_severity NOT NULL,
    details_json JSONB NOT NULL,
    detected_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    resolved_at TIMESTAMPTZ,
    resolution_note TEXT
);

CREATE INDEX idx_violation_unresolved ON invariant_violations(severity, detected_at) 
    WHERE resolved_at IS NULL;
