package router

import (
	"task/controller"
	"task/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	router.POST("/gtest",controller.GroupTest,middlewares.AuthMiddleware)

	router.POST("/create",controller.CreateTask,middlewares.AuthMiddleware)
	router.POST("/upimg",controller.UploadTaskImg,middlewares.AuthMiddleware)
	router.POST("/taskimg",controller.UploadTaskImg,middlewares.AuthMiddleware)
	router.GET("/tasks",controller.GetTasks,middlewares.AuthMiddleware)
	router.GET("/info",controller.GetTask,middlewares.AuthMiddleware)

	return router
}
