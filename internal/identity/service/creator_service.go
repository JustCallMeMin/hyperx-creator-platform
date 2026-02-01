package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hyperx/backend/internal/domain/identity"
	"github.com/hyperx/backend/internal/identity/ports"
	"github.com/hyperx/backend/internal/pkg/authority"
)

type CreatorService struct {
	creatorRepo ports.CreatorRepository
	accountRepo ports.AccountRepository
	authority   *authority.Resolver
}

func NewCreatorService(creatorRepo ports.CreatorRepository, accountRepo ports.AccountRepository, authority *authority.Resolver) *CreatorService {
	return &CreatorService{
		creatorRepo: creatorRepo,
		accountRepo: accountRepo,
		authority:   authority,
	}
}

func (s *CreatorService) CreateProfile(ctx context.Context, accountID uuid.UUID, slug, displayName string) (*identity.CreatorProfile, error) {
	// 1. Verify account exists
	acc, err := s.accountRepo.GetByID(ctx, accountID)
	if err != nil {
		return nil, fmt.Errorf("account not found: %w", err)
	}

	// 2. Check authority (Is account suspended?)
	res, err := s.authority.EvaluateAccess(ctx, accountID, "create_profile") // Using custom string for now or extending ActionType
	if err != nil {
		return nil, fmt.Errorf("authority check failed: %w", err)
	}
	if !res.Allowed {
		return nil, fmt.Errorf("action blocked: %s", res.Reason)
	}

	// 3. Business rule: System accounts cannot be creators
	if !acc.CanActAsCreator() {
		return nil, fmt.Errorf("system accounts cannot create profiles")
	}

	// 4. Create profile
	cp := &identity.CreatorProfile{
		CreatorID:          uuid.New(),
		AccountID:          accountID,
		Slug:               slug,
		DisplayName:        displayName,
		VerificationStatus: identity.VerificationUnverified,
		CreatedAt:          time.Now().UTC(),
	}

	if err := cp.Validate(); err != nil {
		return nil, fmt.Errorf("invalid profile data: %w", err)
	}

	if err := s.creatorRepo.Create(ctx, cp); err != nil {
		return nil, fmt.Errorf("failed to save creator profile: %w", err)
	}

	return cp, nil
}

func (s *CreatorService) GetProfile(ctx context.Context, id uuid.UUID) (*identity.CreatorProfile, error) {
	return s.creatorRepo.GetByID(ctx, id)
}
