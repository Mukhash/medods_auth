package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Mukhash/medods_auth/config"
	handler "github.com/Mukhash/medods_auth/internal/controller/httphandlers"
	"github.com/Mukhash/medods_auth/internal/repository/store"
	"github.com/Mukhash/medods_auth/internal/service"
	"github.com/Mukhash/medods_auth/pkg/database/mongodb"
	"github.com/Mukhash/medods_auth/pkg/httpserver"
	"go.uber.org/zap"
)

var (
	configFilePath = "./config/config"
)

func main() {
	var err error

	cfgFile, err := config.LoadConfig(configFilePath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	fmt.Println(cfg.DB.URL)
	ctx, cancel := context.WithCancel(context.Background())

	logger := zap.NewExample()

	defer logger.Sync()

	mongoClient, err := mongodb.New(cfg, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	repo := store.NewStore(ctx, logger, mongoClient)
	authService := service.New(repo, cfg, logger)
	handler := handler.New(cfg, logger, authService)

	srv := httpserver.New(ctx, cfg, logger, handler)

	go func(cancel func()) {
		srv.Start(cancel)
	}(cancel)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		logger.Error("signal.Notify")
	case <-ctx.Done():
		logger.Error("ctx.Done")
	}

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server shutdown...")
	}

	logger.Info("Server Exited Properly")
}
