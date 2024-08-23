package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAdminInfo(ctx echo.Context) error {
	userid := ctx.Get("adminid").(string)

	return ctx.JSON(http.StatusOK,echo.Map{
		"adminid" : userid,
	})
}
