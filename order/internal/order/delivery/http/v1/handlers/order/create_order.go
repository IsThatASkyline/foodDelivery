package order

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/delivery/http/v1/middleware"
	"github.com/gin-gonic/gin"
)

func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	const op = "order/create_order" //TODO: gowrap

	err := h.orderUseCase.CreateOrder(ctx)
	if err != nil {
		middleware.MapErrors(ctx, err) //TODO: middleware
		return
	}
}
