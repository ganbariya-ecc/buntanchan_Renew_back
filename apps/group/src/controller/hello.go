package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK,"hello world")
}