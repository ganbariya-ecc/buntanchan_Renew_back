package controller

import (
	"auth/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAdminInfo(ctx echo.Context) error {
	userid := ctx.Get("adminid").(string)

	// ユーザー取得
	adminUser,err := service.AdminGetInfo(userid)

	// エラー処理
	if err != nil {
		log.Println("failed to get admininfo : " + err.Error())
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK,echo.Map{
		"adminid" : adminUser.UserID,
		"name" : adminUser.UserName,
	})
}

func GetUsers(ctx echo.Context) error {
	userid := ctx.Get("adminid").(string)

	// ユーザー一覧取得
	users,err := service.AdminGetUsers(userid)

	// エラー処理
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK,echo.Map{
		"users" : users,
	})
}