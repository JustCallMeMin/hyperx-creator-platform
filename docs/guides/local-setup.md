# Local Developer Setup Guide (Windows/No-Docker)
**Author:** DevOps Engineer
**Date:** 2026-02-01

## 1. Prerequisites
- **PostgreSQL:** Installed and running locally (Default port: 5432).
- **Go:** Version 1.22+.
- **PowerShell:** Default on Windows.

## 2. Configuration
1. Copy the configuration template:
   ```powershell
   Copy-Item .env.example .env
   ```
2. **CRITICAL:** Open `.env` and update the database credentials to match your local Postgres installation:
   ```ini
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres      <-- Change to your local user (often 'postgres')
   DB_PASSWORD=yourpass  <-- Change to your local password
   DB_NAME=hyperx        <-- Ensure this database exists!
   ```

3. **Create Database:**
   If you haven't created the database yet, run this in your terminal (assuming `psql` is in PATH):
   ```powershell
   psql -U postgres -c "CREATE DATABASE hyperx;"
   ```

## 3. Managing Migrations
We provide a PowerShell script that builds and runs the Go migration tool.

### Apply Migrations (Up)
```powershell
./migrate.ps1 -Cmd up
```

### Rollback (Down)
```powershell
./migrate.ps1 -Cmd down
```

### Check Version
```powershell
./migrate.ps1 -Cmd version
```

## 4. Running Tests
```powershell
go test ./... -v
```
