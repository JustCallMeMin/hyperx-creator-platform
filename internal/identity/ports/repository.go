package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/hyperx/backend/internal/domain/identity"
)

// AccountRepository defines the port for account data persistence
type AccountRepository interface {
	Create(ctx context.Context, account *identity.Account) error
	GetByID(ctx context.Context, id uuid.UUID) (*identity.Account, error)
	GetByEmail(ctx context.Context, email string) (*identity.Account, error)
	Update(ctx context.Context, account *identity.Account) error
}

// CreatorRepository defines the port for creator profile data persistence
type CreatorRepository interface {
	Create(ctx context.Context, creator *identity.CreatorProfile) error
	GetByID(ctx context.Context, id uuid.UUID) (*identity.CreatorProfile, error)
	GetByAccountID(ctx context.Context, accountID uuid.UUID) (*identity.CreatorProfile, error)
	GetBySlug(ctx context.Context, slug string) (*identity.CreatorProfile, error)
	Update(ctx context.Context, creator *identity.CreatorProfile) error
}
