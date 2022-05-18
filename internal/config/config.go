package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/kodaf1/go-cloud-storage/pkg/logging"
)

type Config struct {
	IsDebug *bool `env:"IS_DEBUG,required" envDefault:"true"`
	Listen  struct {
		Type   string `env:"TYPE" envDefault:"port"`
		BindIP string `env:"BIND_IP" envDefault:"0.0.0.0"`
		Port   string `env:"PORT" envDefault:"8080"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application config")

		err := godotenv.Load(".env")
		if err != nil {
			logger.Fatal(err)
		}

		instance = &Config{}
		if err := env.Parse(instance); err != nil {
			logger.Fatal(err)
		}

	})
	return instance
}
