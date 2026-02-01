package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hyperx/backend/internal/domain/governance"
	"github.com/hyperx/backend/internal/governance/ports"
)

type AdminService struct {
	repo ports.AdminActionRepository
}

func NewAdminService(repo ports.AdminActionRepository) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) SuspendAccount(ctx context.Context, targetAccountID, adminUserID uuid.UUID, reason string, duration time.Duration) error {
	now := time.Now().UTC()
	var expiresAt *time.Time
	if duration > 0 {
		exp := now.Add(duration)
		expiresAt = &exp
	}

	action := &governance.AdminAction{
		ActionID:        uuid.New(),
		ActionType:      governance.ActionSuspendAccount,
		TargetAccountID: targetAccountID,
		Reason:          reason,
		AdminUserID:     adminUserID,
		EffectiveAt:     now,
		ExpiresAt:       expiresAt,
		CreatedAt:       now,
	}

	return s.repo.Create(ctx, action)
}

func (s *AdminService) FreezePayout(ctx context.Context, targetAccountID, adminUserID uuid.UUID, reason string) error {
	now := time.Now().UTC()
	action := &governance.AdminAction{
		ActionID:        uuid.New(),
		ActionType:      governance.ActionFreezePayout,
		TargetAccountID: targetAccountID,
		Reason:          reason,
		AdminUserID:     adminUserID,
		EffectiveAt:     now,
		CreatedAt:       now,
	}

	return s.repo.Create(ctx, action)
}

func (s *AdminService) RevokeAction(ctx context.Context, actionID uuid.UUID) error {
	return s.repo.Invalidate(ctx, actionID, time.Now().UTC())
}
