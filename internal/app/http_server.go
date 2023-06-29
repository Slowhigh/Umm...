package app

import (
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	maxHeaderBytes = 1 << 20
	stackSize      = 1 << 10
	bodyLimit      = "2M"
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
	gzipLevel      = 5
)

func (a *app) runHttpServer() error {
	a.mapRoutes()

	a.echo.Server.ReadTimeout = readTimeout
	a.echo.Server.WriteTimeout = writeTimeout
	a.echo.Server.MaxHeaderBytes = maxHeaderBytes

	return a.echo.Start(a.cfg.Http.Port)
}

func (a *app) mapRoutes() {
	a.echo.Use(a.middlewareManager.RequestLoggerMiddleware)
	a.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: stackSize,
	}))
	a.echo.Use(middleware.RequestID())
	a.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	a.echo.Use(middleware.BodyLimit(bodyLimit))
}
