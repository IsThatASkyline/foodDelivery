package application

type CreateOrderUseCase struct {
	storage Storage
}

func NewCreateOrderUseCase(storage Storage) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		storage: storage,
	}
}
