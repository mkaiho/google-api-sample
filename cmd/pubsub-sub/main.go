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

	config, err := infrastructure.LoadGCPPubsubConfigEnv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	service, err := infrastructure.NewGCPPubsubService(ctx, config)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	sub := infrastructure.NewGCPPubsubSubscriber(service)
	subscription := sub.Subscription(config.Subscription())
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
