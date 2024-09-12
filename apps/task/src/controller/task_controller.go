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
	log.Println(data.ExpirationDate)
	log.Println(utils.NowTime().Unix())
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
	taskid,err := services.CreateTask(data,creator_data)

	// エラー処理
	if err != nil {
		log.Println("failed to create task : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"result": "failed to create task",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": taskid,
	})
}

func GetTasks(ctx echo.Context) error {
	// 作成者のデータ
	member_data := ctx.Get("member").(protoc.MemberData)

	// 自分向けのタスクを取得する
	tasks,err := services.GetTasks(member_data)

	// エラー処理
	if err != nil {
		log.Println("Failed to get task : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError,echo.Map{
			"result" : "Failed to get task",
		})
	}

	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : tasks,
	})
}

func GetTask(ctx echo.Context) error {
	// 取得者のデータ
	member_data := ctx.Get("member").(protoc.MemberData)

	// TaskID 取得
	taskid := ctx.Request().Header.Get("taskid")

	// タスクの詳細を取得
	task,err := services.GetTask(member_data,taskid)

	// エラー処理
	if err != nil {
		log.Println("Failed to get task : " + err.Error())
		return ctx.JSON(http.StatusInternalServerError,echo.Map{
			"result" : "Failed to get task",
		})
	}

	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : task,
	})
}

func UploadTaskImg(ctx echo.Context) error {
	// 取得者のデータ
	member_data := ctx.Get("member").(protoc.MemberData)

	// TaskID 取得
	taskid := ctx.Request().Header.Get("taskid")

	// タスクIDがない場合
	if (taskid == "") {
		return ctx.JSON(http.StatusBadRequest,echo.Map{
			"result" : "Failed to get file",
		})
	}

	// 画像を取得
	imgFile,err := ctx.FormFile("img")

	// エラー処理
	if err != nil {
		log.Println("Failed to get file")
		return ctx.JSON(http.StatusBadRequest,echo.Map{
			"result" : "Failed to get file",
		})
	}

	// 画像をアップロードする
	herr := services.UploadTaskImg(imgFile,member_data,taskid)

	if herr != nil {
		log.Println(herr.Message)
		return ctx.JSON(herr.Status,echo.Map{
			"result" : herr.Message,
		})
	}

	return ctx.NoContent(http.StatusOK)
}

func GetTaskImg(ctx echo.Context) error {
	// 取得者のデータ
	member_data := ctx.Get("member").(protoc.MemberData)

	// TaskID 取得
	taskid := ctx.Request().Header.Get("taskid")

	// タスク画像を取得する
	imgPath,err := services.GetImage(taskid,member_data)

	// エラー処理
	if err != nil {
		//404 の場合
		if err.Status == http.StatusNotFound {
			return ctx.File("./assets/defaultImages/404.jpg")
		}
	
		log.Println("failed to get image : " + err.Error())
		return ctx.JSON(err.Status,err.Message)
	}

	// 画像を返す
	return ctx.File(imgPath)
}