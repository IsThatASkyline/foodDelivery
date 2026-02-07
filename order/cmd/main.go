package main

import (
	"context"
	"github.com/IsThatASkyline/foodDelivery/order/internal/config"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/adapters/postgres"
	server "github.com/IsThatASkyline/foodDelivery/order/internal/server/http"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("can't load config: %v", err)
	}

	db, err := postgres.NewDBPool(ctx, cfg.DB.ConnectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	defer db.Close()
	log.Println("db Pool has been initialized")

	connForMigrations := stdlib.OpenDBFromPool(db)
	if err = goose.Up(connForMigrations, cfg.DB.MigrationPath); err != nil {
		log.Fatalf("failed to run migrations: %s", err)
	}

	// Создаем приложение
	app := server.NewApp()

	// Устанавливаем все роуты
	app.SetupRoutes(
		order.NewModule(db),
	)

	// Запускаем сервер
	app.StartServe(ctx)
}
