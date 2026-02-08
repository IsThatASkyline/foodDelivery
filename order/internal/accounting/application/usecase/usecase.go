package usecase

type PaymentUseCase struct {
	storage        Storage
	txManager      TxManager
	paymentGateway PaymentGateway
}

func NewPaymentUseCase(storage Storage, txManager TxManager, paymentGateway PaymentGateway) *PaymentUseCase {
	return &PaymentUseCase{
		storage:        storage,
		txManager:      txManager,
		paymentGateway: paymentGateway,
	}
}
