package router

import (
	"group/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"group/middlewares"
)

func InitRouter() *echo.Echo {
	// Echo instance
	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// Routes
	router.GET("/", controller.Hello)

	router.POST("/atest", controller.Auth_Test)

	// グループ作成エンドポイント
	router.POST("/create", controller.CreateGroup, middlewares.AuthMiddleware)

	currentg := router.Group("/current")
	{
		currentg.Use(middlewares.AuthMiddleware)

		// グループ取得エンドポイント
		currentg.GET("", controller.GetCurrentGroup)

		// メンバー取得エンドポイント
		currentg.GET("/members", controller.GetCurrentMembers)
	}

	return router
}
