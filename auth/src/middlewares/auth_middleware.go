package middlewares

import (
	"auth/model"
	"auth/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// エラーの時はエラーを返す
		if err := next(ctx); err != nil {
			ctx.Error(err)
		}

		// 初期化
		ctx.Set("userid", "")

		// UserID取得
		userid, err := utils.GetValue(ctx, "userid")

		// エラー処理
		if err != nil {
			log.Println("failed to get userid : " + err.Error())
			return ctx.NoContent(http.StatusUnauthorized)
		}

		// ユーザー取得
		userData, err := model.GetUserByID(userid)

		// エラー処理
		if err != nil {
			log.Println("failed to get user : " + err.Error())
			return ctx.NoContent(http.StatusUnauthorized)
		}

		log.Println("userid : " + userid)
		log.Println(userData)

		// ユーザーIDをセット
		ctx.Set("userid", userid)

		return nil
	}
}
