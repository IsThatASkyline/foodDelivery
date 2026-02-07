package application

import "context"

type Storage interface {
	CreateOrder(ctx context.Context)
}
