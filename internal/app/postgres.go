package app

import (
	"context"
	"time"

	"github.com/avast/retry-go"
	"github.com/pkg/errors"
	"github.com/slowhigh/Umm/pkg/postgres"
	"github.com/slowhigh/Umm/pkg/utils"
)

func (a *app) connectPostgres(ctx context.Context) error {
	retryOption := []retry.Option{
		retry.Attempts(a.cfg.Timeouts.PostgresInitRetryCount),
		retry.Delay(time.Duration(a.cfg.Timeouts.PostgresInitMilliseconds) * time.Millisecond),
		retry.DelayType(retry.BackOffDelay),
		retry.LastErrorOnly(true),
		retry.Context(ctx),
		retry.OnRetry(func(n uint, err error) {
			a.log.Errorf("retry connect postgres err: %v", err)
		}),
	}

	return retry.Do(func() error {
		pgxConn, err := postgres.NewPgxConn(a.cfg.Postgresql)
		if err != nil {
			return errors.Wrap(err, "postgres.NewPgxConn")
		}
		a.pgxConn = pgxConn
		a.log.Infof("postgres connected (poolStat: %s)", utils.GetPostgresStats(a.pgxConn.Stat()))
		return nil
	}, retryOption...)
}
