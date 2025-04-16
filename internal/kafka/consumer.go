package kafka

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/segmentio/kafka-go"
	"quick-poll/internal/models"
	"quick-poll/internal/repository"
	"quick-poll/pkg/logger"
)

type KafkaConsumer struct {
	reader *kafka.Reader
	repo   repository.IRepository
	logger *logger.Logger
}

func NewConsumer(brokers []string, topic string, groupID string, repo repository.IRepository, logger logger.Logger) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:    topic,
			GroupID:  groupID,
			MinBytes: 10e3,
			MaxBytes: 10e6,
		}),
		repo:   repo,
		logger: &logger,
	}
}

func (kc *KafkaConsumer) Consume(ctx context.Context) error {
	for {
		msg, err := kc.reader.ReadMessage(ctx)
		if err != nil {
			kc.logger.Errorf("failed reading message: %v", err)
			return fmt.Errorf("failed to read message: %w", err)
		}

		var vote models.Poll
		if err = json.Unmarshal(msg.Value, &vote); err != nil {
			kc.logger.Errorf("failed unmarshalling: %v", err)
			continue
		}

		for option := range vote.Options {
			if err = kc.repo.Vote(ctx, vote.ID, option); err != nil {
				kc.logger.Errorf("failed make a vote: %v", err)
			}
			break
		}
	}
}
