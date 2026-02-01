package ports

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hyperx/backend/internal/domain/governance"
)

// AdminActionRepository defines the port for governance action persistence
type AdminActionRepository interface {
	Create(ctx context.Context, action *governance.AdminAction) error
	GetActiveByTargetAccount(ctx context.Context, accountID uuid.UUID, now time.Time) ([]*governance.AdminAction, error)
	Invalidate(ctx context.Context, actionID uuid.UUID, now time.Time) error
}
