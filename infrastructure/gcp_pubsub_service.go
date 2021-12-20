package infrastructure

import (
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
	"github.com/mkaiho/google-api-sample/adapter/gcppubsub"
)

/** GCP Pub/Sub Service **/
func NewGCPPubsubService(ctx context.Context, config gcppubsub.GCPPubsubConfig) (*gcpPubsubService, error) {
	if ctx == nil {
		return nil, errors.New("ctx is required")
	}
	opts := newGCPPubsubOption(ctx, config)
	client, err := pubsub.NewClient(ctx, config.ProjectID(), opts)
	if err != nil {
		return nil, err
	}

	return &gcpPubsubService{
		client: client,
	}, nil
}

type gcpPubsubService struct {
	client *pubsub.Client
}
