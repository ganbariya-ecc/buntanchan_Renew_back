package router

import (
	"auth/controller"
	"auth/middlewares"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
)

func InitRouter() *echo.Echo {
	// Echo instance
	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// 認証用のRedis
	client := redis.NewClient(&redis.Options{
		Addr: "aredis:6379",
	})

	// New default RedisStore
	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		log.Fatal("failed to create redis store: ", err)
	}

	// Example changing configuration for sessions
	store.KeyPrefix("AuthSession_")
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   86400 * 365,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		HttpOnly: true,
	})

	// 静的ファイル提供
	router.Static("/", "./statics")

	// ミドルウェア設定
	router.Use(session.Middleware(store))

	// Routes
	// router.GET("/", controller.Hello)

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
	router.POST("/logout", controller.Logout, middlewares.AuthMiddleware)

	// Authed グループ
	authedg := router.Group("/authed")
	{
		// Auth ミドルウェア
		authedg.Use(middlewares.AuthMiddleware)

		authedg.POST("/jwt", controller.GetJWT)

		authedg.GET("/info",controller.GetInfo)
	}

	// 管理者のグループ
	adming := router.Group("/admin")
	{
		adming.POST("/signup",controller.AdminSignup)
		adming.POST("/login",controller.AdminLogin)
		adming.POST("/logout",controller.AdminLogout, middlewares.AdminAuthMiddleware)
	}

	// 管理者の操作 api
	admincg := router.Group("/adminc")
	{
		admincg.Use(middlewares.AdminAuthMiddleware)
		admincg.GET("/info",controller.GetAdminInfo)
		admincg.GET("/users",controller.GetUsers)
		admincg.GET("/userinfo",controller.AdminGetUserInfo)
	}

	// ラベル管理 API
	labelcg := admincg.Group("/labels")
	{
		// ラベルを更新する
		labelcg.POST("/update",controller.UpdateLabels)
	}

	// ユーザー管理 API
	usercg := admincg.Group("/user")
	{
		usercg.GET("/info",controller.UserInfo)
		usercg.DELETE("/delete",controller.DeleteUser)
		usercg.POST("/loginas",controller.LoginAS)
	}

	return router
}
