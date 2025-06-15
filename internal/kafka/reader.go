package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Reader interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
	Close() error
}
