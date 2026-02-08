package dto

import "github.com/google/uuid"

type MakePayment struct {
	OrderID uuid.UUID
}

type ChangeOrderStatus struct {
	ID        uuid.UUID
	NewStatus string
}

type Order struct {
	ID         uuid.UUID
	TotalPrice int64
	Status     string
}
