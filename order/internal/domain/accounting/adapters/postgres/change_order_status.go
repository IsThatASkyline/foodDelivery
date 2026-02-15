package postgres

import (
	"context"
	"fmt"
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/accounting/application/dto"
)

func (s *storage) ChangeOrderStatus(ctx context.Context, in dto.ChangeOrderStatus) error {
	const op = "accounting.adapters.postgres..ChangeOrderStatus"

	query := `UPDATE orders.orders SET status = $2 WHERE id = $1`

	if _, err := s.getDB(ctx).Exec(ctx, query, in.ID, in.NewStatus); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
