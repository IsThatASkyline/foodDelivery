package usecase

type OrderUseCase struct {
	storage   Storage
	txManager TxManager
}

func NewOrderUseCase(storage Storage, txManager TxManager) *OrderUseCase {
	return &OrderUseCase{
		storage:   storage,
		txManager: txManager,
	}
}
