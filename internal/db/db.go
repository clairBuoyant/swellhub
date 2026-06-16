// Package db owns the Postgres connection pool and schema migrations.
package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib" // registers the "pgx" database/sql driver for goose
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

// NewPool opens a pgx connection pool and verifies connectivity.
func NewPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("connect postgres: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}
	return pool, nil
}

// RunGoose runs a goose migration command ("up", "down", "status", "reset", …)
// against the given DSN using the embedded migrations.
func RunGoose(dsn, command string) error {
	goose.SetBaseFS(migrationsFS)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer sqlDB.Close()

	if err := goose.RunContext(context.Background(), command, sqlDB, "migrations"); err != nil {
		return fmt.Errorf("goose %s: %w", command, err)
	}
	return nil
}

// Migrate applies all outstanding up migrations against the given DSN.
func Migrate(dsn string) error { return RunGoose(dsn, "up") }
