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
        Addr:     "aredis:6379",
    })

    // New default RedisStore
    store, err := redisstore.NewRedisStore(context.Background(), client)
    if err != nil {
        log.Fatal("failed to create redis store: ", err)
    }

    // Example changing configuration for sessions
    store.KeyPrefix("AuthSession_")
    store.Options(sessions.Options{
        Path:   "/",
        Domain: "",
        MaxAge:  86400 * 365,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		HttpOnly: true,
    })

	// ミドルウェア設定
	router.Use(session.Middleware(store))

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
