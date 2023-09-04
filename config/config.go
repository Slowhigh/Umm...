package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/slowhigh/Umm/pkg/constants"
	"github.com/slowhigh/Umm/pkg/event_stores"
	"github.com/slowhigh/Umm/pkg/kafka_client"
	"github.com/slowhigh/Umm/pkg/logger"
	"github.com/slowhigh/Umm/pkg/migrations"
	"github.com/slowhigh/Umm/pkg/postgres"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Umm config path")
}

type Config struct {
	ServiceName          string                          `mapstructure:"serviceName"`
	Logger               logger.LogConfig                `mapstructure:"logger"`
	Timeouts             Timeouts                        `mapstructure:"timeouts" validate:"required"`
	Postgresql           postgres.Config                 `mapstructure:"postgres"`
	Kafka                *kafka_client.Config             `mapstructure:"kafka" validate:"required"`
	KafkaPublisherConfig event_stores.KafkaEventBusConfig `mapstructure:"kafkaPublisherConfig" validate:"required"`
	Http                 Http                            `mapstructure:"http"`
	MigrationsConfig     migrations.Config               `mapstructure:"migrations" validate:"required"`
}

type Timeouts struct {
	PostgresInitMilliseconds int  `mapstructure:"postgresInitMilliseconds" validate:"required"`
	PostgresInitRetryCount   uint `mapstructure:"postgresInitRetryCount" validate:"required"`
}

type MongoCollections struct {
	BankAccounts string `mapstructure:"bankAccounts" validate:"required"`
}

type Http struct {
	Port                string   `mapstructure:"port" validate:"required"`
	Development         bool     `mapstructure:"development"`
	BasePath            string   `mapstructure:"basePath" validate:"required"`
	BankAccountsPath    string   `mapstructure:"bankAccountsPath" validate:"required"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	postgresHost := os.Getenv(constants.PostgresqlHost)
	if postgresHost != "" {
		cfg.Postgresql.Host = postgresHost
	}

	postgresPort := os.Getenv(constants.PostgresqlPort)
	if postgresPort != "" {
		cfg.Postgresql.Port = postgresPort
	}

	dbUrl := os.Getenv(constants.MigrationsDbUrl)
	if dbUrl != "" {
		cfg.MigrationsConfig.DbURL = dbUrl
	}

	kafkaBrokers := os.Getenv(constants.KafkaBrokers)
	if kafkaBrokers != "" {
		cfg.Kafka.Brokers = []string{kafkaBrokers}
	}

	return cfg, nil
}
