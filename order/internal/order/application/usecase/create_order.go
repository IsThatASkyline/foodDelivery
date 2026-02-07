package usecase

import (
	"context"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/application/errors"
)

func (uc *OrderUseCase) CreateOrder(ctx context.Context) error {
	return errors.ErrEntityNotFound
}
