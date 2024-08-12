package service

import (
	"auth/model"

	"github.com/markbates/goth"
)

// Oauth の完了
func CallbackOauth(user goth.User) (string,error) {
	userName := ""

	// ユーザー名が空かどうか
	if user.Name != "" {
		// ユーザー名
		userName = user.Name
	} else if user.FirstName != "" {
		// ファーストネーム設定
		userName := user.FirstName

		// ラスト名設定
		if user.LastName != "" {
			//ラスト名設定
			userName += user.LastName
		}
	} else {
		//ニックネームに設定
		userName = user.NickName
	}

	// ユーザー作成
	userid,err := model.CreateOauthUser(userName,[]model.UserLabel{},user.Email,user.Provider)

	// エラー処理
	if err != nil {
		return "",err
	}

	return userid,nil
}