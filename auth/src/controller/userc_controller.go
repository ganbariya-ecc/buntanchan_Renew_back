package controller

import (
	"auth/service"
	"auth/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteUser(ctx echo.Context) error {
	// adminid 取得
	adminid := ctx.Get("adminid").(string)

	// ユーザーID
	userid := ctx.Request().Header.Get("userid")

	return service.Admin_DeleteUser(adminid,userid)
}

func UserInfo(ctx echo.Context) error {
	// adminid 取得
	adminid := ctx.Get("adminid").(string)

	// ユーザーID
	userid := ctx.Request().Header.Get("userid")

	// ユーザー取得
	user_data,err := service.Admin_GetUser_Info(adminid,userid)

	// エラー処理
	if err != nil {
		log.Println("Failed to obtain user information : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError,echo.Map{
			"result" : "Failed to obtain user information",
		})
	}

	return ctx.JSON(http.StatusOK,echo.Map{
		"user" : user_data,
	})
}

func LoginAS(ctx echo.Context) error {
	// adminid 取得
	// adminid := ctx.Get("adminid").(string)

	// ユーザーID
	userid := ctx.Request().Header.Get("userid")

	// ユーザーとして　ログイン
	userData,err := service.GetUserInfo(userid)

	// エラー処理
	if err != nil {
		log.Println("failed to get user : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError,echo.Map{
			"result" : "failed to get user",
		})
	}

	// セッション保存
	err = utils.SetAuth(ctx, "userid", userData.UserID)
	if err != nil {
		return err
	}

	// 認証済みにする
	err = utils.SetAuth(ctx, "authorized", true)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : "Success",
	})
}