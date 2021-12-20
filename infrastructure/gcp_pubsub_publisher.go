package infrastructure

import (
	"context"

	"cloud.google.com/go/pubsub"
)

/** GCP Pub/Sub Publisher **/
func NewGCPPubsubPublisher(service *gcpPubsubService) *gcpPubsubPublisher {
	return &gcpPubsubPublisher{
		service: service,
	}
}

type gcpPubsubPublisher struct {
	service *gcpPubsubService
}

func (p *gcpPubsubPublisher) Topic(topicID string) *gcpPusubTopic {
	topic := p.service.client.Topic(topicID)
	return &gcpPusubTopic{
		topic: topic,
	}
}

/** GCP Pub/Sub Topic **/
type gcpPusubTopic struct {
	topic *pubsub.Topic
}

func (t *gcpPusubTopic) Publish(ctx context.Context, msg *pubsub.Message) *gcpPublishResult {
	return &gcpPublishResult{
		result: t.topic.Publish(ctx, msg),
	}
}

/** GCP Pub/Sub Topic **/
type gcpPublishResult struct {
	result *pubsub.PublishResult
}

func (r *gcpPublishResult) Get(ctx context.Context) (serverID string, err error) {
	return r.result.Get(ctx)
}
