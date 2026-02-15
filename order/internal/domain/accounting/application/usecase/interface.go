package usecase

import (
	"context"
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/accounting/application/dto"
	"github.com/google/uuid"
)

type TxManager interface {
	PerformTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type Storage interface {
	GetOrderByID(ctx context.Context, id uuid.UUID) (dto.Order, error)
	ChangeOrderStatus(ctx context.Context, in dto.ChangeOrderStatus) error
}

type PaymentGateway interface {
	MakePayment(ctx context.Context, orderID uuid.UUID) error
}
