package main

import (
	"context"
	"log"

	"github.com/mkaiho/google-api-sample/infrastructure"
)

func main() {
	gbpCredentials, err := infrastructure.LoadGBPCredentialEnv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	_, err = infrastructure.NewGBPNotification(context.Background(), gbpCredentials)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
