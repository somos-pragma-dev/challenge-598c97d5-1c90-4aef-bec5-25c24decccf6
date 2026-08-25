package handlers

import (
	"context"
	"log"

	pb "github.com/yourusername/payment-microservice/src/internal/handlers"
	"github.com/yourusername/payment-microservice/src/internal/services"
)

type PaymentHandler struct {}

func (h *PaymentHandler) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Println("Processing payment")
	service := services.NewPaymentService()
	response, err := service.ProcessPayment(ctx, req)
	if err!= nil {
		return nil, err
	}
	return response, nil
}