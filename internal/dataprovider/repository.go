package dataprovider

import "github.com/Mukhash/medods_auth/internal/models"

type Repository interface {
	FindUser(uuid string) (models.User, error)
}
