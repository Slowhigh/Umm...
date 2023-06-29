package app

import (
	"time"

	"github.com/slowhigh/Umm/pkg/migrations"
)

const (
	waitShutDownDuration = 3 * time.Second
)

func (a *app) runMigrate() error {
	a.log.Infof("Run migrations with config: %+v", a.cfg.MigrationsConfig)

	version, dirty, err := migrations.RunMigrations(a.cfg.MigrationsConfig)
	if err != nil {
		a.log.Errorf("RunMigrations err: %v", err)
		return err
	}

	a.log.Infof("Migrations successfully created: version: %d, dirty: %v", version, dirty)
	return nil
}

func (a *app) waitShutDown(duration time.Duration) {
	go func() {
		time.Sleep(duration)
		a.doneCh <- struct{}{}
	}()
}
