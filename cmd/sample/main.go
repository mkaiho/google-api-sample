package main

import (
	"context"
	"log"

	"github.com/mkaiho/google-api-sample/infrastructure"
)

func main() {
	ctx := context.Background()

	gbpConfig, err := infrastructure.LoadGBPConfigEnv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	_, err = infrastructure.NewGBPNotification(ctx, gbpConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
