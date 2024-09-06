package service

import (
	"auth/model"
	"errors"
	"log"
)

func BasicLogin(userid string, password string) (model.User,error) {
	//ユーザーデータを取得
	UserData,err := model.GetUserByID(userid)

	//エラー処理
	if err != nil {
		return model.User{},err
	}

	// Oauth ユーザーの場合
	if UserData.AuthType == model.AuthType_O {
		return model.User{},errors.New("Oauthユーザーは認証できません")
	}

	// パスワード認証
	if !UserData.ValidatePassword(password) {
		return model.User{},errors.New("password or userid is wrong")
	}

	return UserData, nil
}

func BasicSignup(userName string, password string) (model.User,error) {
	//ユーザーデータ作成
	userid,err := model.CreateUser(userName,[]model.UserLabel{},password)

	//エラー処理
	if err != nil {
		log.Println("failed to create user : " + err.Error())
		return model.User{},errors.New("failed to create user")
	}

	// ユーザー取得
	return model.GetUserByID(userid)
}