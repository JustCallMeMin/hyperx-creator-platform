package identity

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
)

// VerificationStatus represents the KYC/verification state of a creator
type VerificationStatus string

const (
	VerificationUnverified VerificationStatus = "unverified"
	VerificationPending    VerificationStatus = "pending"
	VerificationVerified   VerificationStatus = "verified"
	VerificationRejected   VerificationStatus = "rejected"
)

// Account represents the central identity container
type Account struct {
	AccountID       uuid.UUID `json:"account_id" db:"account_id"`
	Email           string    `json:"email" db:"email"`
	PasswordHash    string    `json:"-" db:"password_hash"`
	MFAEnabled      bool      `json:"mfa_enabled" db:"mfa_enabled"`
	IsSystemAccount bool      `json:"is_system_account" db:"is_system_account"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate performs business validation on Account
func (a *Account) Validate() error {
	if a.Email == "" {
		return ErrInvalidEmail
	}
	if !isValidEmail(a.Email) {
		return ErrInvalidEmail
	}
	if a.PasswordHash == "" && !a.IsSystemAccount {
		return ErrPasswordRequired
	}
	return nil
}

// CanActAsCreator checks if account can create a creator profile
func (a *Account) CanActAsCreator() bool {
	return !a.IsSystemAccount
}

// isValidEmail performs basic email validation
func isValidEmail(email string) bool {
	// Simple validation - contains @ and domain
	if len(email) < 3 {
		return false
	}
	atIndex := -1
	for i, c := range email {
		if c == '@' {
			atIndex = i
			break
		}
	}
	return atIndex > 0 && atIndex < len(email)-1
}

// Domain errors
var (
	ErrInvalidEmail     = fmt.Errorf("invalid email format")
	ErrPasswordRequired = fmt.Errorf("password is required for non-system accounts")
)

// CreatorProfile represents the creator role/context
type CreatorProfile struct {
	CreatorID          uuid.UUID          `json:"creator_id" db:"creator_id"`
	AccountID          uuid.UUID          `json:"account_id" db:"account_id"`
	Slug               string             `json:"slug" db:"slug"`
	DisplayName        string             `json:"display_name" db:"display_name"`
	VerificationStatus VerificationStatus `json:"verification_status" db:"verification_status"`
	ShortDescription   *string            `json:"short_description,omitempty" db:"short_description"`
	CreatedAt          time.Time          `json:"created_at" db:"created_at"`
}

var slugRegex = regexp.MustCompile(`^[a-z0-9-]+$`)

// Validate performs business validation on CreatorProfile
func (c *CreatorProfile) Validate() error {
	if c.Slug == "" {
		return ErrSlugRequired
	}
	if !c.ValidateSlug() {
		return ErrInvalidSlug
	}
	if c.DisplayName == "" {
		return ErrDisplayNameRequired
	}
	if len(c.DisplayName) > 255 {
		return ErrDisplayNameTooLong
	}
	return nil
}

// ValidateSlug checks if slug is alphanumeric with hyphens only
func (c *CreatorProfile) ValidateSlug() bool {
	if len(c.Slug) < 3 || len(c.Slug) > 100 {
		return false
	}
	return slugRegex.MatchString(c.Slug)
}

// CanReceivePayouts checks if creator can receive payouts
func (c *CreatorProfile) CanReceivePayouts() bool {
	return c.VerificationStatus == VerificationVerified
}

// IsVerified checks if creator is verified
func (c *CreatorProfile) IsVerified() bool {
	return c.VerificationStatus == VerificationVerified
}

// Creator-specific errors
var (
	ErrSlugRequired        = fmt.Errorf("slug is required")
	ErrInvalidSlug         = fmt.Errorf("slug must be lowercase alphanumeric with hyphens")
	ErrDisplayNameRequired = fmt.Errorf("display name is required")
	ErrDisplayNameTooLong  = fmt.Errorf("display name must be 255 characters or less")
)
