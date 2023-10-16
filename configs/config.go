package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server Http `mapstructure:"http" validate:"required"`
}

type Http struct {
	Port    int    `mapstructure:"port" validate:"required"`
	Timeout string `mapstructure:"timeout" validate:"required"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")

	cfg := &Config{}

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("viper.ReadInConfig: %w", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("viper.Unmarshal: %w", err)
	}

	return cfg, nil
}
