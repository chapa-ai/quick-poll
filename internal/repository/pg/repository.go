package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"quick-poll/internal/repository"
	"quick-poll/pkg/logger"
	"sync"
)

var (
	once        sync.Once
	pool        *pgxpool.Pool
	connPoolErr error
)

type Repository struct {
	logger logger.Logger
	db     *pgxpool.Pool
}

func ConnectDB(dsn string) (*pgxpool.Pool, error) {
	once.Do(func() {
		cfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			connPoolErr = err
			return
		}

		pool, err = pgxpool.ConnectConfig(context.Background(), cfg)
		if err != nil {
			connPoolErr = err
			return
		}

		if err = pool.Ping(context.Background()); err != nil {
			connPoolErr = err
			pool.Close()
			return
		}
	})
	return pool, connPoolErr
}

func NewCounterRepository(log logger.Logger, connDB *pgxpool.Pool) repository.IRepository {
	return &Repository{
		logger: log,
		db:     connDB,
	}
}
