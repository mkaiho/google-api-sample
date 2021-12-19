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

	gcpConfig, err := infrastructure.LoadGCPConfigEnv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	pub, err := infrastructure.NewGCPPublisher(ctx, gcpConfig.ProjectID(), *gcpConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	topic := pub.Topic(gcpConfig.PubsubTopic())
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello World!"),
	})

	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("ID: %s\n", id)
}
