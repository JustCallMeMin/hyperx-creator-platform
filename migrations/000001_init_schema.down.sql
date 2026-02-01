-- Migration: Rollback Initial Schema
-- Version: 000001

DROP TABLE IF EXISTS invariant_violations;
DROP TABLE IF EXISTS payouts;
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS ledger_entries;
DROP TABLE IF EXISTS policy_versions;
DROP TABLE IF EXISTS access_entitlements;
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS memberships;
DROP TABLE IF EXISTS digital_products;
DROP TABLE IF EXISTS tiers;
DROP TABLE IF EXISTS admin_actions;
DROP TABLE IF EXISTS creator_profiles;
DROP TABLE IF EXISTS accounts;

DROP TYPE IF EXISTS violation_severity;
DROP TYPE IF EXISTS violation_type;
DROP TYPE IF EXISTS payout_status;
DROP TYPE IF EXISTS ledger_category;
DROP TYPE IF EXISTS ledger_entry_type;
DROP TYPE IF EXISTS entitlement_source;
DROP TYPE IF EXISTS entitlement_status;
DROP TYPE IF EXISTS payment_provider;
DROP TYPE IF EXISTS subscription_status;
DROP TYPE IF EXISTS membership_status;
DROP TYPE IF EXISTS billing_interval;
DROP TYPE IF EXISTS tier_status;
DROP TYPE IF EXISTS admin_action_type;
DROP TYPE IF EXISTS verification_status;
