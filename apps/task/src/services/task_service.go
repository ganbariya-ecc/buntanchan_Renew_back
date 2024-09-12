package services

import (
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"path"
	"task/model"
	"task/sdks/groupsdk"
	"task/sdks/groupsdk/protoc"
	"task/utils"
)

type TaskData struct {
	TaskName       string //タスク名
	Explanation    string // タスクの説明
	ExpirationDate int64  // タスクの有効期限
	OrderTargetID  string // 依頼先ID
	Point          int    //タスクのポイント
}

func CreateTask(data TaskData,Creator protoc.MemberData) (string,error) {
	// ターゲットメンバーを取得
	target_member,err := groupsdk.GetMember(data.OrderTargetID)

	// エラー処理
	if err != nil {
		return "",err
	}

	// ターゲットのグループと自分のグループが違うとき
	if target_member.GroupID != Creator.GroupID {
		return "",errors.New("Unable to create tasks for other groups")
	}

	// タスクを作成する
	taskid,err := model.CreateTask(data.TaskName,Creator.GroupID,Creator.Memberid,data.Explanation,data.OrderTargetID,data.ExpirationDate,data.Point)

	// エラー処理
	if err != nil {
		return "",err
	}

	log.Println(taskid)

	return taskid,nil
}

func GetTasks(getUser protoc.MemberData) ([]model.Task,error) {
	// グループからタスクを取得
	tasks,err := model.GetTasks(getUser.GroupID)

	// エラー処理
	if err != nil {
		return []model.Task{},err
	}

	// オーナーか Admin の場合タスクを返す
	if getUser.MemberRole == "Owner" || getUser.MemberRole == "Admin" {
		return tasks,nil
	}

	// 戻す引数
	return_tasks := []model.Task{}

	// メンバーの場合
	for _, val := range tasks {
		// 自分向けのタスクの場合
		if val.OrderTargetID == getUser.Memberid {
			return_tasks = append(return_tasks, val)
		}
	}

	return return_tasks,nil
}

func GetTask(getUser protoc.MemberData,taskid string) (model.Task,error) {
	// グループからタスクを取得
	task,err := model.GetTask(taskid)

	// エラー処理
	if err != nil {
		return model.Task{},err
	}

	// オーナーか Admin の場合タスクを返す
	if getUser.MemberRole == "Owner" || getUser.MemberRole == "Admin" {
		return task,nil
	}

	// 自分向けのタスクの場合
	if task.OrderTargetID == getUser.Memberid {
		return task,nil
	}

	return model.Task{},errors.New("Failed to get task")
}

func UploadTaskImg(img *multipart.FileHeader,uploader protoc.MemberData,taskID string) *utils.HttpError {
	// タスクを取得
	task,err := model.GetTask(taskID)

	// エラー処理
	if err != nil {
		return utils.NewHttpError(http.StatusNotFound,"Task Not Found")
	}

	// 自分のタスクではない場合
	if task.CreatorID != uploader.Memberid {
		return utils.NewHttpError(http.StatusForbidden,"No permission")
	}

	// ファイルを開く
	openFile,err := img.Open()

	// エラー処理
	if err != nil {
		return utils.NewHttpError(http.StatusBadRequest,"Could not open file")
	}

	// 画像を読み込む
	imgData,err := utils.LoadStream(openFile)

	// エラー処理
	if err != nil {
		return utils.NewHttpError(http.StatusInternalServerError,"Could not load file")
	}
	
	// 画像をリサイズする
	resizeImg := utils.ResizeImage(imgData.Src,1280,720)

	// 画像を保存する
	err = utils.SaveImage(path.Join(model.TaskImageDir,taskID + ".jpeg"),resizeImg)

	// エラー処理
	if err != nil {
		return utils.NewHttpError(http.StatusInternalServerError,"failed to save image")
	}

	return nil
}