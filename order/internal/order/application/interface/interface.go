package _interface

import (
	"context"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/application/dto"
)

type Storage interface {
	CreateOrder(ctx context.Context)
	GetMenuItems(ctx context.Context) ([]dto.MenuItem, error)
}
