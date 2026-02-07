package usecase

import "github.com/IsThatASkyline/foodDelivery/order/internal/order/application/interface"

type OrderUseCase struct {
	storage _interface.Storage
}

func NewOrderUseCase(storage _interface.Storage) *OrderUseCase {
	return &OrderUseCase{
		storage: storage,
	}
}
