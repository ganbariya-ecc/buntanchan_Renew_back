package model

import "time"

type Role string

const (
	Owner  = Role("Owner")
	Admin  = Role("Admin")
	Member = Role("Member")
)

type MemberData struct {
	MemberID   string    `gorm:"primaryKey"` //ユーザーID
	Name       string    // メンバー名
	GroupID    string    //種族してるグループID
	CreatedAt  time.Time // GORMによって自動的に管理される作成時間
	UpdatedAt  time.Time // GORMによって自動的に管理される更新時間
	MemberRole Role      // グループの役職
	Point      int64     `gorm:"default:0"` // 獲得ポイント数
}

// MemberID (UserID) でメンバーを取得する
func GetMember(memberid string) (MemberData, error) {
	var getData MemberData

	// 1件取得
	result := dbconn.Where(MemberData{
		MemberID: memberid,
	}).First(&getData)

	// エラー処理
	if result.Error != nil {
		return MemberData{}, result.Error
	}

	return getData, nil
}

func (group *Group) GetMembers() ([]MemberData, error) {
	var members []MemberData

	// メンバー一覧を取得
	err := dbconn.Model(group).Association("Members").Find(&members)

	// エラー処理
	if err != nil {
		return []MemberData{}, err
	}

	return members, nil
}
