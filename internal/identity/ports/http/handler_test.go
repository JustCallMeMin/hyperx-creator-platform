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
	id_postgres "github.com/hyperx/backend/internal/identity/adapters/postgres"
	"github.com/hyperx/backend/internal/identity/adapters/sessions"
	id_service "github.com/hyperx/backend/internal/identity/service"
	"github.com/hyperx/backend/internal/pkg/authority"
	"github.com/hyperx/backend/internal/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_Integration(t *testing.T) {
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
	creRepo := id_postgres.NewCreatorRepository(pool)
	authResolver := authority.NewResolver(nil)
	accSvc := id_service.NewAccountService(accRepo, authResolver)
	creSvc := id_service.NewCreatorService(creRepo, accRepo, authResolver)
	authSvc := id_service.NewAuthService(accRepo, sessions.NewJWTManager("secret", 1*time.Hour))

	h := NewHandler(accSvc, creSvc, authSvc)
	r := chi.NewRouter()
	h.Routes(r)

	t.Run("Create Account API", func(t *testing.T) {
		reqBody, _ := json.Marshal(CreateAccountRequest{
			Email:    "api@example.com",
			Password: "password123",
		})
		req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(reqBody))
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)

		var resp map[string]interface{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, "api@example.com", resp["email"])
	})

	t.Run("Get Account API", func(t *testing.T) {
		// First create one
		acc, _ := accSvc.CreateAccount(ctx, "get@example.com", "pass")

		req, _ := http.NewRequest("GET", "/accounts/"+acc.AccountID.String(), nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "get@example.com")
	})
}
