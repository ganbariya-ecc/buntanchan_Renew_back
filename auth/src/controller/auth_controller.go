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

// ユーザーIDとパスワードの構造体
type AuthData struct {
	UserID   string `json:"userid"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// ユーザーIDとパスワードでログイン
func BasicLogin(ctx echo.Context) error {
	// パラメーター取得
	data := new(AuthData)
	if err := ctx.Bind(data); err != nil {
		return err
	}

	// バリデーション
	if data.UserID == "" || data.Password == "" {
		return errors.New("userid or password is empty")
	}

	// パスワード認証
	userData, err := service.BasicLogin(data.UserID, data.Password)

	// エラー処理
	if err != nil {
		return err
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

	return ctx.JSON(http.StatusOK, userData.UserID)
}

// ユーザーIDとパスワードでサインアップ
func BasicSignup(ctx echo.Context) error {
	// パラメーター取得
	binddata := new(AuthData)
	if err := ctx.Bind(&binddata); err != nil {
		return err
	}

	// バリデーション
	if binddata.UserName == "" || binddata.Password == "" {
		return errors.New("userid or password is empty")
	}

	// ユーザー作成
	userData, err := service.BasicSignup(binddata.UserName, binddata.Password)

	// エラー処理
	if err != nil {
		return err
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

	log.Println(userData)

	return ctx.JSON(http.StatusOK, userData.UserID)
}

func Logout(ctx echo.Context) error {
	// セッション削除
	// セッション取得
	AuthSession, err := session.Get(utils.AuthSessionName, ctx)

	// エラー処理
	if err != nil {
		return err
	}

	// 認証済み解除
	err = utils.SetAuth(ctx, "authorized", false)
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
