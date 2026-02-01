# HyperX Backend - Makefile

.PHONY: help migrate-up migrate-down migrate-create server test

help:
	@echo "Available targets:"
	@echo "  migrate-up      - Run database migrations (up)"
	@echo "  migrate-down    - Rollback database migrations (down)"
	@echo "  migrate-create  - Create new migration file"
	@echo "  server          - Run the server"
	@echo "  test            - Run tests"

migrate-up:
	@powershell -ExecutionPolicy Bypass -File scripts/migrate_up.ps1

migrate-down:
	@powershell -ExecutionPolicy Bypass -File scripts/migrate_down.ps1

migrate-create:
	@powershell -ExecutionPolicy Bypass -File scripts/migrate_create.ps1

server:
	go run cmd/server/main.go

test:
	go test ./...
