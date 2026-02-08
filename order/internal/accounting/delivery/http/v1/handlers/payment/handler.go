package payment

import (
	"context"
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/application/dto"
)

type paymentUseCase interface {
	MakePayment(ctx context.Context, payments dto.MakePayment) error
}

type PaymentHandler struct {
	paymentUseCase paymentUseCase
}

func NewPaymentHandler(paymentUseCase paymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: paymentUseCase,
	}
}
