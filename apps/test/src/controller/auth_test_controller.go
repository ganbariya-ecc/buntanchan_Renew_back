package controller

import (
	"log"
	"net/http"
	"template/sdks/authsdk"

	"github.com/labstack/echo/v4"
)

func Auth_Test(ctx echo.Context) error {
	log.Println(ctx.Request().Header.Get("Authorized"))

	// 認証を実行する
	user, err := authsdk.Auth(ctx.Request().Header.Get("Authorized"))

	// エラー処理
	if err != nil {
		log.Println(user)
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"result": "failed",
		})
	}

	log.Println(user)

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "ok",
	})
}

