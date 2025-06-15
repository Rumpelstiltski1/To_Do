package kafka

import "context"

type KafkaProducer interface {
	SendMessage(ctx context.Context, key, value []byte) error
	Close() error
}
