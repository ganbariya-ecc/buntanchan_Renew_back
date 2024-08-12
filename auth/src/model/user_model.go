package model

import (
	"auth/utils"
	"log"
	"time"
)

type AType string

const (
	AuthType_O = AType("Oauth2")
	AuthType_P = AType("Password")
)

type User struct {
	//ユーザーの情報
	UserID   string      `gorm:"primaryKey"` //ユーザーID
	UserName string      //ユーザー名
	Labels   []UserLabel `gorm:"foreignKey:UID"` //ラベル

	AuthType AType `gorm:"not null"` //認証の種類

	// Oauth 関連
	Email    string //メールアドレス
	Provider string `gorm:"default:'None'"` //プロバイダー

	// 認証関連
	Password string //パスワードハッシュ

	//その他
	CreatedAt time.Time //作成日時
	UpdatedAt time.Time //更新日時
}

type UserLabel struct {
	LabelID string `gorm:"primaryKey"`
	UID     string
	Name    string `gorm:"not null"`
}

func CreateLabel(name string) UserLabel {
	//ラベルID生成
	LabelID := utils.GenID()

	// ユーザーモデル作成
	return UserLabel{
		LabelID: LabelID,
		Name:    name,
	}
}

// Oauth のユーザーを作成
func CreateOauthUser(UserName string, Labels []UserLabel, email string, provider string) (string, error) {
	//ユーザーID生成
	UserID := utils.GenID()

	//ユーザーデータ作成
	UserData := User{
		UserID:   UserID,
		UserName: UserName,
		Labels:   Labels,
		AuthType: AuthType_O,
		Email:    email,
		Provider: provider,
		Password: "",
	}

	//データベースに保存
	result := dbconn.Save(&UserData)

	//エラー処理
	if result.Error != nil {
		log.Println("failed to create oauth2 user : " + result.Error.Error())
		return "", result.Error
	}

	return UserID, nil
}

// パスワード認証ユーザーを作成
func CreateUser(UserName string, Labels []UserLabel, Password string) (string, error) {
	//ユーザーID生成
	UserID := utils.GenID()

	//パスワードをハッシュ化
	hashed, err := utils.HashPassword(Password)

	//エラー処理
	if err != nil {
		return "", err
	}

	//ユーザーデータ作成
	UserData := User{
		UserID:   UserID,
		UserName: UserName,
		Labels:   Labels,
		AuthType: AuthType_P,
		Email:    "",
		Provider: "",
		Password: hashed,
	}

	//データベースに保存
	result := dbconn.Save(&UserData)

	//エラー処理
	if result.Error != nil {
		log.Println("failed to create user : " + result.Error.Error())
		return "", result.Error
	}

	return UserID, nil
}
