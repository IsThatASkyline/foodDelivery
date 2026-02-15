package postgres

import (
	"context"
	"fmt"
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/application/dto"
	"github.com/jackc/pgx/v5"
)

func (s *storage) AddItemsToOrder(ctx context.Context, in dto.AddItemsToOrder) error {
	const op = "order.adapters.postgres.AddItemsToOrder"

	query :=
		`
INSERT INTO orders.order_items (order_id, menu_item_id, quantity, total_price)
VALUES ($1, $2, $3, $4)
`
	var batch pgx.Batch
	for _, item := range in.Items {
		batch.Queue(query, in.OrderID, item.ID, item.Quantity, item.TotalPrice)
	}

	res := s.getDB(ctx).SendBatch(ctx, &batch)
	defer func() { _ = res.Close() }()

	for range in.Items {
		if _, err := res.Exec(); err != nil {
			return fmt.Errorf("%s: failed to send batch: %w", op, err)
		}
	}

	return nil
}
