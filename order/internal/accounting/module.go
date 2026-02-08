package accounting

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/adapters/http/payment_gateway"
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/adapters/postgres"
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/application/usecase"
	v1handlers "github.com/IsThatASkyline/foodDelivery/order/internal/accounting/delivery/http/v1/handlers"
	paymentV1Handlers "github.com/IsThatASkyline/foodDelivery/order/internal/accounting/delivery/http/v1/handlers/payment"
	v1routes "github.com/IsThatASkyline/foodDelivery/order/internal/accounting/delivery/http/v1/routes"
	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v5/pgxpool"
)

type module struct {
	v1handlers v1handlers.Handlers
}

func NewModule(db *pgxpool.Pool) *module {
	txManager := postgres.NewTransactionRepo(db)
	paymentStorage := postgres.NewStorage(db)

	paymentGatewayClient := payment_gateway.NewPaymentGatewayClient()

	paymentUseCase := usecase.NewPaymentUseCase(paymentStorage, txManager, paymentGatewayClient)
	paymentV1Handler := paymentV1Handlers.NewPaymentHandler(paymentUseCase)

	return &module{
		v1handlers: v1handlers.Handlers{
			PaymentHandler: paymentV1Handler,
		},
	}
}

func (m *module) RegisterRoutes(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	v1routes.Setup(v1, m.v1handlers)
}
