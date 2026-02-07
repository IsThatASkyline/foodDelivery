package order

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/delivery/http/v1/handlers/order"
	"github.com/gin-gonic/gin"
)

func Setup(routerGroup *gin.RouterGroup, orderHandler *order.OrderHandler) {
	routerGroup.POST("", orderHandler.CreateOrder)

	routerGroup.GET("/menu", orderHandler.GetMenuItems)
}
