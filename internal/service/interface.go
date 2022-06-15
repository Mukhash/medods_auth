package service

import "github.com/Mukhash/medods_auth/internal/models"

type AuthService interface {
	Auth(uuid string) models.Token
	Refresh(refreshToken string) error
}
