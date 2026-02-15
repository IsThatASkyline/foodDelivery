package postgres

import (
	"context"
	"fmt"
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/application/dto"
	"github.com/google/uuid"
)

func (s *storage) GetOrderByID(ctx context.Context, id uuid.UUID) (dto.Order, error) {
	const op = "order.adapters.postgres.GetOrderByID"

	query :=
		`
	SELECT o.id,
	       o.total_price,
	       o.status,
	       mi.id,
	       mi.name,
	       oi.quantity,
		   oi.total_price
	FROM orders.orders o
	LEFT JOIN orders.order_items oi ON oi.order_id = o.id
	LEFT JOIN orders.menu_items mi ON mi.id = oi.menu_item_id
	WHERE o.id = $1
`
	rows, err := s.getDB(ctx).Query(ctx, query, id)
	if err != nil {
		return dto.Order{}, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var resp dto.Order
	items := make([]dto.Item, 0)

	for rows.Next() {
		var item dto.Item

		err := rows.Scan(
			&resp.ID,
			&resp.TotalPrice,
			&resp.Status,
			&item.ID,
			&item.Name,
			&item.Quantity,
			&item.TotalPrice,
		)
		if err != nil {
			return dto.Order{}, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return dto.Order{}, err
	}

	resp.Items = items
	return resp, nil
}
