package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/mkaiho/google-api-sample/infrastructure"
)

func main() {
	var mu sync.Mutex
	ctx := context.Background()

	gcpConfig, err := infrastructure.LoadGCPConfigEnv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	sub, err := infrastructure.NewGCPSubscriber(ctx, gcpConfig.ProjectID(), *gcpConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	subscription := sub.Subscription(gcpConfig.PubsubSubscription())
	cctx, cancel := context.WithCancel(ctx)

	fmt.Println("Listening...")
	err = subscription.Receive(cctx, func(c context.Context, m *pubsub.Message) {
		defer fmt.Println("End receiving")
		m.Ack()
		fmt.Printf("Got message: %q\n", string(m.Data))
		mu.Lock()
		defer mu.Unlock()
		cancel()
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
