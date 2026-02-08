package order

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/delivery/http/v1/handlers/common"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/delivery/http/v1/handlers/order/requests"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/delivery/http/v1/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	const op = "order.delivery.http.CreateOrder" //TODO: gowrap

	var req requests.CreateOrder
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.orderUseCase.CreateOrder(ctx, req.ToDTO())
	if err != nil {
		middleware.MapErrors(ctx, err) //TODO: middleware
		return
	}
	ctx.JSON(http.StatusCreated, common.Response{
		Status: http.StatusCreated,
		Data:   data,
	})
}
