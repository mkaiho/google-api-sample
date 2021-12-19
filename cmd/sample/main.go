package main

import (
	"context"
	"log"

	"github.com/mkaiho/google-api-sample/infrastructure"
)

func main() {
	ctx := context.Background()

	gbpCredentials, err := infrastructure.LoadGBPCredentialEnv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	_, err = infrastructure.NewGBPNotification(ctx, gbpCredentials)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
