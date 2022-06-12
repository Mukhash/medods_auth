package main

import (
	"context"

	"github.com/Mukhash/medods_auth/config"
	"github.com/Mukhash/medods_auth/internal/controller/handler"
	"github.com/Mukhash/medods_auth/internal/dataprovider/store"
	"github.com/Mukhash/medods_auth/internal/service"
	"github.com/Mukhash/medods_auth/pkg/database/mongodb"
	"github.com/Mukhash/medods_auth/pkg/httpserver"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	cfg := config.DefaultConfig()
	logger := zap.NewExample()

	defer logger.Sync()

	mongoClient, err := mongodb.New(cfg, logger)
	if err != nil {
		logger.Fatal("connecting to mongoDB...")
	}

	repo := store.NewStore(ctx, mongoClient)
	authService := service.New(repo)
	handler := handler.New(authService)

	srv := httpserver.New(ctx, cfg, logger, handler)

	srv.Start()
}
