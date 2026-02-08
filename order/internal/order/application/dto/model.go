package dto

import (
	"github.com/google/uuid"
)

type Order struct {
	ID            uuid.UUID
	Items         []Item
	TotalPrice    int64
	Status        string
	PaymentMethod string
	Customer      Customer
	Delivery      Delivery
}

type Item struct {
	ID         int64
	Name       string
	Quantity   int
	TotalPrice int64
}

type Customer struct {
	ID                    int64
	CustomerContactNumber string
	CustomerAddress       string
}

type Delivery struct {
	Courier         Courier
	DeliveryTime    string
	DeliveryAddress string
}

type Courier struct {
	ID            int64
	Name          string
	ContactNumber string
}
