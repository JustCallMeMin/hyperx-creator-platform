package authority

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hyperx/backend/internal/domain/governance"
	"github.com/hyperx/backend/internal/governance/ports"
)

type ActionType string

const (
	ActionViewContent   ActionType = "view_content"
	ActionRequestPayout ActionType = "request_payout"
	ActionCreateTier    ActionType = "create_tier"
	ActionSubscribe     ActionType = "subscribe"
)

type Result struct {
	Allowed bool
	Reason  string
	Scope   string
}

type Resolver struct {
	adminActionRepo ports.AdminActionRepository
}

func NewResolver(adminActionRepo ports.AdminActionRepository) *Resolver {
	return &Resolver{
		adminActionRepo: adminActionRepo,
	}
}

// EvaluateAccess checks if an account is allowed to perform a specific action
func (r *Resolver) EvaluateAccess(ctx context.Context, accountID uuid.UUID, action ActionType) (Result, error) {
	now := time.Now().UTC()
	actions, err := r.adminActionRepo.GetActiveByTargetAccount(ctx, accountID, now)
	if err != nil {
		return Result{Allowed: false, Reason: "internal_error"}, fmt.Errorf("failed to fetch admin actions: %w", err)
	}

	for _, a := range actions {
		switch a.ActionType {
		case governance.ActionSuspendAccount:
			return Result{
				Allowed: false,
				Reason:  fmt.Sprintf("account suspended: %s", a.Reason),
				Scope:   "account",
			}, nil

		case governance.ActionFreezePayout:
			if action == ActionRequestPayout {
				return Result{
					Allowed: false,
					Reason:  fmt.Sprintf("payouts frozen: %s", a.Reason),
					Scope:   "creator",
				}, nil
			}

		case governance.ActionRevokeAccess:
			if action == ActionViewContent {
				// Note: In a real scenario, we'd check if the target_id matches the specific content/membership
				return Result{
					Allowed: false,
					Reason:  fmt.Sprintf("access revoked: %s", a.Reason),
					Scope:   "membership",
				}, nil
			}
		}
	}

	return Result{Allowed: true}, nil
}
