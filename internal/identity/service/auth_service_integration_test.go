package service

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/hyperx/backend/internal/identity/adapters/postgres"
	"github.com/hyperx/backend/internal/identity/adapters/sessions"
	"github.com/hyperx/backend/internal/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthService_Integration(t *testing.T) {
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
	jwtMgr := sessions.NewJWTManager("secret", 1*time.Hour)
	accSvc := NewAccountService(repo, nil) // Authority not needed for basic creation
	authSvc := NewAuthService(repo, jwtMgr)

	t.Run("Successful Login", func(t *testing.T) {
		email := "auth_test@example.com"
		password := "secret123"

		// Create account
		acc, err := accSvc.CreateAccount(ctx, email, password)
		require.NoError(t, err)
		assert.Equal(t, email, acc.Email)

		// Try login
		token, err := authSvc.Login(ctx, email, password)
		require.NoError(t, err)
		assert.NotEmpty(t, token)

		// Verify token
		claims, err := jwtMgr.Verify(token)
		require.NoError(t, err)
		assert.Equal(t, acc.AccountID, claims.UserID)
	})

	t.Run("Invalid Credentials - Wrong Password", func(t *testing.T) {
		email := "wrong_pass@example.com"
		password := "secret123"

		_, err := accSvc.CreateAccount(ctx, email, password)
		require.NoError(t, err)

		_, err = authSvc.Login(ctx, email, "wrong")
		assert.ErrorIs(t, err, ErrInvalidCredentials)
	})

	t.Run("Invalid Credentials - Missing User", func(t *testing.T) {
		_, err = authSvc.Login(ctx, "nonexistent@example.com", "any")
		assert.ErrorIs(t, err, ErrInvalidCredentials)
	})
}
