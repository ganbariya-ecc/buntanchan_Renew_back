package controller

import (
	"auth/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(ctx echo.Context) error {
	// セッション取得
	userid, err := utils.GetAuth(ctx, "userid")

	// エラー処理
	if err == nil {
		log.Println("userid : " + userid.(string))
	} else {
		log.Println(err)
	}

	return ctx.String(http.StatusOK, "hello world")
}
