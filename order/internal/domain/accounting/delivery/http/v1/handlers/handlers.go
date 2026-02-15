package handlers

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/accounting/delivery/http/v1/handlers/payment"
)

type Handlers struct {
	PaymentHandler *payment.PaymentHandler
}
