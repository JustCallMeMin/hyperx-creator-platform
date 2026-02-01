package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hyperx/backend/internal/domain/identity"
	"github.com/hyperx/backend/internal/identity/ports"
	"github.com/hyperx/backend/internal/pkg/authority"
	"golang.org/x/crypto/bcrypt"
)

type AccountService struct {
	repo      ports.AccountRepository
	authority *authority.Resolver
}

func NewAccountService(repo ports.AccountRepository, authority *authority.Resolver) *AccountService {
	return &AccountService{
		repo:      repo,
		authority: authority,
	}
}

func (s *AccountService) CreateAccount(ctx context.Context, email, password string) (*identity.Account, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	acc := &identity.Account{
		AccountID:       uuid.New(),
		Email:           email,
		PasswordHash:    string(hashedPassword),
		MFAEnabled:      false,
		IsSystemAccount: false,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}

	if err := acc.Validate(); err != nil {
		return nil, fmt.Errorf("invalid account data: %w", err)
	}

	if err := s.repo.Create(ctx, acc); err != nil {
		return nil, fmt.Errorf("failed to save account: %w", err)
	}

	return acc, nil
}

func (s *AccountService) GetAccount(ctx context.Context, id uuid.UUID) (*identity.Account, error) {
	// Check if account is suspended via Authority Resolver
	// Note: Even for a simple GET, the Authority Resolver might define if the account is "visible"
	// For now, only write operations are strictly enforced, but we'll demonstrate the pattern.

	acc, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("account not found: %w", err)
	}

	return acc, nil
}
