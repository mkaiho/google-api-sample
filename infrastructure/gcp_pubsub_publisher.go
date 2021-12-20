package infrastructure

import (
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
)

/** TODO: Refactoring **/

/** GCP Pub/Sub Publisher **/
func NewGCPPublisher(ctx context.Context, projectID string, config gcpConfig) (*gcpPublisher, error) {
	if ctx == nil {
		return nil, errors.New("ctx is required")
	}
	opts := newGCPOption(ctx, config)
	client, err := pubsub.NewClient(ctx, projectID, opts)
	if err != nil {
		return nil, err
	}

	return &gcpPublisher{
		client: client,
	}, nil
}

type gcpPublisher struct {
	client *pubsub.Client
}

func (p *gcpPublisher) Topic(topicID string) gcpTopic {
	topic := p.client.Topic(topicID)
	return gcpTopic{
		topic: topic,
	}
}

/** GCP Pub/Sub Topic **/
type gcpTopic struct {
	topic *pubsub.Topic
}

func (t *gcpTopic) Publish(ctx context.Context, msg *pubsub.Message) gcpPublishResult {
	return gcpPublishResult{
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
