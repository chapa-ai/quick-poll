package repository

import (
	"context"
	"quick-poll/internal/models"
)

type IRepository interface {
	Create(ctx context.Context, question string, options []string) (*models.Poll, error)
	Vote(ctx context.Context, pollID string, option string) error
	GetByID(ctx context.Context, pollID string) (*models.Poll, error)
}

type KafkaProducer interface {
	Publish(ctx context.Context, topic string, key, value []byte) error
}
