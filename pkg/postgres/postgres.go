package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Config struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
	SSLMode  bool   `mapstructure:"sslMode"`
}

const (
	minConns          = 10
	maxConns          = 50
	healthCheckPeriod = 1 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	lazyConnect       = false
)

func NewPgxConn(cfg Config) (*pgxpool.Pool, error) {
	ctx := context.Background()
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password)

	pgxCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	pgxCfg.MinConns = minConns
	pgxCfg.MaxConns = maxConns
	pgxCfg.HealthCheckPeriod = healthCheckPeriod
	pgxCfg.MaxConnIdleTime = maxConnIdleTime
	pgxCfg.MaxConnLifetime = maxConnLifetime
	pgxCfg.LazyConnect = lazyConnect

	pool, err := pgxpool.ConnectConfig(ctx, pgxCfg)
	if err != nil {
		return nil, errors.Wrap(err, "pgx.ConnectConfig")
	}

	return pool, nil
}
