package model

import "task/utils"

type Status string

const (
	Rejected   = Status("Rejected")
	InProgress = Status("InProgress")
	Reported   = Status("Reported")
	Completed  = Status("Completed")
)

type Task struct {
	TaskID         string `gorm:"primaryKey"`
	TaskName       string //タスク名
	GroupID        string // タスクの属しているグループID
	CreatorID      string //作成者ID
	Explanation    string // タスクの説明
	ExpirationDate int64  // タスクの有効期限
	OrderTargetID  string // 依頼先ID
	Point          int    //タスクのポイント
	Status         Status //タスクの状態
}

func CreateTask(TaskName, GroupID, CreatorID, Explanation, OrderTargetID string, ExpirationDate int64,Point int) (string, error) {
	// タスクID生成
	taskid := utils.GenID()

	// 作成するタスク
	taskData := Task{
		TaskID:         taskid,
		TaskName:       TaskName,
		GroupID:        GroupID,
		CreatorID:      CreatorID,
		Explanation:    Explanation,
		ExpirationDate: ExpirationDate,
		OrderTargetID:  OrderTargetID,
		Status:         InProgress,
	}

	// データを保存
	result := dbconn.Save(&taskData)

	// エラー処理
	if result.Error != nil {
		return "", result.Error
	}

	return taskid, nil
}

func DeleteTask(taskid string) error {
	// データを保存
	result := dbconn.Where(Task{
		TaskID: taskid,
	}).Delete(&Task{})

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// タスクのステータスを更新
func (task *Task) UpdateStatus(status Status) error {
	// ステータス更新
	task.Status = status

	// データ更新
	result := dbconn.Save(task)

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	return nil
}
