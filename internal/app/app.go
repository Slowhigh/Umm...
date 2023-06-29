package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/slowhigh/Umm/config"
	"github.com/slowhigh/Umm/pkg/logger"
	"github.com/slowhigh/Umm/pkg/middlewares"
)

type app struct {
	log               logger.Logger
	cfg               config.Config
	middlewareManager middlewares.MiddlewareManager
	pgxConn           *pgxpool.Pool
	doneCh            chan struct{}
	echo              *echo.Echo
}

func NewApp(log logger.Logger, cfg config.Config) *app {
	return &app{log: log, cfg: cfg, echo: echo.New()}
}

func (a *app) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	a.middlewareManager = middlewares.NewMiddlewareManager(a.log, a.cfg)

	if err := a.connectPostgres(ctx); err != nil {
		return err
	}
	defer a.pgxConn.Close()

	if err := a.runMigrate(); err != nil {
		return err
	}

	go func() {
		if err := a.runHttpServer(); err != nil {
			a.log.Errorf("runHttpServer (err: %v)", err)
			cancel()
		}
	}()
	a.log.Infof("%s is listening on PORT: %v", a.cfg.ServiceName, a.cfg.Http.Port)

	<-ctx.Done()
	a.waitShutDown(waitShutDownDuration)
	
	if err := a.echo.Shutdown(ctx); err != nil {
		a.log.Warnf("Shutdown (err: %v)", err)
	}

	<-a.doneCh
	a.log.Infof("%s app exited properly", a.cfg.ServiceName)
	return nil
}
