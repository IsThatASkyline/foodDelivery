package application

import (
	"context"
	"github.com/IsThatASkyline/foodDelivery/order/internal/errors"
)

func (uc *CreateOrderUseCase) CreateOrder(ctx context.Context) error {
	return errors.ErrEntityNotFound
}
