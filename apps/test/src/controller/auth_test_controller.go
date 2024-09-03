package controller

import (
	"log"
	"net/http"
	"test/sdk"

	"github.com/labstack/echo/v4"
)

func Auth_Test(ctx echo.Context) error {
	log.Println(ctx.Request().Header.Get("Authorized"))

	user,err := sdk.Auth(ctx.Request().Header.Get("Authorized"))

	// エラー処理
	if err != nil {
		log.Println(user)
		return ctx.JSON(http.StatusUnauthorized,echo.Map{
			"result" : "failed",
  		})
	}

	log.Println(user)

	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : "ok",
	})
}