package application

import (
	"context"
	"encoding/json"
	"fmt"
)

type Offset int

const (
	// Beginning starts consuming from the earliest available message
	Beginning Offset = iota
	// Latest starts consuming from the most recent message
	Latest
	// Custom allows specifying a particular offset value
	Custom
)

// SubscriptionOptions contains configuration for message subscription
type SubscriptionOptions struct {
	Offset      Offset
	CustomValue int64
}

type MessageHandler[T any] func(message T) error

type MessagingPort interface {
	PublishMessage(ctx context.Context, topic string, message interface{}) error

	// Subscribe registers a handler for messages on a topic
	// The handler will receive messages deserialized into the appropriate type
	Subscribe(ctx context.Context, topic string, groupID string, options SubscriptionOptions, handler interface{}) error
}

func DebugString(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprint("Error marshaling: %v", err)
	}

	return string(bytes)
}
