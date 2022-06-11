package handler

import (
	"github.com/Mukhash/medods_auth/internal/service"
	"github.com/labstack/echo"
)

type Handler struct {
	AuthService service.AuthService
}

func New(authService service.AuthService) *Handler {
	return &Handler{
		AuthService: authService,
	}
}

func (h *Handler) Auth(c echo.Context) error {
	return nil
}

func (h *Handler) Refresh(c echo.Context) error {
	return nil
}
