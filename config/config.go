package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

// Config default config structure
type Config struct {
	Exchange ExchangeConfig `env:",prefix=EXCHANGE_,required"`
	Database DatabaseConfig `env:",prefix=DB_,required"`
	Logger   LoggerConfig   `env:",prefix=LOGGER_"`
}

// LoggerConfig for logger configuration
type LoggerConfig struct {
	DisableCaller     bool   `env:"CALLER,default=false"`
	DisableStacktrace bool   `env:"STACKTRACE,default=false"`
	Level             string `env:"LEVEL,default=debug"`
}

// ExchangeConfig for exchange configuration
type ExchangeConfig struct {
	Url      string   `env:"URL,required"`
	Origin   string   `env:"URL,required"`
	Protocol string   `env:"URL,default="`
	Symbols  []string `env:"SYMBOLS,required"`
	Channels []string `env:"CHANNELS,required"`
}

// DatabaseConfig for db config
type DatabaseConfig struct {
	Host     string `env:"HOST,required"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD"`
	Base     string `env:"BASE"`
}

// NewConfig init default config for application
func NewConfig(ctx context.Context) (*Config, error) {
	var cfg Config

	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil

}
