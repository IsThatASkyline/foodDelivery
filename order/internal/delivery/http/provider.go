package api

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/adapters/postgres"
	"github.com/IsThatASkyline/foodDelivery/order/internal/application"
	"github.com/IsThatASkyline/foodDelivery/order/internal/delivery/http/v1/handlers"
	"github.com/IsThatASkyline/foodDelivery/order/internal/delivery/http/v1/handlers/order"

	"github.com/jackc/pgx/v5/pgxpool"
)

type provider struct {
	handlers handlers.Handlers
}

func NewProvider(db *pgxpool.Pool) (*provider, error) {
	const op = "delivery.http.provider.NewProvider"

	orderStorage := postgres.NewStorage(db)
	orderUseCase := application.NewCreateOrderUseCase(orderStorage)
	orderHandler := order.NewOrderHandler(orderUseCase)

	return &provider{
		handlers: handlers.Handlers{
			OrderHandler: orderHandler,
		},
	}, nil
}

func (provider *provider) SetupDependencies(app *App) {
	app.handlers = provider.handlers
}
