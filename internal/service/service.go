package service

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"quick-poll/config"
	"quick-poll/internal/kafka"
	"quick-poll/internal/repository"
	"quick-poll/internal/repository/pg"
	"quick-poll/pkg/logger"
)

type Service struct {
	Ctx    context.Context
	Cfg    config.Config
	DB     repository.IRepository
	Broker kafka.KafkaProducer
	Logger *logger.Logger
}

func NewService(ctx context.Context, cfg config.Config, connDB *pgxpool.Pool, broker kafka.KafkaProducer, log logger.Logger) *Service {
	return &Service{
		Ctx:    ctx,
		Cfg:    cfg,
		DB:     pg.NewCounterRepository(log, connDB),
		Broker: broker,
		Logger: &log,
	}
}
