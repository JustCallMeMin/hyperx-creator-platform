package entity

import (
	"time"

	"github.com/google/uuid"
)

// LedgerEntryType represents credit or debit
type LedgerEntryType string

const (
	LedgerCredit LedgerEntryType = "credit"
	LedgerDebit  LedgerEntryType = "debit"
)

// LedgerCategory represents the business purpose of the entry
type LedgerCategory string

const (
	CategoryEarning       LedgerCategory = "earning"
	CategoryPurchase      LedgerCategory = "purchase"
	CategoryPlatformFee   LedgerCategory = "platform_fee"
	CategoryProcessingFee LedgerCategory = "processing_fee"
	CategoryTax           LedgerCategory = "tax"
	CategoryPayout        LedgerCategory = "payout"
	CategoryAdjustment    LedgerCategory = "adjustment"
)

// LedgerEntry represents an immutable financial fact
type LedgerEntry struct {
	EntryID         uuid.UUID       `json:"entry_id" db:"entry_id"`
	Sequence        int64           `json:"sequence" db:"sequence"`
	AccountID       uuid.UUID       `json:"account_id" db:"account_id"`
	EntryType       LedgerEntryType `json:"entry_type" db:"entry_type"`
	AmountMinor     int             `json:"amount_minor" db:"amount_minor"`
	Currency        string          `json:"currency" db:"currency"`
	Category        LedgerCategory  `json:"category" db:"category"`
	ReferenceID     *uuid.UUID      `json:"reference_id,omitempty" db:"reference_id"`
	OccurredAt      time.Time       `json:"occurred_at" db:"occurred_at"`
	PolicyVersionID *uuid.UUID      `json:"policy_version_id,omitempty" db:"policy_version_id"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
}

// PolicyVersion represents a snapshot of platform fee rules
type PolicyVersion struct {
	PolicyVersionID    uuid.UUID  `json:"policy_version_id" db:"policy_version_id"`
	PlatformFeePercent float64    `json:"platform_fee_percent" db:"platform_fee_percent"`
	EffectiveFrom      time.Time  `json:"effective_from" db:"effective_from"`
	EffectiveTo        *time.Time `json:"effective_to,omitempty" db:"effective_to"`
	CreatedAt          time.Time  `json:"created_at" db:"created_at"`
}

// Wallet represents a projected balance (CQRS Read Model)
type Wallet struct {
	WalletID          uuid.UUID `json:"wallet_id" db:"wallet_id"`
	AccountID         uuid.UUID `json:"account_id" db:"account_id"`
	BalanceSettled    int       `json:"balance_settled" db:"balance_settled"`
	BalancePending    int       `json:"balance_pending" db:"balance_pending"`
	Currency          string    `json:"currency" db:"currency"`
	LastUpdatedAt     time.Time `json:"last_updated_at" db:"last_updated_at"`
	LastEntrySequence *int64    `json:"last_entry_sequence,omitempty" db:"last_entry_sequence"`
}
