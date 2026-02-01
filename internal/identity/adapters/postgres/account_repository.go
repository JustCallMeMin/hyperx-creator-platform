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

var (
	ErrNotFound  = errors.New("resource not found")
	ErrDuplicate = errors.New("resource already exists")
)

type AccountRepository struct {
	pool *pgxpool.Pool
}

func NewAccountRepository(pool *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{pool: pool}
}

func (r *AccountRepository) Create(ctx context.Context, account *identity.Account) error {
	query := `
		INSERT INTO accounts (account_id, email, password_hash, mfa_enabled, is_system_account, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.pool.Exec(ctx, query,
		account.AccountID,
		account.Email,
		account.PasswordHash,
		account.MFAEnabled,
		account.IsSystemAccount,
		account.CreatedAt,
		account.UpdatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // unique_violation
			return ErrDuplicate
		}
		return fmt.Errorf("failed to create account: %w", err)
	}

	return nil
}

func (r *AccountRepository) GetByID(ctx context.Context, id uuid.UUID) (*identity.Account, error) {
	query := `
		SELECT account_id, email, password_hash, mfa_enabled, is_system_account, created_at, updated_at
		FROM accounts
		WHERE account_id = $1
	`
	var acc identity.Account
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&acc.AccountID,
		&acc.Email,
		&acc.PasswordHash,
		&acc.MFAEnabled,
		&acc.IsSystemAccount,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get account by id: %w", err)
	}

	return &acc, nil
}

func (r *AccountRepository) GetByEmail(ctx context.Context, email string) (*identity.Account, error) {
	query := `
		SELECT account_id, email, password_hash, mfa_enabled, is_system_account, created_at, updated_at
		FROM accounts
		WHERE email = $1
	`
	var acc identity.Account
	err := r.pool.QueryRow(ctx, query, email).Scan(
		&acc.AccountID,
		&acc.Email,
		&acc.PasswordHash,
		&acc.MFAEnabled,
		&acc.IsSystemAccount,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get account by email: %w", err)
	}

	return &acc, nil
}

func (r *AccountRepository) Update(ctx context.Context, account *identity.Account) error {
	query := `
		UPDATE accounts
		SET email = $2, password_hash = $3, mfa_enabled = $4, is_system_account = $5, updated_at = $6
		WHERE account_id = $1
	`
	tag, err := r.pool.Exec(ctx, query,
		account.AccountID,
		account.Email,
		account.PasswordHash,
		account.MFAEnabled,
		account.IsSystemAccount,
		account.UpdatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrDuplicate
		}
		return fmt.Errorf("failed to update account: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
