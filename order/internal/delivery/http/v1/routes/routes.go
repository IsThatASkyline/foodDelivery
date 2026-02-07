package client

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/delivery/http/v1/handlers"
	"github.com/IsThatASkyline/foodDelivery/order/internal/delivery/http/v1/routes/order"
	"github.com/gin-gonic/gin"
)

func Setup(
	routerGroup *gin.RouterGroup,
	handlers handlers.Handlers,
) {
	orderGroup := routerGroup.Group("/orders")
	order.Setup(orderGroup, handlers.OrderHandler)
}
