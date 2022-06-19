package store

import (
	"context"
	"fmt"

	"github.com/Mukhash/medods_auth/internal/models"
	"github.com/Mukhash/medods_auth/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	Ctx    context.Context
	Logger *zap.Logger
	Client *mongodb.Client
}

func NewStore(ctx context.Context, logger *zap.Logger, client *mongodb.Client) *Store {
	return &Store{
		Ctx:    ctx,
		Logger: logger,
		Client: client,
	}
}

// FindSession if no user is found mongo.ErrNoDocements if thrown.
func (s *Store) FindSession(uuid string) (*models.User, error) {
	user := &models.User{}
	usersCollection := s.Client.DB.Collection("users")

	res := usersCollection.FindOne(s.Ctx, bson.D{{"uuid", uuid}})

	if res.Err() != nil {
		return nil, res.Err()
	}

	if err := res.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) InsertSession(user *models.User) error {
	usersCollection := s.Client.DB.Collection("users")

	hashedToken, err := bcrypt.GenerateFromPassword([]byte(user.RefreshToken), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	filter := bson.D{{"uuid", user.UUID}}
	update := bson.D{{"$set", bson.D{{"refreshToken", hashedToken}}}}

	upsert := true
	opts := &options.UpdateOptions{Upsert: &upsert}

	_, err = usersCollection.UpdateOne(s.Ctx, filter, update, opts)
	if err != nil {
		return err
	}
	s.Logger.Info(fmt.Sprintf("MongoDB UpdateOne: guid %s", user.UUID))

	return nil
}

func (s *Store) InsertRefresh(user *models.User) error {
	usersCollection := s.Client.DB.Collection("users")

	hashedToken, err := bcrypt.GenerateFromPassword([]byte(user.RefreshToken), 0)
	if err != nil {
		return err
	}

	filter := bson.D{{"uuid", user.UUID}}
	update := bson.D{{"$set", bson.D{{"refreshToken", hashedToken}}}}

	_, err = usersCollection.UpdateOne(s.Ctx, filter, update)
	if err != nil {
		return err
	}
	s.Logger.Info(fmt.Sprintf("MongoDB UpdateOne: guid %s", user.UUID))

	return nil
}
