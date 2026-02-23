package order

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/delivery/http/v1/handlers/common"
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/delivery/http/v1/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *OrderHandler) GetOrderByID(ctx *gin.Context) {
	const op = "order.delivery.http.GetOrderByID" //TODO: gowrap

	var req struct {
		ID string `uri:"id"`
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.orderUseCase.GetOrderByID(ctx, id)
	if err != nil {
		middleware.MapErrors(ctx, err) //TODO: middleware
		return
	}
	data.PaymentMethod = "cash"
	ctx.JSON(http.StatusOK, common.Response{
		Status: http.StatusOK,
		Data:   data,
	})
}
