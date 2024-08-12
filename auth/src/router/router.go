package router

import (
	"auth/controller"

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

	oauthg := router.Group("/oauth")
	// Oauth エンドポイント
	oauthg.GET("/:provider",controller.StartOauth)
	oauthg.GET("/callback/:provider",controller.CallbackOauth)

	return router
}