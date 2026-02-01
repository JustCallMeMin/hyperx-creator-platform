package governance

import (
	"time"

	"github.com/google/uuid"
)

// AdminActionType represents the category of governance action
type AdminActionType string

const (
	ActionSuspendAccount   AdminActionType = "suspend_account"
	ActionFreezePayout     AdminActionType = "freeze_payout"
	ActionRevokeAccess     AdminActionType = "revoke_access"
	ActionGrantEntitlement AdminActionType = "grant_entitlement"
)

// AdminAction represents a governance fact issued by an admin
type AdminAction struct {
	ActionID        uuid.UUID       `json:"action_id" db:"action_id"`
	ActionType      AdminActionType `json:"action_type" db:"action_type"`
	TargetAccountID uuid.UUID       `json:"target_account_id" db:"target_account_id"`
	Reason          string          `json:"reason" db:"reason"`
	AdminUserID     uuid.UUID       `json:"admin_user_id" db:"admin_user_id"`
	EffectiveAt     time.Time       `json:"effective_at" db:"effective_at"`
	ExpiresAt       *time.Time      `json:"expires_at,omitempty" db:"expires_at"`
	InvalidatedAt   *time.Time      `json:"invalidated_at,omitempty" db:"invalidated_at"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
}

// IsEffective checks if the action is currently active based on temporal rules
func (a *AdminAction) IsEffective(now time.Time) bool {
	if a.InvalidatedAt != nil {
		return false
	}
	if now.Before(a.EffectiveAt) {
		return false
	}
	if a.ExpiresAt != nil && now.After(*a.ExpiresAt) {
		return false
	}
	return true
}

// IsExpired checks if the action has naturally reached its end date
func (a *AdminAction) IsExpired(now time.Time) bool {
	if a.ExpiresAt == nil {
		return false
	}
	return now.After(*a.ExpiresAt)
}

// Invalidate marks the action as no longer effective
func (a *AdminAction) Invalidate(now time.Time) {
	a.InvalidatedAt = &now
}
