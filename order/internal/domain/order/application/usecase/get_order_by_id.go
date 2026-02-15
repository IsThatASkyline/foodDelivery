package usecase

import (
	"context"
	"fmt"
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/application/dto"
	"github.com/google/uuid"
)

func (uc *OrderUseCase) GetOrderByID(ctx context.Context, id uuid.UUID) (dto.Order, error) {
	const op = "order.application.GetOrderByID"

	data, err := uc.storage.GetOrderByID(ctx, id)
	if err != nil {
		return dto.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	return convertPricesInRubbles(data), nil
}

func convertPricesInRubbles(in dto.Order) dto.Order {
	in.TotalPrice = FromKopecksToRubbles(in.TotalPrice)
	for i, item := range in.Items {
		item.TotalPrice = FromKopecksToRubbles(item.TotalPrice)
		in.Items[i] = item
	}
	return in
}
