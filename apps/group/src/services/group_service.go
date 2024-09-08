package services

import (
	"errors"
	"group/model"
	"group/sdks/authsdk"
	"group/sdks/authsdk/protoc"
	"group/utils"
)

type GroupData struct {
	Name    string `json:"name"`
	Members []Member `json:"members"`
}

type Member struct {
	Name    string `json:"name"`
	IsAdmin bool   `json:"admin"`
}

// グループを作成する関数
func CreateGroup(owner *protoc.User,data GroupData) (error) {
	// オーナーのグループを取得
	_,err := model.GetGroupByOwnerID(owner.UserID)

	// エラー処理
	if err != nil {
		// エラーの場合
		// グループを作成する

		// メンバーデータを作成する
		members := []model.MemberData{}

		// メンバーデータを渡す
		for _, val := range data.Members {
			//パスワード生成 (数字のみ)
			passwd,err := utils.GenPasswd(12)
			
			// エラー処理
			if err != nil {
				return err
			}

			// ユーザーを作成
			userid,err := authsdk.CreateUser(val.Name,passwd)

			// エラー処理
			if err != nil {
				return err
			}

			// メンバー情報を追加
			members = append(members, model.MemberData{
				MemberID:   userid,
				Name:       val.Name,
				GroupID:    "",
				MemberRole: role_to_str(val.IsAdmin),
			})
		}

		// グループを作成する
		_,err := model.CreatedGroup(owner.UserID,data.Name,members)

		return err
	} else {
		// グループが見つかった時
		return errors.New("Already created a group")
	}
}

func role_to_str(isAdmin bool) model.Role {
	if (isAdmin) {
		return model.Admin
	}

	return model.Member
}