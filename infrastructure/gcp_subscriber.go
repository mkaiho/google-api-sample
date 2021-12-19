package infrastructure

import (
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
)

/** TODO: Refactoring **/

/** GCP Pub/Sub Subscriber **/
func NewGCPSubscriber(ctx context.Context, projectID string, config gcpConfig) (*gcpSubscriber, error) {
	if ctx == nil {
		return nil, errors.New("ctx is required")
	}
	opts := newGCPOption(ctx, config)
	client, err := pubsub.NewClient(ctx, projectID, opts)
	if err != nil {
		return nil, err
	}

	return &gcpSubscriber{
		client: client,
	}, nil
}

type gcpSubscriber struct {
	client *pubsub.Client
}

func (p *gcpSubscriber) Subscription(subscriptionID string) gcpSubscription {
	subscription := p.client.Subscription(subscriptionID)
	return gcpSubscription{
		subscription: subscription,
	}
}

/** GCP Pub/Sub Subscription **/
type gcpSubscription struct {
	subscription *pubsub.Subscription
}

func (s *gcpSubscription) Receive(ctx context.Context, f func(context.Context, *pubsub.Message)) error {
	return s.subscription.Receive(ctx, f)
}
