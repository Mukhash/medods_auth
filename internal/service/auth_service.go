package service

import (
	"github.com/Mukhash/medods_auth/config"
	"github.com/Mukhash/medods_auth/internal/dataprovider"
	"github.com/Mukhash/medods_auth/internal/models"
	"go.uber.org/zap"
)

type authService struct {
	logger *zap.Logger
	cfg    *config.Config
	repo   dataprovider.Repository
}

func New(repo dataprovider.Repository, cfg *config.Config, logger *zap.Logger) *authService {
	return &authService{repo: repo, cfg: cfg, logger: logger}
}

func (a *authService) Auth(uuid string) (*models.Token, error) {
	user, err := a.repo.FindUser(uuid)
	if err != nil {
		return nil, err
	}

}

func (a *authService) Refresh() {

}
