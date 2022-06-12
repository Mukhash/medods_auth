package service

import "github.com/Mukhash/medods_auth/internal/dataprovider"

type authService struct {
	Repo dataprovider.Repository
}

func New(repo dataprovider.Repository) *authService {
	return &authService{Repo: repo}
}

func (a *authService) Auth() {

}

func (a *authService) Refresh() {

}
