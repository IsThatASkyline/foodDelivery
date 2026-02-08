package payment_gateway

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func (c *PaymentGatewayClient) MakePayment(ctx context.Context, orderID uuid.UUID) error {
	const op = "accounting.adapters.http.payment_gateway.MakePayment"

	rand.New(rand.NewSource(time.Now().UnixNano()))
	if rand.Intn(4) == 0 {
		return fmt.Errorf("%s: not enough money on balance", op)
	}
	return nil
}
