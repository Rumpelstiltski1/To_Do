package kafka

import (
	"To_Do/internal/models"
	"To_Do/internal/repository"
	"To_Do/pkg/logger"
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"strconv"
	"strings"
	"time"
)

type KafkaConsumer interface {
	Start(ctx context.Context) error
	Close() error
}

type Consumer struct {
	reader  Reader
	storage *repository.Storage
}

func NewConsumer(broker, topic, groupId string, storage *repository.Storage) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{broker},
		GroupID:        groupId,
		Topic:          topic,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: time.Second,
	})
	return &Consumer{reader: r, storage: storage}
}

func (c *Consumer) Start(ctx context.Context) error {
	logger.Logger.Info("Kafka consumer запущен")
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				logger.Logger.Info("Kafka consumer остановлен")
				return nil
			}
			logger.Logger.Error("Ошибка чтения Kafka-сообщения", "err", err)
			continue
		}
		eventStr := string(m.Value)
		logger.Logger.Info("Получено сообщение из Kafka", "topic", m.Topic, "key", string(m.Key), "value", eventStr)

		event, err := parseEvent(eventStr)
		if err != nil {
			logger.Logger.Error("Не удалось распарсить сообщение Kafka", "err", err)
			continue
		}
		if err := c.storage.SaveEvent(ctx, event); err != nil {
			logger.Logger.Error("Не удалось сохранить событие в БД", "err", err)
		}
	}

}

func (c *Consumer) Close() error {
	return c.reader.Close()
}

func parseEvent(s string) (models.TaskEvent, error) {
	event := models.TaskEvent{}

	fields := strings.Split(s, ",") // разделяем строку на массив типа {action=a, title=b и тд.}
	for _, fiels := range fields {
		parts := strings.SplitN(fiels, "=", 2) // разделяем action=a на {action,a}
		if len(parts) != 2 {
			continue
		}
		key, value := parts[0], parts[1]

		switch key {
		case "action":
			event.Action = value
		case "id":
			pr, err := strconv.Atoi(value)
			if err != nil {
				logger.Logger.Error("Не удалось обработать id", "err", err)
				continue
			}
			event.TaskID = pr
		case "title":
			clean := strings.Trim(value, `"`)
			event.Title = &clean
		case "content":
			clean := strings.Trim(value, `"`)
			event.Content = &clean
		case "status":
			b, err := strconv.ParseBool(value)
			if err != nil {
				return event, err
			}
			event.Status = &b
		default:
			return event, errors.New("неизвестное поле в событии")
		}
	}
	if event.Action == "" || event.TaskID == 0 {
		return event, fmt.Errorf("неполные данные события: %+v", event)
	}
	return event, nil
}
