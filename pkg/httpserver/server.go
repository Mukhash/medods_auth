package httpserver

import (
	"context"
	"net/http"

	"github.com/Mukhash/medods_auth/config"
	"github.com/Mukhash/medods_auth/internal/controller/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
)

type Server struct {
	*echo.Echo
	config *config.Config
	ctx    context.Context
}

// CORSMiddlewareWrapper https://github.com/labstack/echo/issues/1146
func CORSMiddlewareWrapper(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := ctx.Request()
		dynamicCORSConfig := middleware.CORSConfig{
			AllowOrigins: []string{req.Header.Get("Origin")},
			AllowHeaders: []string{"Accept", "Cache-Control", "Content-Type", "X-Requested-With", "Content-Type", "api_key", "Authorization"},
			AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		}
		CORSMiddleware := middleware.CORSWithConfig(dynamicCORSConfig)
		CORSHandler := CORSMiddleware(next)
		return CORSHandler(ctx)
	}
}

func New(ctx context.Context, config *config.Config, logger *zap.Logger, handler *handler.Handler) *Server {
	e := echo.New()

	srv := &Server{
		e,
		config,
		ctx,
	}

	e.Use(CORSMiddlewareWrapper)
	e.Static("/static", "./assets")
	v2 := e.Group(config.API.MainPath)
	passRoutes := v2.Group("/auth")
	{
		passRoutes.POST("/auth", handler.Auth)
		passRoutes.POST("/refresh", handler.Refresh)

	}

	logger.Info("try to run api")
	return srv
}

func (srv Server) Start() {

}
