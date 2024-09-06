package middleware

import (
	"group/sdks/authsdk"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Header 取得
		btoken := ctx.Request().Header.Get("Authorization")

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
			return err
		}

		// ユーザーID
		ctx.Set("user",user)

		return nil
	}
}