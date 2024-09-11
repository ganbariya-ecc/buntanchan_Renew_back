package services

import (
	"errors"
	"log"
	"task/model"
	"task/sdks/groupsdk"
	"task/sdks/groupsdk/protoc"
)

type TaskData struct {
	TaskName       string //タスク名
	Explanation    string // タスクの説明
	ExpirationDate int64  // タスクの有効期限
	OrderTargetID  string // 依頼先ID
	Point          int    //タスクのポイント
}

func CreateTask(data TaskData,Creator protoc.MemberData) error {
	// ターゲットメンバーを取得
	target_member,err := groupsdk.GetMember(data.OrderTargetID)

	// エラー処理
	if err != nil {
		return err
	}

	// ターゲットのグループと自分のグループが違うとき
	if target_member.GroupID != Creator.GroupID {
		return errors.New("Unable to create tasks for other groups")
	}

	// タスクを作成する
	create_task,err := model.CreateTask(data.TaskName,Creator.GroupID,Creator.Memberid,data.Explanation,data.OrderTargetID,data.ExpirationDate,data.Point)

	// エラー処理
	if err != nil {
		return err
	}

	log.Println(create_task)

	return nil
}