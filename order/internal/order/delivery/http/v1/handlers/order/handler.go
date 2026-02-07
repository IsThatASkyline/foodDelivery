package order

import (
	"context"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/application/dto"
)

type orderUseCase interface {
	CreateOrder(ctx context.Context) error
	GetMenuItems(ctx context.Context) ([]dto.MenuItem, error)
}

type OrderHandler struct {
	orderUseCase orderUseCase
}

func NewOrderHandler(orderUseCase orderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: orderUseCase,
	}
}
