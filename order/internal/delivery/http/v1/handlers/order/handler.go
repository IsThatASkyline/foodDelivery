package order

import (
	"context"
)

type orderUseCase interface {
	CreateOrder(ctx context.Context) error
}

type OrderHandler struct {
	orderUseCase orderUseCase
}

func NewOrderHandler(orderUseCase orderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: orderUseCase,
	}
}
