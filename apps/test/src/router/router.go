package router

import (
	"test/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() (*echo.Echo) {
	// Echo instance
	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// Routes
	router.GET("/", controller.Hello)

	router.POST("/atest",controller.Auth_Test)

	return router
}