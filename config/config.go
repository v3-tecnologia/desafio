package config

import (
	"log/slog"

	env "github.com/caarlos0/env/v6"
)

type environment struct {
	DatabaseUser       string `env:"DATABASE_USER,required"`
	DatabasePassword   string `env:"DATABASE_PASSWORD,required"`
	DatabaseHost       string `env:"DATABASE_HOST,required"`
	DatabaseName       string `env:"DATABASE_NAME,required"`
	DatabaseSSL        string `env:"DATABASE_SSL,required"`
	AWSSecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY,required"`
	AWSAccessKeyID     string `env:"AWS_ACCESS_KEY,required"`
}

func NewConfig() (*Config, error) {
	slog.Info("Loading environment...")
	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return nil, err
	}

	slog.Info("Environment loaded successfully!")

	cfg := Config{
		DatabaseConfig: &databaseConfig{
			User:         environment.DatabaseUser,
			Password:     environment.DatabasePassword,
			Host:         environment.DatabaseHost,
			DatabaseName: environment.DatabaseName,
			SSL:          environment.DatabaseSSL,
		},
		AWSConfig: &awsConfig{
			AccessKeyID:     environment.AWSAccessKeyID,
			SecretAccessKey: environment.AWSSecretAccessKey,
		},
	}

	return &cfg, nil
}

type Config struct {
	AWSConfig      *awsConfig
	DatabaseConfig *databaseConfig
}

type databaseConfig struct {
	User         string
	Password     string
	Host         string
	DatabaseName string
	SSL          string
}

type awsConfig struct {
	AccessKeyID     string
	SecretAccessKey string
}
