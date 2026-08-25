package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/yourusername/payment-microservice/src/internal/handlers"
	"google.golang.org/grpc"
)

func main() {
	// Create a gRPC server
	grpcServer := grpc.NewServer()
	paymentHandler := &handlers.PaymentHandler{}
	pb.RegisterPaymentServiceServer(grpcServer, paymentHandler)

	// Create a gRPC gateway
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := runtime.RegisterService(ctx, gwmux, "payment.PaymentService", paymentHandler)
	if err!= nil {
		log.Fatalf("failed to register service: %v", err)
	}

	// Serve the gRPC server and gateway
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	log.Println("Server started at :8080")
	if err := httpServer.ListenAndServe(); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}