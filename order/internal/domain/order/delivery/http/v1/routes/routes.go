package client

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/delivery/http/v1/handlers"
	orderRoutes "github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/delivery/http/v1/routes/order"
	"github.com/gin-gonic/gin"
)

func Setup(
	routerGroup *gin.RouterGroup,
	handlers handlers.Handlers,
) {
	orderGroup := routerGroup.Group("/orders")
	orderRoutes.Setup(orderGroup, handlers.OrderHandler)
}
