package external

import (
	"context"
	"log"
)

type ExternalServiceClient struct {}

func NewExternalServiceClient() *ExternalServiceClient {
	return &ExternalServiceClient{}
}

func (c *ExternalServiceClient) CheckFunds(ctx context.Context, accountID string) error {
	log.Println("Checking funds")
	// Implement fund check logic
	return nil
}