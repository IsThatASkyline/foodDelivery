package main

import (
	"context"
	"github.com/IsThatASkyline/foodDelivery/order/internal/adapters/postgres"
	"github.com/IsThatASkyline/foodDelivery/order/internal/config"
	api "github.com/IsThatASkyline/foodDelivery/order/internal/delivery/http"
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

	//connForMigrations := stdlib.OpenDBFromPool(incomeDB)
	//if err = goose.Up(connForMigrations, cfg.IncomeDB.MigrationPath); err != nil {
	//	log.Fatalf("failed to run migrations: %s", err)
	//}

	// Создаем приложение
	app := api.NewApp()

	// Прокидываем зависимости
	provider, err := api.NewProvider(db)
	if err != nil {
		log.Fatalf("failed to initialize provider: %s", err)
	}
	provider.SetupDependencies(app)

	// Устанавливаем все роуты
	app.SetupRoutes()

	// Запускаем сервер
	app.StartServe(ctx)
}
