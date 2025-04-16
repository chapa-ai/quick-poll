package pg

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "log/slog"
	"quick-poll/config"
)

func MigrateUp(cfg config.DB, migrationsPath string) error {
	dsn := cfg.GetMigrateDsn()

	m, err := migrate.New(fmt.Sprintf("file://%s", migrationsPath), dsn)
	if err != nil {
		return err
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Error(fmt.Sprintf("Migrate: up error: %s", err))
		return err
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Info("Migrate: no change")
		return nil
	}

	log.Info("Migrate: up success")
	return nil
}
