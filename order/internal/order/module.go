package order

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/adapters/postgres"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/application/usecase"
	v1handlers "github.com/IsThatASkyline/foodDelivery/order/internal/order/delivery/http/v1/handlers"
	orderV1Handlers "github.com/IsThatASkyline/foodDelivery/order/internal/order/delivery/http/v1/handlers/order"
	v1routes "github.com/IsThatASkyline/foodDelivery/order/internal/order/delivery/http/v1/routes"
	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v5/pgxpool"
)

type module struct {
	v1handlers v1handlers.Handlers
}

func NewModule(db *pgxpool.Pool) *module {
	orderStorage := postgres.NewStorage(db)
	orderUseCase := usecase.NewOrderUseCase(orderStorage)
	orderV1Handler := orderV1Handlers.NewOrderHandler(orderUseCase)

	return &module{
		v1handlers: v1handlers.Handlers{
			OrderHandler: orderV1Handler,
		},
	}
}

func (m *module) RegisterRoutes(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	v1routes.Setup(v1, m.v1handlers)
}
