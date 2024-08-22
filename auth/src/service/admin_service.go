package service

import (
	"auth/model"
	"errors"
	"log"
)

func AdminLogin(userName string, password string) (model.AdminUser,error) {
	//ユーザーデータを取得
	UserData,err := model.GetAdmin(userName)

	//エラー処理
	if err != nil {
		return model.AdminUser{},err
	}

	// パスワード認証
	if !UserData.ValidatePassword(password) {
		return model.AdminUser{},errors.New("password or username is wrong")
	}

	return UserData, nil
}

func AdminSignup(userName string, password string) (model.AdminUser,error) {
	// admin 一覧取得
	admins,err := model.GetAdmins()

	// エラー処理
	if err != nil {
		log.Println("failed to get admins : " + err.Error())
		return model.AdminUser{},errors.New("failed to get admins")
	}

	// ユーザーが存在しない場合
	if len(admins) == 0 {
		//ユーザーデータ作成
		userid,err := model.CreateAdminUser(userName,[]model.AdminLabel{
			{
				Name: "Owner",
			},
		},password)

		//エラー処理
		if err != nil {
			log.Println("failed to create user : " + err.Error())
			return model.AdminUser{},errors.New("failed to create user")
		}

		// ユーザー取得
		return model.GetAdminByID(userid)
	} 

	// Admin が存在する場合
	return model.AdminUser{},errors.New("Sign up is prohibited")
}