package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/mkaiho/google-api-sample/infrastructure"
)

func main() {
	ctx := context.Background()

	config, err := infrastructure.LoadGCPPubsubConfigEnv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	service, err := infrastructure.NewGCPPubsubService(ctx, config)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	publisher := infrastructure.NewGCPPubsubPublisher(service)
	topic := publisher.Topic(config.Topic())

	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello World!"),
	})
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("ID: %s\n", id)
}
