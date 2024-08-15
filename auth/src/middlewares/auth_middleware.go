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
		// // エラーの時はエラーを返す
		// if err := next(ctx); err != nil {
		// 	return err
		// }

		// 初期化
		ctx.Set("userid", "")

		// 認証済みか
		authorized, err := utils.GetValue(ctx, "authorized")

		// エラー処理
		if err != nil {
			log.Println("failed to get authorized : " + err.Error())
			return ctx.NoContent(http.StatusUnauthorized)
		}

		// 認証済みじゃない場合
		if !authorized.(bool) {
			log.Println("unauthorized")
			return ctx.NoContent(http.StatusUnauthorized)
		}

		// UserID取得
		userid, err := utils.GetValue(ctx, "userid")

		// エラー処理
		if err != nil {
			log.Println("failed to get userid : " + err.Error())
			return ctx.NoContent(http.StatusUnauthorized)
		}

		// ユーザー取得
		userData, err := model.GetUserByID(userid.(string))

		// エラー処理
		if err != nil {
			log.Println("failed to get user : " + err.Error())
			return ctx.NoContent(http.StatusUnauthorized)
		}

		log.Println("userid : " + userid.(string))
		log.Println(userData)

		// ユーザーIDをセット
		ctx.Set("userid", userid)

		return next(ctx)
	}
}
