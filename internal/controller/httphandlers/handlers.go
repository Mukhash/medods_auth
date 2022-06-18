package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Mukhash/medods_auth/config"
	"github.com/Mukhash/medods_auth/internal/service"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type Handler struct {
	cfg         *config.Config
	logger      *zap.Logger
	AuthService service.AuthService
}

type AuthRequest struct {
	UUID string `json:"guid"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

func New(cfg *config.Config, logger *zap.Logger, authService service.AuthService) *Handler {
	return &Handler{
		cfg:         cfg,
		logger:      logger,
		AuthService: authService,
	}
}

func (h *Handler) Auth(c echo.Context) error {
	req := &AuthRequest{}

	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		h.logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	tokens, err := h.AuthService.CreateSession(req.UUID)
	if err != nil {
		h.logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	cookieRefresh := http.Cookie{Name: "refreshToken", Value: tokens.Refresh, HttpOnly: true}
	c.SetCookie(&cookieRefresh)

	cookieAccess := http.Cookie{Name: "accessToken", Value: tokens.Access}
	c.SetCookie((&cookieAccess))

	return c.JSON(http.StatusCreated, tokens)
}

func (h *Handler) Refresh(c echo.Context) error {
	req := &RefreshRequest{}

	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		h.logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	accessToken, err := h.AuthService.Refresh(req.RefreshToken)
	if err != nil {
		h.logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	res := RefreshResponse{AccessToken: accessToken}

	return c.JSON(http.StatusCreated, res)
}
