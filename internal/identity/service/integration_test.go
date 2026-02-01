package service

import (
	"context"
	"os"
	"testing"

	"github.com/hyperx/backend/internal/identity/adapters/postgres"
	"github.com/hyperx/backend/internal/pkg/authority"
	"github.com/hyperx/backend/internal/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountService_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:332001@localhost:5432/hyperx_test?sslmode=disable"
	}

	pool, err := database.NewPool(ctx, dbURL)
	require.NoError(t, err)
	defer pool.Close()

	// Cleanup
	pool.Exec(ctx, "TRUNCATE accounts CASCADE")

	repo := postgres.NewAccountRepository(pool)
	authResolver := authority.NewResolver(nil) // Note: No admin action repo needed for basic create
	svc := NewAccountService(repo, authResolver)

	t.Run("Create and Get Account", func(t *testing.T) {
		email := "integration@example.com"
		acc, err := svc.CreateAccount(ctx, email, "password123")
		require.NoError(t, err)
		assert.Equal(t, email, acc.Email)

		fetched, err := svc.GetAccount(ctx, acc.AccountID)
		require.NoError(t, err)
		assert.Equal(t, acc.AccountID, fetched.AccountID)
		assert.Equal(t, email, fetched.Email)
	})

	t.Run("Duplicate Email", func(t *testing.T) {
		email := "dup@example.com"
		_, err := svc.CreateAccount(ctx, email, "pass")
		require.NoError(t, err)

		_, err = svc.CreateAccount(ctx, email, "pass")
		assert.Error(t, err)
	})
}
