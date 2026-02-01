package service

import (
	"context"
	"os"
	"testing"

	governance_postgres "github.com/hyperx/backend/internal/governance/adapters/postgres"
	identity_postgres "github.com/hyperx/backend/internal/identity/adapters/postgres"
	"github.com/hyperx/backend/internal/pkg/authority"
	"github.com/hyperx/backend/internal/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatorService_Integration(t *testing.T) {
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
	pool.Exec(ctx, "TRUNCATE admin_actions CASCADE")

	accRepo := identity_postgres.NewAccountRepository(pool)
	creRepo := identity_postgres.NewCreatorRepository(pool)
	admRepo := governance_postgres.NewAdminActionRepository(pool)

	authResolver := authority.NewResolver(admRepo)

	accSvc := NewAccountService(accRepo, authResolver)
	creSvc := NewCreatorService(creRepo, accRepo, authResolver)

	t.Run("Create Creator Profile", func(t *testing.T) {
		acc, err := accSvc.CreateAccount(ctx, "creator@example.com", "pass")
		require.NoError(t, err)

		cp, err := creSvc.CreateProfile(ctx, acc.AccountID, "myslug", "My Display Name")
		require.NoError(t, err)
		assert.Equal(t, acc.AccountID, cp.AccountID)
		assert.Equal(t, "myslug", cp.Slug)

		fetched, err := creSvc.GetProfile(ctx, cp.CreatorID)
		require.NoError(t, err)
		assert.Equal(t, cp.CreatorID, fetched.CreatorID)
	})

	t.Run("System Account Cannot be Creator", func(t *testing.T) {
		// Manually create a system account (usually we'd have a service method for this)
		// For now, we'll just test the logic
		acc, err := accSvc.CreateAccount(ctx, "system@example.com", "pass")
		require.NoError(t, err)

		_, err = pool.Exec(ctx, "UPDATE accounts SET is_system_account = true WHERE account_id = $1", acc.AccountID)
		require.NoError(t, err)

		_, err = creSvc.CreateProfile(ctx, acc.AccountID, "sysslug", "Sys Creator")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "system accounts cannot create profiles")
	})
}
