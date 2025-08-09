# Go HTTP Template

This project is a Go HTTP server with PostgreSQL and database migrations.

---

## Setup

### Environment variables

Create a `.env` file or export these environment variables:

```bash
PORT=8080
DB_HOST=localhost
DB_USER=myuser
DB_PASSWORD=mypassword
DB_PORT=5432
DB_NAME=mydatabase
DB_SSL_MODE=disable
```

## Build
Compile the Go application and output the binary to bin/go-http-template

```
make build
```

## Run
Build the application and run the compiled binary:
```
make run
```

### Database Migrations
You can manage your database schema migrations using the following commands:

# Create a new migration file:

Replace name_of_migration with a descriptive name for your migration.

```
make migrate-create name_of_migration
```

Apply all pending migrations (up):
```
make migrate-up
```

Rollback the last migration (down):
```
make migrate-down
```

## Project Structure

- `cmd/main.go` — The main application entrypoint
- `cmd/script/main.go` — Migration runner script
- `db/migrations/` — Directory containing SQL migration files
- `Makefile` — Defines build, run, and migration commands
