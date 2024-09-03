package service

import "auth/model"

func Admin_DeleteUser(adminid,userid string) error {
	// ユーザーを取得する
	delete_user,err := model.GetUserByID(userid)

	// エラー処理
	if err != nil {
		return err
	}

	// ユーザーを削除する
 	return delete_user.Delete()
}

func Admin_GetUser_Info(adminid,userid string) (model.User,error) {
	// ユーザーを取得する
	get_user,err := model.GetUserByID(userid)

	// エラー処理
	if err != nil {
		return model.User{},err
	}

	// ユーザーを返す
 	return get_user,nil
}