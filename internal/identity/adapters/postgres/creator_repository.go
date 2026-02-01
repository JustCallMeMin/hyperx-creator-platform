package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/hyperx/backend/internal/domain/identity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreatorRepository struct {
	pool *pgxpool.Pool
}

func NewCreatorRepository(pool *pgxpool.Pool) *CreatorRepository {
	return &CreatorRepository{pool: pool}
}

func (r *CreatorRepository) Create(ctx context.Context, creator *identity.CreatorProfile) error {
	query := `
		INSERT INTO creator_profiles (creator_id, account_id, slug, display_name, verification_status, short_description, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.pool.Exec(ctx, query,
		creator.CreatorID,
		creator.AccountID,
		creator.Slug,
		creator.DisplayName,
		creator.VerificationStatus,
		creator.ShortDescription,
		creator.CreatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrDuplicate
		}
		return fmt.Errorf("failed to create creator profile: %w", err)
	}

	return nil
}

func (r *CreatorRepository) GetByID(ctx context.Context, id uuid.UUID) (*identity.CreatorProfile, error) {
	query := `
		SELECT creator_id, account_id, slug, display_name, verification_status, short_description, created_at
		FROM creator_profiles
		WHERE creator_id = $1
	`
	var cp identity.CreatorProfile
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&cp.CreatorID,
		&cp.AccountID,
		&cp.Slug,
		&cp.DisplayName,
		&cp.VerificationStatus,
		&cp.ShortDescription,
		&cp.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get creator profile by id: %w", err)
	}

	return &cp, nil
}

func (r *CreatorRepository) GetByAccountID(ctx context.Context, accountID uuid.UUID) (*identity.CreatorProfile, error) {
	query := `
		SELECT creator_id, account_id, slug, display_name, verification_status, short_description, created_at
		FROM creator_profiles
		WHERE account_id = $1
	`
	var cp identity.CreatorProfile
	err := r.pool.QueryRow(ctx, query, accountID).Scan(
		&cp.CreatorID,
		&cp.AccountID,
		&cp.Slug,
		&cp.DisplayName,
		&cp.VerificationStatus,
		&cp.ShortDescription,
		&cp.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get creator profile by account id: %w", err)
	}

	return &cp, nil
}

func (r *CreatorRepository) GetBySlug(ctx context.Context, slug string) (*identity.CreatorProfile, error) {
	query := `
		SELECT creator_id, account_id, slug, display_name, verification_status, short_description, created_at
		FROM creator_profiles
		WHERE slug = $1
	`
	var cp identity.CreatorProfile
	err := r.pool.QueryRow(ctx, query, slug).Scan(
		&cp.CreatorID,
		&cp.AccountID,
		&cp.Slug,
		&cp.DisplayName,
		&cp.VerificationStatus,
		&cp.ShortDescription,
		&cp.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get creator profile by slug: %w", err)
	}

	return &cp, nil
}

func (r *CreatorRepository) Update(ctx context.Context, creator *identity.CreatorProfile) error {
	query := `
		UPDATE creator_profiles
		SET slug = $2, display_name = $3, verification_status = $4, short_description = $5
		WHERE creator_id = $1
	`
	tag, err := r.pool.Exec(ctx, query,
		creator.CreatorID,
		creator.Slug,
		creator.DisplayName,
		creator.VerificationStatus,
		creator.ShortDescription,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrDuplicate
		}
		return fmt.Errorf("failed to update creator profile: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
