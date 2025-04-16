package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"quick-poll/internal/service"
)

type Handler struct {
	s *service.Service
}

func NewHandler(c *service.Service) *echo.Echo {
	app := echo.New()
	app.Use(middleware.Logger())

	h := &Handler{
		c,
	}

	app.POST("/polls", h.CreatePoll)
	app.POST("/polls/:id/vote", h.Vote)
	app.GET("/polls/:id/results", h.GetResults)

	return app

}
