package services

import (
	"context"
	"log"

	"github.com/yourusername/payment-microservice/src/internal/repository"
	"github.com/yourusername/payment-microservice/src/pkg/external"
	"github.com/yourusername/payment-microservice/src/pkg/idempotency"
	pb "github.com/yourusername/payment-microservice/src/internal/handlers"
)

type PaymentService struct {}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	idKey := idempotency.GenerateIdempotencyKey(req)
	if idempotency.IsIdempotent(idKey) {
		log.Println("Payment is idempotent")
		return &pb.PaymentResponse{Status: "SUCCESS"}, nil
	}
		// Validate payment
	if err := validatePayment(req); err!= nil {
		return nil, err
	}
		// Check external service
	client := external.NewExternalServiceClient()
	if err := client.CheckFunds(ctx, req.AccountID); err!= nil {
		return nil, err
	}
		// Save payment
	repo := repository.NewPaymentRepository()
	if err := repo.SavePayment(ctx, req); err!= nil {
		return nil, err
	}
		// Mark as idempotent
	idempotency.MarkAsIdempotent(idKey)
	return &pb.PaymentResponse{Status: "SUCCESS"}, nil
}

func validatePayment(req *pb.PaymentRequest) error {
	// Implement validation logic
	return nil
}