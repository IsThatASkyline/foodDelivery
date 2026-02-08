package dto

import "github.com/google/uuid"

type CreateOrder struct {
	Items []OrderItem `json:"items" binding:"required"`
}

type OrderItem struct {
	ID       int64 `json:"id"`
	Quantity int64 `json:"quantity"`
}

type CreateOrderInDB struct {
	UUID       uuid.UUID
	TotalPrice int64
	Status     string
}
