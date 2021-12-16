package main

import (
	"context"
	"log"
	"os"

	"github.com/mkaiho/google-api-sample/infrastructure"
)

const (
	envKeyGoogleAPIKey = "GOOGLE_API_KEY"
)

func main() {
	gbpOptions := infrastructure.NewGBPNotificationOptions()
	gbpOptions.APIKey = os.Getenv(envKeyGoogleAPIKey)
	if len(gbpOptions.APIKey) == 0 {
		log.Fatal("GOOGLE_API_KEY is not set")
		return
	}
	_, err := infrastructure.NewGBPNotification(context.Background(), *gbpOptions)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
}
