package payment

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/application/dto"
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/delivery/http/v1/handlers/common"
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/delivery/http/v1/handlers/payment/requests"
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/delivery/http/v1/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *PaymentHandler) MakePayment(ctx *gin.Context) {
	const op = "accounting.delivery.http.MakePayment" //TODO: gowrap

	var req requests.MakePaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.paymentUseCase.MakePayment(ctx, dto.MakePayment{
		OrderID: req.OrderID,
	})
	if err != nil {
		middleware.MapErrors(ctx, err) //TODO: middleware
		return
	}
	ctx.JSON(http.StatusCreated, common.Response{
		Status: http.StatusCreated,
	})
}
