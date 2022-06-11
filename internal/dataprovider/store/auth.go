package store

import "context"

type Store struct {
	Ctx context.Context
}

func NewStore(ctx context.Context) *Store {
	return &Store{
		Ctx: ctx,
	}
}
