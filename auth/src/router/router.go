package router

import (
	"auth/controller"
	"auth/middlewares"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

func InitRouter() *echo.Echo {
	// Echo instance
	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// ミドルウェア設定
	router.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))))

	// Routes
	router.GET("/", controller.Hello)

	// Oauth グループ
	oauthg := router.Group("/oauth")
	{
		// 認証エンドポイント
		oauthg.GET("/:provider", controller.StartOauth)

		// Oauth Callback エンドポイント
		oauthg.GET("/callback/:provider", controller.CallbackOauth)
	}

	// Auth グループ
	// ログイン
	router.POST("/login", controller.BasicLogin)

	// サインアップ
	router.POST("/signup", controller.BasicSignup)

	// 認証エンドポイント
	router.POST("/logout", controller.Logout,middlewares.AuthMiddleware)

	// Authed グループ
	authedg := router.Group("/authed")
	{
		// Auth ミドルウェア
		authedg.Use(middlewares.AuthMiddleware)

		authedg.POST("/jwt", controller.GetJWT)
	}

	return router
}
