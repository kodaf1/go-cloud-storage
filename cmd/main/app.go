package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/kodaf1/go-cloud-storage/internal/composites"
	"github.com/kodaf1/go-cloud-storage/internal/config"
	"github.com/kodaf1/go-cloud-storage/pkg/logging"
	"github.com/kodaf1/go-cloud-storage/pkg/shutdown"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("logger initialized")

	logger.Info("config initializing")
	cfg := config.GetConfig()

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("mongodb composite initializing")
	mongoDBC, err := composites.NewMongoDBComposite(
		context.Background(),
		cfg.MongoDB.Host,
		cfg.MongoDB.Port,
		cfg.MongoDB.Username,
		cfg.MongoDB.Passowrd,
		cfg.MongoDB.Database,
		cfg.MongoDB.AuthDB,
	)
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}

	logger.Info("s3 composite initializing")
	s3C := composites.NewS3Composite(
		cfg.S3.AccessKey,
		cfg.S3.SecretKey,
		cfg.S3.SessionToken,
		cfg.S3.Endpoint,
		cfg.S3.Region,
		cfg.S3.Bucket,
	)

	logger.Info("user composite initializing")
	userComposite, err := composites.NewFileComposite(mongoDBC, s3C)
	if err != nil {
		logger.Fatal("user composite failed")
	}
	userComposite.Handler.Register(router)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		logger.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 120 * time.Second,
		ReadTimeout:  120 * time.Second,
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM},
		server)

	logger.Println("application initialized and started")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err)
		}
	}
}
