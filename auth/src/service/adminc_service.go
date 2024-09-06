package service

import (
	"auth/model"
)

func AdminGetInfo(userid string) (model.AdminUser, error) {
	// admin ユーザーを取得
	return model.GetAdminByID(userid)
}

func AdminGetUsers(userid string) ([]model.User, error) {
	// admin ユーザー取得
	adminUser, err := model.GetAdminByID(userid)

	// エラー処理
	if err != nil {
		return []model.User{}, err
	}

	// ラベル取得
	_,err = adminUser.GetLabel("Owner")

	// エラー処理
	if err != nil {
		return []model.User{},err
	}

	// ユーザー一覧取得
	return model.GetAllUser()
}

func AdminGetUserInfo(adminid,userid string) (model.User,error) {
	// ユーザー情報取得
	userData,err := model.GetUserByID(userid)

	// エラー処理
	if err != nil {
		return model.User{},err
	}

	return userData,nil
}