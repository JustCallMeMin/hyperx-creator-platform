package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hyperx/backend/internal/domain/governance"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("resource not found")

type AdminActionRepository struct {
	pool *pgxpool.Pool
}

func NewAdminActionRepository(pool *pgxpool.Pool) *AdminActionRepository {
	return &AdminActionRepository{pool: pool}
}

func (r *AdminActionRepository) Create(ctx context.Context, action *governance.AdminAction) error {
	query := `
		INSERT INTO admin_actions (action_id, action_type, target_account_id, reason, admin_user_id, effective_at, expires_at, invalidated_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.pool.Exec(ctx, query,
		action.ActionID,
		action.ActionType,
		action.TargetAccountID,
		action.Reason,
		action.AdminUserID,
		action.EffectiveAt,
		action.ExpiresAt,
		action.InvalidatedAt,
		action.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create admin action: %w", err)
	}

	return nil
}

func (r *AdminActionRepository) GetActiveByTargetAccount(ctx context.Context, accountID uuid.UUID, now time.Time) ([]*governance.AdminAction, error) {
	query := `
		SELECT action_id, action_type, target_account_id, reason, admin_user_id, effective_at, expires_at, invalidated_at, created_at
		FROM admin_actions
		WHERE target_account_id = $1 
		  AND effective_at <= $2 
		  AND (expires_at IS NULL OR expires_at > $2)
		  AND invalidated_at IS NULL
	`
	rows, err := r.pool.Query(ctx, query, accountID, now)
	if err != nil {
		return nil, fmt.Errorf("failed to query active admin actions: %w", err)
	}
	defer rows.Close()

	var actions []*governance.AdminAction
	for rows.Next() {
		var a governance.AdminAction
		err := rows.Scan(
			&a.ActionID,
			&a.ActionType,
			&a.TargetAccountID,
			&a.Reason,
			&a.AdminUserID,
			&a.EffectiveAt,
			&a.ExpiresAt,
			&a.InvalidatedAt,
			&a.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan admin action: %w", err)
		}
		actions = append(actions, &a)
	}

	return actions, nil
}

func (r *AdminActionRepository) Invalidate(ctx context.Context, actionID uuid.UUID, now time.Time) error {
	query := `
		UPDATE admin_actions
		SET invalidated_at = $2
		WHERE action_id = $1
	`
	tag, err := r.pool.Exec(ctx, query, actionID, now)
	if err != nil {
		return fmt.Errorf("failed to invalidate admin action: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
