package controller

import (
	"group/sdks/authsdk/protoc"
	"group/services"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateGroup(ctx echo.Context) error {
	// ユーザー取得
	user := ctx.Get("user").(*protoc.User)

	//値をBind する
	var data services.GroupData
	if err := ctx.Bind(&data); err != nil {
		log.Println("Failed to bind value : " + err.Error())
		return ctx.JSON(http.StatusBadRequest,echo.Map{
			"result" : "Failed to bind value",
		})
	}
	
	// log.Println(user)
	// log.Println(data)

	// ユーザーを検証
	if user.AuthType != "Oauth2" {
		//Oauth ユーザーじゃなかったら
		return ctx.JSON(http.StatusForbidden,echo.Map{
			"result" : "This account cannot create groups",
		})
	}


	// グループを作成
	err := services.CreateGroup(user,data)

	// エラー処理
	if err != nil {
		log.Println("Creation failure : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError,echo.Map{
			"result" : "Creation failure",
		})
	}

	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : "success",
	})
}

// 所属グループ取得
func GetCurrentGroup(ctx echo.Context) error {
	// ユーザー取得
	user := ctx.Get("user").(*protoc.User)

	// 所属しているグループを取得
	group,err := services.GetCurrentGroup(user.UserID)

	// エラー処理
	if err != nil {
		log.Println("get current group Error : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError,echo.Map{
			"result" : "get current group failed",
		})
	}

	// メンバー情報をnil にする
	group.Members = nil

	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : group,
	})
}

func GetCurrentMembers(ctx echo.Context) error {
	// ユーザー取得
	user := ctx.Get("user").(*protoc.User)

	// メンバー一覧取得
	result,err := services.GetCurrentMembers(user.UserID)

	// エラー処理
	if err != nil {
		log.Println("Get Current Members failed : " + err.Error())

		return ctx.JSON(http.StatusInternalServerError,echo.Map{
			"result" : "Get Current Members failed",
		})
	}

	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : result,
	})
}