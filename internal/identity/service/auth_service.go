package service

import (
	"context"
	"errors"

	"github.com/hyperx/backend/internal/identity/adapters/sessions"
	"github.com/hyperx/backend/internal/identity/ports"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type AuthService struct {
	repo       ports.AccountRepository
	jwtManager *sessions.JWTManager
}

func NewAuthService(repo ports.AccountRepository, jwtManager *sessions.JWTManager) *AuthService {
	return &AuthService{
		repo:       repo,
		jwtManager: jwtManager,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	acc, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		// Hide details to avoid email enumeration
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.PasswordHash), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	return s.jwtManager.Generate(acc.AccountID)
}
