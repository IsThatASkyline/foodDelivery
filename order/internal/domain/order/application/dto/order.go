package dto

import "github.com/google/uuid"

type CreateOrder struct {
	Items []OrderItem
}

type OrderItem struct {
	ID       int64
	Quantity int64
}

type CreateOrderInDB struct {
	UUID       uuid.UUID
	TotalPrice int64
	Status     string
}
