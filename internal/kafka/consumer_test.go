package kafka

import (
	"To_Do/internal/models"
	"To_Do/internal/repository"
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockReader struct {
	messages []kafka.Message
	index    int
}

func (r *mockReader) ReadMessage(ctx context.Context) (kafka.Message, error) {
	if r.index >= len(r.messages) {
		time.Sleep(10 * time.Millisecond)
		return kafka.Message{}, ctx.Err()
	}
	msg := r.messages[r.index]
	r.index++
	return msg, nil
}

func (r *mockReader) Close() error {
	return nil
}

func TestConsumer_Start(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	kafkaMsg := `action=create,id=1,title="Learn Go",content="Write unit tests",status=true`

	reader := &mockReader{messages: []kafka.Message{{Value: []byte(kafkaMsg)}}}

	expected := models.TaskEvent{
		TaskID:  1,
		Action:  "create",
		Title:   strPtr("Learn Go"),
		Content: strPtr("Write unit tests"),
		Status:  boolPtr(true),
	}
	mockStorage := new(repository.MockStorage)
	mockStorage.On("SaveEvent", mock.Anything, expected).Return(nil)

	consumer := &Consumer{
		reader:  reader,
		storage: &repository.Storage{EventRepository: mockStorage},
	}

	err := consumer.Start(ctx)
	if err != nil {
		t.Fatalf("ошибка запуска consumer: %v", err)
	}
	mockStorage.AssertCalled(t, "SaveEvent", mock.Anything, expected)
	mockStorage.AssertNumberOfCalls(t, "SaveEvent", 1)
}
func strPtr(s string) *string { return &s }
func boolPtr(b bool) *bool    { return &b }
