package model

import (
	"group/utils"
	"time"
)

type Group struct {
	GroupID   string `gorm:"primaryKey"`
	Name      string
	OwnerID   string    // オーナーのユーザーID
	CreatedAt time.Time // GORMによって自動的に管理される作成時間
	UpdatedAt time.Time // GORMによって自動的に管理される更新時間

	Members []MemberData `gorm:"foreignKey:GroupID"` //メンバー
}

func CreatedGroup(ownerid,name string, members []MemberData) (string, error) {
	// グループID生成
	groupid := utils.GenID()

	// グループを作成する
	create_group := Group{
		GroupID: groupid,
		Name: name,
		OwnerID: ownerid,
		Members: members,
	}

	// データーベースに書き込む
	result := dbconn.Save(&create_group)

	// エラー処理
	if result.Error != nil {
		return "",result.Error
	}

	// 成功したとき
	return groupid,nil
}

// オーナーIDからグループを取得する
func GetGroupByOwnerID(ownerid string) (Group,error) {
	var returnData Group

	// グループを取得
	result := dbconn.Where(Group{
		OwnerID: ownerid,
	}).First(&returnData)

	return returnData,result.Error
}