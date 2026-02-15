package requests

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/domain/order/application/dto"
)

type CreateOrder struct {
	Items []OrderItem `json:"items" binding:"required,dive"`
}

type OrderItem struct {
	ID       int64 `json:"id" binding:"required"`
	Quantity int64 `json:"quantity" binding:"required"`
}

func (r *CreateOrder) ToDTO() dto.CreateOrder {
	orderItems := make([]dto.OrderItem, len(r.Items))
	for i, item := range r.Items {
		orderItems[i] = dto.OrderItem{
			ID:       item.ID,
			Quantity: item.Quantity,
		}
	}
	return dto.CreateOrder{
		Items: orderItems,
	}
}
