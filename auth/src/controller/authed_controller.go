package controller

import (
	"auth/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetJWT(ctx echo.Context) error {
	// ユーザーID取得
	userid := ctx.Get("userid")

	// エラー処理
	if userid == nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error" : "failed to get userid",
		})
	}

	// JWT生成
	token,herr := service.GenJWT(userid.(string))

	// エラー処理
	if herr != nil {
		log.Println("JWT Error : " + herr.Error())
		return ctx.JSON(herr.Status,echo.Map{
			"error" : herr.Message,
		})
	}

	return ctx.JSON(http.StatusOK,echo.Map{
		"jwt": token,
	})
}

func GetInfo(ctx echo.Context) error {
	// ユーザーID取得
	userid := ctx.Get("userid")

	// エラー処理
	if userid == nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error" : "failed to get userid",
		})
	}

	// ユーザーID取得
	user,err := service.GetUserInfo(userid.(string))

	// エラー処理
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,echo.Map{
			"result" : "Failed to obtain user information",
		})
	}

	// ユーザー情報返却
	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : user,
	})
}