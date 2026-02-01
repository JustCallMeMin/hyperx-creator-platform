package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	gov_postgres "github.com/hyperx/backend/internal/governance/adapters/postgres"
	gov_service "github.com/hyperx/backend/internal/governance/service"
	id_postgres "github.com/hyperx/backend/internal/identity/adapters/postgres"
	id_service "github.com/hyperx/backend/internal/identity/service"
	"github.com/hyperx/backend/internal/pkg/authority"
	"github.com/hyperx/backend/internal/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGovernanceHandler_Integration(t *testing.T) {
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

	// Dependencies
	accRepo := id_postgres.NewAccountRepository(pool)
	admRepo := gov_postgres.NewAdminActionRepository(pool)
	admSvc := gov_service.NewAdminService(admRepo)

	// Create some accounts for testing
	accSvc := id_service.NewAccountService(accRepo, authority.NewResolver(admRepo))
	targetAcc, _ := accSvc.CreateAccount(ctx, "target@example.com", "pass")
	adminAcc, _ := accSvc.CreateAccount(ctx, "admin@example.com", "pass")

	h := NewHandler(admSvc)
	r := chi.NewRouter()
	h.Routes(r)

	t.Run("Suspend Account API", func(t *testing.T) {
		reqBody, _ := json.Marshal(SuspendRequest{
			TargetAccountID: targetAcc.AccountID,
			AdminUserID:     adminAcc.AccountID,
			Reason:          "violating terms",
			DurationSeconds: 3600,
		})
		req, _ := http.NewRequest("POST", "/admin/actions/suspend", bytes.NewBuffer(reqBody))
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)

		// Verify it exists in DB
		actions, err := admRepo.GetActiveByTargetAccount(ctx, targetAcc.AccountID, time.Now().UTC())
		require.NoError(t, err)
		assert.Len(t, actions, 1)
		assert.Equal(t, "violating terms", actions[0].Reason)
	})
}
