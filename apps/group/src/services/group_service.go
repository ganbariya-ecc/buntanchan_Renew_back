package services

import (
	"errors"
	"group/model"
	"group/sdks/authsdk"
	"group/sdks/authsdk/protoc"
	"group/utils"
)

type GroupData struct {
	Name    string   `json:"name"`
	Members []Member `json:"members"`
}

type Member struct {
	Name    string `json:"name"`
	IsAdmin bool   `json:"admin"`
}

// グループを作成する関数
func CreateGroup(owner *protoc.User, data GroupData) error {
	// オーナーのグループを取得
	_, err := model.GetGroupByOwnerID(owner.UserID)

	// エラー処理
	if err != nil {
		// エラーの場合
		// グループを作成する

		// メンバーデータを作成する
		members := []model.MemberData{}

		// オーナーを追加
		members = append(members, model.MemberData{
			MemberID:   owner.UserID,
			Name:       owner.UserName,
			GroupID:    "",
			MemberRole: model.Owner,
		})

		// メンバーデータを渡す
		for _, val := range data.Members {
			//パスワード生成 (数字のみ)
			passwd, err := utils.GenPasswd(12)

			// エラー処理
			if err != nil {
				return err
			}

			// ユーザーを作成
			userid, err := authsdk.CreateUser(val.Name, passwd)

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
		_, err := model.CreatedGroup(owner.UserID, data.Name, members)

		return err
	} else {
		// グループが見つかった時
		return errors.New("Already created a group")
	}
}

func role_to_str(isAdmin bool) model.Role {
	if isAdmin {
		return model.Admin
	}

	return model.Member
}

type CurrentGroupData struct {
	Group  model.Group
	Mydata model.MemberData
}

func GetCurrentGroup(userid string) (CurrentGroupData, error) {
	// ユーザーを取得
	getUser, err := model.GetMember(userid)

	// エラー処理
	if err != nil {
		return CurrentGroupData{}, err
	}

	// グループを取得
	group, err := model.GetGroup(getUser.GroupID)

	// エラー処理
	if err != nil {
		return CurrentGroupData{}, err
	}

	// 自身を取得
	myData,err := model.GetMember(userid)

	// エラー処理
	if err != nil {
		return CurrentGroupData{},err
	}

	return CurrentGroupData{
		Group: group,
		Mydata: myData,
	}, nil
}

type MembersData struct {
	UserID   string
	UserName string
	Password string
	Role     model.Role
	Point    int64
}

func GetCurrentMembers(userid string) ([]MembersData, error) {
	// 所属しているグループを取得
	group, err := GetCurrentGroup(userid)

	// エラー処理
	if err != nil {
		return []MembersData{}, err
	}

	// メンバーを取得
	get_member, err := model.GetMember(userid)

	// エラー処理
	if err != nil {
		return []MembersData{}, err
	}

	// メンバー一覧を取得
	members, err := group.Group.GetMembers()

	// エラー処理
	if err != nil {
		return []MembersData{}, err
	}

	// 返すデータ
	return_data := []MembersData{}

	// メンバーを回す
	for _, val := range members {
		// 自身のID と一致する場合は飛ばす
		// if val.MemberID == userid {
		// 	continue
		// }

		switch get_member.MemberRole {
		case model.Owner:
			// ユーザの詳細を取得する
			userData, err := authsdk.GetUserAll(val.MemberID)

			// エラー処理
			if err != nil {
				return []MembersData{}, err
			}

			// データを追加
			return_data = append(return_data, MembersData{
				UserID:   userData.UserID,
				UserName: userData.UserName,
				Password: userData.Password,
				Role:     val.MemberRole,
				Point:    val.Point,
			})
		case model.Admin, model.Member:
			// データを追加
			return_data = append(return_data, MembersData{
				UserID:   val.MemberID,
				UserName: val.Name,
				Password: "",
				Role:     val.MemberRole,
				Point:    val.Point,
			})
		}
	}

	return return_data, nil
}
