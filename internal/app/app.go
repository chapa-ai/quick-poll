package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"quick-poll/config"
	"quick-poll/internal/handler"
	"quick-poll/internal/kafka"
	"quick-poll/internal/repository/pg"
	"quick-poll/internal/service"
	"quick-poll/pkg/logger"
	"strings"
	"syscall"
)

const migrationsPath = "migrations"

type App struct {
	cfg config.Config
	log logger.Logger
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	connDB, err := pg.ConnectDB(a.cfg.GetDbConfig().GetDsn())
	if err != nil {
		a.log.Errorf("connect to db failed: %v", err)
		return err
	}

	if err = pg.MigrateUp(a.cfg.DB, migrationsPath); err != nil {
		a.log.Errorf("migrations failed: %v", err)
		return err
	}
	kafkaBrokers := strings.Split(a.cfg.Kafka.Brokers, ",")
	broker := kafka.New(kafkaBrokers, a.cfg.Kafka.Topic)

	s := service.NewService(ctx, a.cfg, connDB, *broker, a.log)
	h := handler.NewHandler(s)

	//consumer := kafka.NewConsumer([]string{"kafka:9092"}, "polls", "vote-consumers", s.DB)
	consumer := kafka.NewConsumer([]string{a.cfg.Kafka.Brokers}, a.cfg.Kafka.Topic, a.cfg.Kafka.Group, s.DB, a.log)

	go func() {
		err = consumer.Consume(context.Background())
		if err != nil {
			a.log.Fatalf("consumer failed: %v", err)
		}
	}()

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		<-exit
		h.Shutdown(ctx)
		cancel()
	}()

	return h.Start(fmt.Sprintf(":%s", a.cfg.App.Port))
}

func New(cfg config.Config, log logger.Logger) *App {
	return &App{
		cfg: cfg,
		log: log,
	}

}
