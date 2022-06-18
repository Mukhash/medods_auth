package main

import (
	"context"
	"log"

	"github.com/Mukhash/medods_auth/config"
	"github.com/Mukhash/medods_auth/internal/controller/handler"
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

	ctx := context.Background()

	logger := zap.NewExample()

	defer logger.Sync()

	mongoClient, err := mongodb.New(cfg, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	repo := store.NewStore(ctx, mongoClient)
	authService := service.New(repo, cfg, logger)
	handler := handler.New(authService)

	srv := httpserver.New(ctx, cfg, logger, handler)

	srv.Start()
}
