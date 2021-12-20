package infrastructure

import (
	"context"

	"cloud.google.com/go/pubsub"
)

/** GCP Pub/Sub Subscriber **/
func NewGCPPubsubSubscriber(service *gcpPubsubService) *gcpPubsubSubscriber {
	return &gcpPubsubSubscriber{
		service: service,
	}
}

type gcpPubsubSubscriber struct {
	service *gcpPubsubService
}

func (p *gcpPubsubSubscriber) Subscription(subscriptionID string) *gcpPubsubSubscription {
	return &gcpPubsubSubscription{
		subscription: p.service.client.Subscription(subscriptionID),
	}
}

/** GCP Pub/Sub Subscription **/
type gcpPubsubSubscription struct {
	subscription *pubsub.Subscription
}

func (s *gcpPubsubSubscription) Receive(ctx context.Context, f func(context.Context, *pubsub.Message)) error {
	return s.subscription.Receive(ctx, f)
}
