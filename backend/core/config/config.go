package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server struct {
			Port string `env:"SERVER_PORT,required"`
			Host string `env:"SERVER_HOST,required"`
		}
		Database struct {
			ConnString string `env:"POSTGRE_CONN_STRING,required"`
		}
		NSQConfig struct {
			ReadTimeout  int    `env:"NSQ_READ_TIMEOUT"`
			WriteTimeout int    `env:"NSQ_WRITE_TIMEOUT"`
			MaxAttempts  int    `env:"NSQ_MAX_ATTEMPTS"`
			MaxInFlight  int    `env:"NSQ_MAX_IN_FLIGHT"`
			NSQd         string `env:"NSQD_STRING"`
		}
		RedisConfig struct {
			Addr     string `env:"REDIS_ADDRESS"`
			Password string `env:"REDIS_PASSWORD"`
		}
	}
)

func InitConfig() (*Config, error) {
	var cfg = Config{}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
