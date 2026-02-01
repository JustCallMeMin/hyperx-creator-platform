$DatabaseURL = "postgres://postgres:332001@localhost:5432/postgres?sslmode=disable"
$testDbName = "hyperx_test"

# Drop if exists and create
Write-Host "Setting up test database: $testDbName"
$env:PGPASSWORD = "332001"
psql -U postgres -h localhost -p 5432 -c "DROP DATABASE IF EXISTS $testDbName;"
psql -U postgres -h localhost -p 5432 -c "CREATE DATABASE $testDbName;"

# Migrate test database
$testDbURL = "postgres://postgres:332001@localhost:5432/$testDbName?sslmode=disable"
./scripts/migrate_up.ps1 -DatabaseURL $testDbURL

Write-Host "Test database setup complete."
