package client

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/accounting/delivery/http/v1/handlers"
	paymentRoutes "github.com/IsThatASkyline/foodDelivery/order/internal/domain/accounting/delivery/http/v1/routes/payment"
	"github.com/gin-gonic/gin"
)

func Setup(
	routerGroup *gin.RouterGroup,
	handlers handlers.Handlers,
) {
	paymentGroup := routerGroup.Group("/payments")
	paymentRoutes.Setup(paymentGroup, handlers.PaymentHandler)
}
