package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	gov_postgres "github.com/hyperx/backend/internal/governance/adapters/postgres"
	gov_http "github.com/hyperx/backend/internal/governance/ports/http"
	gov_service "github.com/hyperx/backend/internal/governance/service"
	id_postgres "github.com/hyperx/backend/internal/identity/adapters/postgres"
	"github.com/hyperx/backend/internal/identity/adapters/sessions"
	id_http "github.com/hyperx/backend/internal/identity/ports/http"
	id_service "github.com/hyperx/backend/internal/identity/service"
	"github.com/hyperx/backend/internal/pkg/authority"
	"github.com/hyperx/backend/internal/pkg/database"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Initialize Database
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL must be set")
	}
	pool, err := database.NewPool(ctx, dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize database pool: %v", err)
	}
	defer pool.Close()

	// 2. Initialize Repositories
	accountRepo := id_postgres.NewAccountRepository(pool)
	creatorRepo := id_postgres.NewCreatorRepository(pool)
	adminActionRepo := gov_postgres.NewAdminActionRepository(pool)

	// 3. Initialize Shared Authority Resolver
	authResolver := authority.NewResolver(adminActionRepo)

	// 4. Initialize Core Security & Sessions
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "development-secret-key-please-change"
	}
	jwtDurationStr := os.Getenv("JWT_DURATION_HOURS")
	jwtDuration := 24 * time.Hour
	if d, err := time.ParseDuration(jwtDurationStr + "h"); err == nil {
		jwtDuration = d
	}
	jwtManager := sessions.NewJWTManager(jwtSecret, jwtDuration)

	// 5. Initialize Services
	accountService := id_service.NewAccountService(accountRepo, authResolver)
	creatorService := id_service.NewCreatorService(creatorRepo, accountRepo, authResolver)
	adminService := gov_service.NewAdminService(adminActionRepo)
	authService := id_service.NewAuthService(accountRepo, jwtManager)

	// 6. Initialize HTTP Handlers
	identityHandler := id_http.NewHandler(accountService, creatorService, authService)
	governanceHandler := gov_http.NewHandler(adminService)

	// 6. Setup Router
	r := chi.NewRouter()

	// Global Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"ok","service":"hyperx-backend"}`)
	})

	// Register API Routes
	r.Route("/api/v1", func(r chi.Router) {
		// Public routes (Login & Account Creation)
		r.Group(func(r chi.Router) {
			identityHandler.Routes(r)
		})

		// Protected routes (Governance & Sensitive Identity)
		r.Group(func(r chi.Router) {
			r.Use(id_http.AuthMiddleware(jwtManager))
			governanceHandler.Routes(r)
		})
	})

	// Server configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 7. Graceful shutdown
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
