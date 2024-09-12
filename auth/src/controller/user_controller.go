package controller

import (
	"auth/model"
	"log"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
)

func GetIcon(ctx echo.Context) error {
	userid := ctx.Param("userid")

	// ユーザーを取得する
	user,err := model.GetUserByID(userid)

	// エラー処理
	if err != nil {
		log.Println("Failed to get icon : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError,echo.Map{
			"result" : "Failed to get icon",
		})
	}

	return ctx.File(path.Join(model.UserIconDir,user.UserID + ".jpeg"))
}