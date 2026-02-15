package usecase

import (
	"context"
	dto2 "github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/application/dto"
	"github.com/google/uuid"
)

type TxManager interface {
	PerformTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type Storage interface {
	CreateOrder(ctx context.Context, in dto2.CreateOrderInDB) error
	GetOrderByID(ctx context.Context, id uuid.UUID) (dto2.Order, error)
	GetMenuItems(ctx context.Context) ([]dto2.MenuItem, error)
	AddItemsToOrder(ctx context.Context, in dto2.AddItemsToOrder) error
}
