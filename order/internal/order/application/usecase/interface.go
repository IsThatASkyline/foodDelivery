package usecase

import (
	"context"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/application/dto"
	"github.com/google/uuid"
)

type TxManager interface {
	PerformTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type Storage interface {
	CreateOrder(ctx context.Context, in dto.CreateOrderInDB) error
	GetOrderByID(ctx context.Context, id uuid.UUID) (dto.Order, error)
	GetMenuItems(ctx context.Context) ([]dto.MenuItem, error)
	AddItemsToOrder(ctx context.Context, in dto.AddItemsToOrder) error
}
