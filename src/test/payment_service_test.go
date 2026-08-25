package test

import (
	"testing"
	"context"

	"github.com/yourusername/payment-microservice/src/internal/services"
	pb "github.com/yourusername/payment-microservice/src/internal/handlers"
)

func TestProcessPayment(t *testing.T) {
	service := services.NewPaymentService()
	req := &pb.PaymentRequest{
		AccountID: "1234567890",
		Amount:     100.0,
	}
	_, err := service.ProcessPayment(context.Background(), req)
	if err!= nil {
		 t.Errorf("ProcessPayment failed: %v", err)
	}
}