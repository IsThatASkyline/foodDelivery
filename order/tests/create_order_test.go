package tests

import (
	"context"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/adapters/postgres"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/application/dto"
	"github.com/IsThatASkyline/foodDelivery/order/internal/order/application/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderUseCase_CreateOrder(t *testing.T) {
	ctx := context.Background()

	db, cleanup, err := SetupTestDB(ctx)
	require.NoError(t, err)
	defer cleanup()

	storage := postgres.NewStorage(db)
	txManager := postgres.NewTransactionRepo(db)
	uc := usecase.NewOrderUseCase(storage, txManager)

	order := dto.CreateOrder{
		Items: []dto.OrderItem{
			{
				ID:       1,
				Quantity: 3,
			},
			{
				ID:       2,
				Quantity: 1,
			},
		},
	}

	orderID, err := uc.CreateOrder(ctx, order)

	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, orderID)
}
