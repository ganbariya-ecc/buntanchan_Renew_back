package model

import "time"

type Role string

const (
	Owner  = Role("Owner")
	Admin  = Role("Admin")
	Member = Role("Member")
)

type MemberData struct {
	MemberID   string    `gorm:"primaryKey"`
	Name       string    // メンバー名
	GroupID    string    //種族してるグループID
	CreatedAt  time.Time // GORMによって自動的に管理される作成時間
	UpdatedAt  time.Time // GORMによって自動的に管理される更新時間
	MemberRole Role      // グループの役職
}
