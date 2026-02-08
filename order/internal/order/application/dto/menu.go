package dto

import "github.com/google/uuid"

type MenuItem struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type AddItemsToOrder struct {
	OrderID uuid.UUID
	Items   []ItemInOrder
}

type ItemInOrder struct {
	ID         int64
	Quantity   int64
	TotalPrice int64
}
