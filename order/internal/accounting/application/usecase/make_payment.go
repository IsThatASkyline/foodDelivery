package usecase

import (
	"context"
	"fmt"
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/application/consts"
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/application/dto"
	apperrors "github.com/IsThatASkyline/foodDelivery/order/internal/accounting/application/errors"
)

func (uc *PaymentUseCase) MakePayment(ctx context.Context, payment dto.MakePayment) error {
	const op = "accounting.application.MakePayment"

	err := uc.txManager.PerformTransaction(ctx, func(ctx context.Context) error {
		order, err := uc.storage.GetOrderByID(ctx, payment.OrderID)
		if err != nil {
			return err
		}
		if order.Status != consts.OrderStatusCreated {
			return apperrors.ErrCannotChangeOrderInThisStatus
		}
		err = uc.paymentGateway.MakePayment(ctx, order.ID)
		if err != nil {
			return fmt.Errorf("%w: %w", apperrors.ErrCannotMakePayment, err)
		}

		err = uc.storage.ChangeOrderStatus(ctx, dto.ChangeOrderStatus{
			ID:        payment.OrderID,
			NewStatus: consts.OrderStatusPaid,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
