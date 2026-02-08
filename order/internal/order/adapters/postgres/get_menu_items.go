package postgres

import (
	"context"
	"fmt"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/application/dto"
)

func (s *storage) GetMenuItems(ctx context.Context) ([]dto.MenuItem, error) {
	const op = "order.adapters.postgres.GetMenuItems"

	query :=
		`
	SELECT id, name, price FROM orders.menu_items
`
	rows, err := s.getDB(ctx).Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	resp := make([]dto.MenuItem, 0)
	for rows.Next() {
		var item dto.MenuItem
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		resp = append(resp, item)
		if rows.Err() != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return resp, nil
}
