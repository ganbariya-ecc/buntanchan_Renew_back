package controller

import (
	"auth/service"
	"auth/utils"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// ユーザー名とパスワードの構造体
type AdminData struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// ユーザー名とパスワードでログイン
func AdminLogin(ctx echo.Context) error {
	// パラメーター取得
	data := new(AdminData)
	if err := ctx.Bind(data); err != nil {
		return err
	}

	// バリデーション
	if data.UserName == "" || data.Password == "" {
		return errors.New("userName or password is empty")
	}

	// パスワード認証
	userData, err := service.AdminLogin(data.UserName, data.Password)

	// エラー処理
	if err != nil {
		return err
	}

	// セッション保存
	err = utils.SetAdmin(ctx, "userid", userData.UserID)
	if err != nil {
		return err
	}

	// 認証済みにする
	err = utils.SetAdmin(ctx, "authorized", true)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, userData.UserID)
}

// ユーザー名とパスワードでサインアップ
func AdminSignup(ctx echo.Context) error {
	// パラメーター取得
	binddata := new(AdminData)
	if err := ctx.Bind(&binddata); err != nil {
		return err
	}

	// バリデーション
	if binddata.UserName == "" || binddata.Password == "" {
		return errors.New("username or password is empty")
	}

	// ユーザー作成
	userData, err := service.AdminSignup(binddata.UserName, binddata.Password)

	// エラー処理
	if err != nil {
		return err
	}

	// セッション保存
	err = utils.SetAdmin(ctx, "userid", userData.UserID)
	if err != nil {
		return err
	}

	// 認証済みにする
	err = utils.SetAdmin(ctx, "authorized", true)
	if err != nil {
		return err
	}

	log.Println(userData)

	return ctx.JSON(http.StatusOK, userData.UserID)
}

func AdminLogout(ctx echo.Context) error {
	// セッション削除
	AuthSession, err := session.Get(utils.AdminSessionName, ctx)

	// エラー処理
	if err != nil {
		return err
	}

	// 認証済み解除
	err = utils.SetAdmin(ctx, "authorized", false)
	if err != nil {
		return err
	}

	// セッション削除
	AuthSession.Options.MaxAge = -1

	// セッション更新
	err = AuthSession.Save(ctx.Request(), ctx.Response())

	// エラー処理
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"result": "success",
	})
}

