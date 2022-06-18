package mongodb

import (
	"context"

	"github.com/Mukhash/medods_auth/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Client struct {
	*mongo.Client
	DB     *mongo.Database
	logger *zap.Logger
}

func New(cfg *config.Config, logger *zap.Logger) (*Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.DB.URL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	logger.Info("Connected to MongoDB...")

	authDatabase := client.Database("auth")

	return &Client{client, authDatabase, logger}, nil
}
