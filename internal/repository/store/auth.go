package store

import (
	"context"

	"github.com/Mukhash/medods_auth/internal/models"
	"github.com/Mukhash/medods_auth/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	Ctx    context.Context
	Client *mongodb.Client
}

func NewStore(ctx context.Context, client *mongodb.Client) *Store {
	return &Store{
		Ctx:    ctx,
		Client: client,
	}
}

// FindUser if no user is found mongo.ErrNoDocements if thrown.
func (s *Store) FindUser(uuid string) (*models.User, error) {
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

func (s *Store) CreateUser(user *models.User) error {
	usersCollection := s.Client.DB.Collection("users")

	hashedToken, err := bcrypt.GenerateFromPassword([]byte(user.RefreshToken), 0)
	if err != nil {
		return err
	}

	filter := bson.D{{"uuid", user.UUID}}
	update := bson.D{{"$setOnInsert", bson.D{{"refreshToken", hashedToken}}}}

	upsert := true
	opts := &options.UpdateOptions{Upsert: &upsert}

	_, err = usersCollection.UpdateOne(s.Ctx, filter, update, opts)
	if err != nil {
		return err
	}

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
	return nil
}