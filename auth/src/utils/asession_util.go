package utils

import (
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func SetData(ctx echo.Context, key string, value interface{}) error {
	// セッション取得
	AuthSession, err := session.Get("AuthSession", ctx)

	// エラー処理
	if err != nil {
		return err
	}

	// セッションの設定
	AuthSession.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   31536000,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	// データをセット
	AuthSession.Values[key] = value

	// セッション保存
	return AuthSession.Save(ctx.Request(), ctx.Response())
}

func GetValue(ctx echo.Context, key string) (interface{}, error) {
	// セッション取得
	AuthSession, err := session.Get("AuthSession", ctx)

	// エラー処理
	if err != nil {
		return "", err
	}

	// データを取得
	val, exits := AuthSession.Values[key]

	// データを取得
	if exits {
		return val, nil
	}

	// トークン取得
	return "", errors.New("data not found")
}