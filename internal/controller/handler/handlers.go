package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Mukhash/medods_auth/internal/service"
	"github.com/labstack/echo"
)

type Handler struct {
	AuthService service.AuthService
}

type AuthRequest struct {
	UUID string `json:"guid"`
}

func New(authService service.AuthService) *Handler {
	return &Handler{
		AuthService: authService,
	}
}

func (h *Handler) Auth(c echo.Context) error {
	req := &AuthRequest{}

	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		return errors.New("could not decode data")
	}

	tokens, err := h.AuthService.Auth(req.UUID)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		return errors.New("could not decode data")
	}

	return c.JSON(http.StatusCreated, tokens)
}

func (h *Handler) Refresh(c echo.Context) error {
	return nil
}
