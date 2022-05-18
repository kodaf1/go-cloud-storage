package main

import (
	"fmt"

	"github.com/kodaf1/go-cloud-storage/internal/config"
	"github.com/kodaf1/go-cloud-storage/pkg/logging"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("logger initialized")

	logger.Info("config initializing")
	cfg := config.GetConfig()

	fmt.Printf("%v\n", cfg)
}
