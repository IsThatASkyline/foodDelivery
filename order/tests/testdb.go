package tests

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func SetupTestDB(ctx context.Context) (*pgxpool.Pool, func(), error) {
	var (
		dbname   = "dbname"
		user     = "user"
		password = "pass"
	)

	// 1. Start the postgres ctr and run any migrations on it
	container, err := postgres.Run(
		ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbname),
		postgres.WithUsername(user),
		postgres.WithPassword(password),
		postgres.BasicWaitStrategies(),
		postgres.WithSQLDriver("pgx"),
	)
	if err != nil {
		return nil, nil, err
	}

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, nil, err
	}

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, nil, err
	}

	ctxPing, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := pool.Ping(ctxPing); err != nil {
		return nil, nil, err
	}

	if err := applyGooseMigrations(connStr); err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		pool.Close()
		_ = container.Terminate(ctx)
	}

	return pool, cleanup, nil
}

func applyGooseMigrations(connStr string) error {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	root, err := os.Getwd()
	if err != nil {
		return err
	}

	migrationsPath := filepath.Join(root, "..", "migrations")

	return goose.Up(db, migrationsPath)
}
