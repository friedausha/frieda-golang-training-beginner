package aws

import (
	"context"
	"frieda-golang-training-beginner/domain"
)

type MessageClient interface {
	// Creates a new long polling queue and returns its URL.
	CreateQueue(ctx context.Context, queueName string, isDLX bool) (string, error)
	// Send a message to queue and returns its message ID.
	Send(ctx context.Context, req *domain.SendRequest) (string, error)
}