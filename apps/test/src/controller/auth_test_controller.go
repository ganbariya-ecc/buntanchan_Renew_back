package controller

import (
	"log"
	"net/http"
	"test/sdk"

	"github.com/labstack/echo/v4"
)

func Auth_Test(ctx echo.Context) error {
	log.Println(ctx.Request().Header.Get("Authorized"))

	sdk.Auth(ctx.Request().Header.Get("Authorized"))

	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : "ok",
	})
}