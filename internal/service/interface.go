package service

import "github.com/Mukhash/medods_auth/internal/models"

type AuthService interface {
	Auth(payload string) (*models.Token, error)
	Refresh(refreshToken string) (string, error)
}
