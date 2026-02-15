package handlers

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/delivery/http/v1/handlers/order"
)

type Handlers struct {
	OrderHandler *order.OrderHandler
}
