package usecase

import (
	"context"
	"fmt"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/application/dto"
)

func (uc *OrderUseCase) GetMenuItems(ctx context.Context) ([]dto.MenuItem, error) {
	const op = "order.application.GetMenuItems"

	data, err := uc.storage.GetMenuItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	for i, item := range data {
		item.Price = FromKopecksToRubbles(item.Price)
		data[i] = item
	}
	return data, nil
}
