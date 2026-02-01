# Migration Down Script
# Rolls back database migrations using golang-migrate CLI

param(
    [string]$DatabaseURL = $env:DATABASE_URL,
    [int]$Steps = 1
)

# Load from .env file if it exists and DATABASE_URL is not set
if ([string]::IsNullOrEmpty($DatabaseURL) -and (Test-Path ".env")) {
    Get-Content ".env" | ForEach-Object {
        if ($_ -match '^DATABASE_URL=(.+)$') {
            $DatabaseURL = $matches[1]
        }
    }
}

# Fallback to error if still not set
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

Write-Host "Rolling back $Steps migration(s)..."
migrate -path migrations -database $DatabaseURL down $Steps

if ($LASTEXITCODE -eq 0) {
    Write-Host "Rollback completed successfully!" -ForegroundColor Green
}
else {
    Write-Error "Rollback failed with exit code $LASTEXITCODE"
    exit $LASTEXITCODE
}
