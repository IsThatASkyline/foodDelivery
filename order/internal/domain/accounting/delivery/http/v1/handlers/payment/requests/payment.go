package requests

import "github.com/google/uuid"

type MakePaymentRequest struct {
	OrderID uuid.UUID `json:"order_id" binding:"required"`
}
