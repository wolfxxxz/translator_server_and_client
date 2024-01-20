package config

import (
	"client/internal/apperrors"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const path = "configs/.env"

type Config struct {
	Token        string `env:"TOKEN"`
	AppPort      string `env:"APP_PORT"`
	Host         string `env:"HOST"`
	LogLevel     string `env:"LOGGER_LEVEL"`
	TimeZone     string `env:"TIME_ZONE"`
	TimeoutQuery string `env:"TIMEOUT_QUERY"`
}

func NewConfig(logger *logrus.Logger) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		appErr := apperrors.EnvConfigLoadError.AppendMessage(err)
		return nil, appErr
	}

	conf := &Config{}
	if err := env.Parse(conf); err != nil {
		appErr := apperrors.EnvConfigParseError.AppendMessage(err)
		return nil, appErr
	}

	logger.Info("Config has been parsed")
	return conf, nil
}
