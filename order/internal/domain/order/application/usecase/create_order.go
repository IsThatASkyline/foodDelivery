package usecase

import (
	"context"
	"fmt"
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/application/consts"
	dto2 "github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/application/dto"
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/application/errors"
	"github.com/google/uuid"
)

func (uc *OrderUseCase) CreateOrder(ctx context.Context, order dto2.CreateOrder) (uuid.UUID, error) {
	const op = "order.application.CreateOrder"

	err := validateCreateOrder(order)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%s: %w", op, err)
	}

	orderID := uuid.New()
	err = uc.txManager.PerformTransaction(ctx, func(ctx context.Context) error {
		menuItems, err := uc.storage.GetMenuItems(ctx)
		if err != nil {
			return err
		}

		itemsInOrder, orderTotalPrice := calculatePrices(menuItems, order.Items)

		err = uc.storage.CreateOrder(ctx, dto2.CreateOrderInDB{
			UUID:       orderID,
			TotalPrice: orderTotalPrice,
			Status:     consts.OrderStatusCreated,
		})
		if err != nil {
			return err
		}

		err = uc.storage.AddItemsToOrder(ctx, dto2.AddItemsToOrder{
			OrderID: orderID,
			Items:   itemsInOrder,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%s: %w", op, err)
	}

	return orderID, nil
}

func validateCreateOrder(order dto2.CreateOrder) error {
	if len(order.Items) <= 0 {
		return errors.ErrItemsMustBeMoreThanZero
	}
	for _, item := range order.Items {
		if item.Quantity <= 0 {
			return errors.ErrItemsMustBeMoreThanZero
		}
	}
	return nil
}

func calculatePrices(pricesList []dto2.MenuItem, items []dto2.OrderItem) ([]dto2.ItemInOrder, int64) {
	pricesMap := make(map[int64]dto2.MenuItem, len(pricesList))
	for _, item := range pricesList {
		pricesMap[item.ID] = item
	}

	result := make([]dto2.ItemInOrder, len(items))
	var orderTotalPrice int64

	for i, item := range items {
		totalPrice := pricesMap[item.ID].Price * item.Quantity
		result[i] = dto2.ItemInOrder{
			ID:         item.ID,
			Quantity:   item.Quantity,
			TotalPrice: totalPrice,
		}
		orderTotalPrice += totalPrice
	}
	return result, orderTotalPrice
}
