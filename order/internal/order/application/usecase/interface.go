package usecase

import (
	"context"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/application/dto"
)

type TxManager interface {
	PerformTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type Storage interface {
	CreateOrder(ctx context.Context, in dto.CreateOrderInDB) error
	GetMenuItems(ctx context.Context) ([]dto.MenuItem, error)
	AddItemsToOrder(ctx context.Context, in dto.AddItemsToOrder) error
}
