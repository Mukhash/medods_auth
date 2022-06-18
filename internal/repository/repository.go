package repository

import "github.com/Mukhash/medods_auth/internal/models"

type Repository interface {
	FindSession(uuid string) (*models.User, error)
	InsertSession(user *models.User) error
	InsertRefresh(user *models.User) error
}
