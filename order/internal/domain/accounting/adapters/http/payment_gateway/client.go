package payment_gateway

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type PaymentGatewayClient struct {
	client *http.Client
}

func NewPaymentGatewayClient() *PaymentGatewayClient {
	timeout := time.Second * 10
	netDialer := &net.Dialer{
		Timeout:   timeout,
		KeepAlive: 10 * time.Minute,
	}
	transport := &http.Transport{
		TLSHandshakeTimeout:   time.Minute,
		MaxConnsPerHost:       10,
		MaxIdleConnsPerHost:   5,
		ResponseHeaderTimeout: timeout,
		DialContext:           netDialer.DialContext,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
	return &PaymentGatewayClient{
		client: &http.Client{
			Transport: transport,
			Timeout:   timeout,
		},
	}
}
