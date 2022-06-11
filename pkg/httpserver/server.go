package api

import (
	"context"

	"github.com/Mukhash/medods_auth/config"
	"github.com/Mukhash/medods_auth/internal/controller/handler"
	"github.com/Mukhash/medods_auth/internal/service"
	"github.com/labstack/echo"
)

type Server struct {
	*echo.Echo
	config *config.Config
	ctx    context.Context
}

func NewServer(ctx context.Context, config *config.Config) *Server {
	e := echo.New()

	srv := &Server{
		e,
		config,
		ctx,
	}
	as := service.New()
	handler := handler.New(as)
	//e.Use(CORSMiddlewareWrapper)
	e.Static("/static", "./assets")
	v2 := e.Group(config.API.MainPath)
	//v2.GET("/sayhi", handler.SayHi, mdw.Auth)
	passRoutes := v2.Group("/auth")
	{
		passRoutes.POST("/auth", handler.Auth)
		passRoutes.POST("/refresh", handler.Refresh)

	}

	//logger.Info("try to run api")
	return srv
}

func (srv Server) Start() {

}
