package middlewares

import (
	"group/sdks/authsdk"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Header 取得
		btoken := ctx.Request().Header.Get("Authorization")

		log.Println("token : ",btoken)

		// トークンがない場合
		if btoken == "" {
			return ctx.JSON(http.StatusUnauthorized,echo.Map{
				"result" : "There is no token",
 			})
		}

		// ユーザー取得
		user,err := authsdk.Auth(btoken)

		// エラー処理
		if err != nil {
			log.Println("Authentication failed : " + err.Error())
			return ctx.JSON(http.StatusUnauthorized,echo.Map{
				"result" : "failed to auth",
			})
		}

		// ユーザーデータ
		ctx.Set("user",&user)

		return next(ctx)
	}
}