# Database Setup Script
# Drops and recreates the hyperx database, then runs all migrations

param(
    [string]$Password = "332001"
)

$env:PGPASSWORD = $Password

Write-Host "Dropping existing database..." -ForegroundColor Yellow
psql -U postgres -h localhost -d postgres -c "DROP DATABASE IF EXISTS hyperx;"

Write-Host "Creating fresh database..." -ForegroundColor Yellow
psql -U postgres -h localhost -d postgres -c "CREATE DATABASE hyperx;"

Write-Host "Running migration 000001..." -ForegroundColor Cyan
psql -U postgres -h localhost -d hyperx -f migrations/000001_init_schema.up.sql

if ($LASTEXITCODE -eq 0) {
    Write-Host "Migration 000001 completed!" -ForegroundColor Green
    
    Write-Host "Running migration 000002..." -ForegroundColor Cyan
    psql -U postgres -h localhost -d hyperx -f migrations/000002_db_optimizations.up.sql
    
    if ($LASTEXITCODE -eq 0) {
        Write-Host "All migrations completed successfully!" -ForegroundColor Green
    }
    else {
        Write-Error "Migration 000002 failed!"
        exit 1
    }
}
else {
    Write-Error "Migration 000001 failed!"
    exit 1
}

Remove-Item Env:\PGPASSWORD
