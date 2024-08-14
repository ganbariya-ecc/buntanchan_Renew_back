package controller

import (
	"auth/service"
	"auth/utils"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/markbates/goth/gothic"
)

func StartOauth(ctx echo.Context) error {
	// パラメーターからプロバイダー取得
	provider := ctx.Param("provider")

	//リクエストにプロバイダー設定
	ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(),"provider",provider)))
	
	// Oauth セッション開始
	gothic.BeginAuthHandler(ctx.Response().Writer,ctx.Request())

	return nil
}

func CallbackOauth(ctx echo.Context) error {
	// パラメーターからプロバイダー取得
	provider := ctx.Param("provider")

	//リクエストにプロバイダー設定
	ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(),"provider",provider)))

	// Oauth ユーザー取得
	gothUser,err := gothic.CompleteUserAuth(ctx.Response(),ctx.Request())

	// エラー処理
	if err != nil {
		log.Println("Oauth2 Callback Error : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError,"failed to complete Oauth")
	}

	// Oauth完了
	_,err = service.CallbackOauth(gothUser)

	// エラー処理
	if err != nil {
		log.Println("failed to oauth callback : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError,nil)
	}

	// セッション保存
	err = utils.SetData(ctx,"userid",gothUser.UserID)
	if err != nil {
		log.Println("failed to set token : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError,nil)
	}

	return ctx.Redirect(http.StatusFound,"/auth/")
}