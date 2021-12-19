package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mkaiho/google-api-sample/adapter/gbpapi"
	"github.com/mkaiho/google-api-sample/infrastructure"
)

func main() {
	ctx := context.Background()

	gbpConfig, err := infrastructure.LoadGBPConfigEnv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	notificationClient, err := infrastructure.NewGBPNotification(ctx, gbpConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	accountID, err := gbpapi.ParseAccountID(123456)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	settings, err := notificationClient.GetNotificationSetting(ctx, *accountID)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("%v", settings)
}
