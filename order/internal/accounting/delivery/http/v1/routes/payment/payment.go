package payment

import (
	"github.com/IsThatASkyline/foodDelivery/order/internal/accounting/delivery/http/v1/handlers/payment"
	"github.com/gin-gonic/gin"
)

func Setup(routerGroup *gin.RouterGroup, paymentHandler *payment.PaymentHandler) {
	routerGroup.POST("", paymentHandler.MakePayment)
}
