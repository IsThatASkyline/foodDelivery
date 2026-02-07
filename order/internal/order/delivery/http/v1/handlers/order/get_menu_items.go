package order

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/delivery/http/v1/handlers/common"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/delivery/http/v1/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *OrderHandler) GetMenuItems(ctx *gin.Context) {
	const op = "order.delivery.http.GetMenuItems" //TODO: gowrap

	data, err := h.orderUseCase.GetMenuItems(ctx)
	if err != nil {
		middleware.MapErrors(ctx, err) //TODO: middleware
		return
	}
	ctx.JSON(http.StatusOK, common.Response{
		Status: http.StatusOK,
		Data:   data,
	})
}
