package repository

import "github.com/Mukhash/medods_auth/internal/models"

type Repository interface {
	FindUser(uuid string) (*models.User, error)
	CreateUser(user *models.User) error
	InsertRefresh(user *models.User) error
}
