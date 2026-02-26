# AGENTS.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview
This is a Go REST API built with the Gin web framework, designed to be a todo application. The project uses PostgreSQL via pgx driver, JWT for authentication, and golang-migrate for database migrations.

**Module name**: `todo_api` (as defined in go.mod)

## Architecture

### Project Structure
Standard Go project layout:
- `cmd/api/main.go` - Application entry point (note: package is `cmd`, not `main`)
- `internal/` - Private application code organized by layer:
  - `config/` - Configuration management
  - `database/` - Database connection setup
  - `handlers/` - HTTP request handlers (Gin handlers)
  - `middleware/` - HTTP middleware (likely auth, logging, etc.)
  - `models/` - Data models/structs
  - `repository/` - Data access layer (repository pattern)
- `migrations/` - Database migration files (golang-migrate)

### Key Dependencies
- **Web Framework**: Gin (github.com/gin-gonic/gin)
- **Database**: PostgreSQL via pgx/v5 with connection pooling (pgxpool)
- **Auth**: JWT tokens (github.com/golang-jwt/jwt/v5)
- **Migrations**: golang-migrate with postgres tag
- **Config**: godotenv for environment variables

## Common Commands

### Development
```powershell
# Run with hot reload (using air)
air

# Build manually
go build -o ./tmp/main.exe ./cmd/api

# Run without hot reload
go run ./cmd/api
```

### Database Migrations
```powershell
# Create a new migration
migrate create -ext sql -dir migrations -seq <migration_name>

# Run migrations
migrate -path migrations -database "postgresql://user:password@localhost:5432/dbname?sslmode=disable" up

# Rollback migrations
migrate -path migrations -database "postgresql://user:password@localhost:5432/dbname?sslmode=disable" down
```

### Dependencies
```powershell
# Install dependencies (if go.mod exists)
go mod download

# Update dependencies
go get -u ./...:

# Tidy dependencies
go mod tidy
```

## Important Notes

### Air Configuration Issue
The `.air.toml` file uses the deprecated `bin` setting. If air fails to run properly:
1. Update `.air.toml` line 8 from `bin = "./tmp/main.exe"` to `build.entrypoint = "./tmp/main.exe"` (or remove the line as it's auto-detected)
2. Clear the tmp directory: `Remove-Item -Path ./tmp/* -Force`
3. Restart air

### Main Package Declaration
The `cmd/api/main.go` file declares `package cmd` instead of `package main`. This is non-standard and may cause issues - the main function in the entry point should be in `package main`.

### Environment Variables
The project uses a `.env` file (via godotenv). This likely contains:
- Database connection strings
- JWT secrets
- Server configuration
- Build environment variables (currently contains GOOS, GOARCH, CGO_ENABLED)

### API Server
The API runs on port 3000 by default (configured in main.go with `router.Run(":3000")`).

## Development Workflow
1. Ensure PostgreSQL is running and accessible
2. Load environment variables from `.env`
3. Run database migrations before starting the server
4. Use `air` for development with hot reload
5. The main entry endpoint returns a health check at GET /
