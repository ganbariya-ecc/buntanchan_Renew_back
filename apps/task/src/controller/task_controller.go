package controller

import (
	"log"
	"net/http"
	"task/sdks/groupsdk/protoc"
	"task/services"
	"task/utils"

	"github.com/labstack/echo/v4"
)


func CreateTask(ctx echo.Context) error {
	// データバインド
	var data services.TaskData

	// 値をバインドする
	err := ctx.Bind(&data)

	// エラー処理
	if err != nil {
		log.Println("Failed to bind value : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"result": "Failed to bind value",
		})
	}

	// 作成者のデータ
	creator_data := ctx.Get("member").(protoc.MemberData)

	// バリデーション 
	// 現在時刻より過去の場合
	if data.ExpirationDate < utils.NowTime().Unix() {
		log.Println("create task error : The specified time is in the past")
		return ctx.JSON(http.StatusBadRequest,echo.Map{
			"result" : "The specified time is in the past",
		})
	}

	// タスク名がないとき
	if data.TaskName == "" {
		log.Println("Missing task name")
		return ctx.JSON(http.StatusBadRequest,echo.Map{
			"result" : "Missing task name",
		})
	}

	// タスクのポイントがマイナスの時
	if data.Point < 0 {
		log.Println("Invalid points for task")
		return ctx.JSON(http.StatusBadRequest,echo.Map{
			"result" : "Invalid points for task",
		})
	}

	// タスクを作成する
	err = services.CreateTask(data,creator_data)

	// エラー処理
	if err != nil {
		log.Println("failed to create task : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"result": "failed to create task",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "ok",
	})
}
