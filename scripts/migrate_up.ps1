# Migration Up Script
# Runs database migrations using golang-migrate CLI

param(
    [string]$DatabaseURL = $env:DATABASE_URL
)

# Load from .env file if it exists and DATABASE_URL is not set
if ([string]::IsNullOrEmpty($DatabaseURL) -and (Test-Path ".env")) {
    Get-Content ".env" | ForEach-Object {
        if ($_ -match '^DATABASE_URL=(.+)$') {
            $DatabaseURL = $matches[1]
        }
    }
}

# Fallback to default if still not set
if ([string]::IsNullOrEmpty($DatabaseURL)) {
    Write-Error "DATABASE_URL not found in environment or .env file"
    exit 1
}

Write-Host "Using DATABASE_URL from .env"

# Check if migrate CLI is installed
$migrateCmd = Get-Command migrate -ErrorAction SilentlyContinue
if (-not $migrateCmd) {
    Write-Error "migrate CLI not found. Install it with: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
    exit 1
}

Write-Host "Running migrations UP..."
migrate -path migrations -database $DatabaseURL up

if ($LASTEXITCODE -eq 0) {
    Write-Host "Migrations completed successfully!" -ForegroundColor Green
}
else {
    Write-Error "Migration failed with exit code $LASTEXITCODE"
    exit $LASTEXITCODE
}
