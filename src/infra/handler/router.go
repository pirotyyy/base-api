package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	// Create Echo Instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Health Check
	e.GET("/health", func(c echo.Context) error { return c.JSON(http.StatusOK, "basic-api") })

	return e
}
