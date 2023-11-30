package pg

import (
	"context"
	"gophermart/internal/config"
	"gophermart/internal/logger"
	"time"
)

func (p Pg) Up(ctx context.Context) error {
	logger.Log.Info("Start storage up")

	ctxT, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := p.migrationUp(ctxT); err != nil {
		return err
	}

	logger.Log.Info("Stop storage up")

	return nil
}

func (p Pg) Down(_ context.Context) error {
	logger.Log.Info("Storage down")

	if config.Instance.DownMigrations() {
		defer p.migrationDown()
	}

	defer p.db.Close()

	return nil
}
