package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetJWT(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"jwt":"jwtdata",
	})
}