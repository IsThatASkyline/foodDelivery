package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/application/dto"
	apperrors "github.com/IsThatASkyline/foodDelivery/order/internal/accounting/application/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *storage) GetOrderByID(ctx context.Context, id uuid.UUID) (dto.Order, error) {
	const op = "accounting.adapters.postgres.GetOrderByID"

	query :=
		`
	SELECT o.id,
	       o.total_price,
	       o.status
	FROM orders.orders o
	WHERE o.id = $1
`
	var resp dto.Order
	err := s.getDB(ctx).QueryRow(ctx, query, id).Scan(
		&resp.ID,
		&resp.TotalPrice,
		&resp.Status,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.Order{}, fmt.Errorf("%s: %w", op, apperrors.ErrEntityNotFound)
		}
		return dto.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}
