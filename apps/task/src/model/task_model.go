package model

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
	Status         Status //タスクの状態
}
