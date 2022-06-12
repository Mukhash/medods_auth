package store

import (
	"context"

	"github.com/Mukhash/medods_auth/pkg/database/mongodb"
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
