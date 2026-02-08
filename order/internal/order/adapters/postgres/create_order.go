package postgres

import (
	"context"
	"fmt"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/application/dto"
)

func (s *storage) CreateOrder(ctx context.Context, in dto.CreateOrderInDB) error {
	const op = "order.adapter.postgres.CreateOrder"

	query :=
		`
	INSERT INTO orders.orders (id, total_price, status)
	VALUES ($1, $2, $3)
`
	_, err := s.getDB(ctx).Exec(ctx, query, in.UUID, in.TotalPrice, in.Status)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
