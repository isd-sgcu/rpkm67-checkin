package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port        string
	Env         string
	ServiceName string
}

type DbConfig struct {
	Url string
}

type TracerConfig struct {
	Endpoint string
}

type Config struct {
	App    AppConfig
	Db     DbConfig
	Tracer TracerConfig
}

func LoadConfig() (*Config, error) {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	appConfig := AppConfig{
		Port:        os.Getenv("APP_PORT"),
		Env:         os.Getenv("APP_ENV"),
		ServiceName: os.Getenv("APP_SERVICE_NAME"),
	}

	dbConfig := DbConfig{
		Url: os.Getenv("DB_URL"),
	}

	tracerConfig := TracerConfig{
		Endpoint: os.Getenv("TRACER_ENDPOINT"),
	}

	return &Config{
		App:    appConfig,
		Db:     dbConfig,
		Tracer: tracerConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}
