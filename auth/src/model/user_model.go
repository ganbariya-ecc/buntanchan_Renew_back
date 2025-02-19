package model

import (
	"auth/utils"
	"errors"
	"log"
	"os"
	"path"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AType string

const (
	AuthType_O = AType("Oauth2")
	AuthType_B = AType("Basic")
)

type User struct {
	//ユーザーの情報
	UserID   string      `gorm:"primaryKey"` //ユーザーID
	UserName string      //ユーザー名
	Labels   []UserLabel `gorm:"foreignKey:UID"` //ラベル

	AuthType AType `gorm:"not null"` //認証の種類

	// Oauth 関連
	Email    string //メールアドレス
	Provider string //プロバイダー
	ProviderID string //プロバイダーID

	// 認証関連
	Password string //パスワード

	//その他
	CreatedAt time.Time //作成日時
	UpdatedAt time.Time //更新日時

	HashPassword bool // パスワードをハッシュ化するか
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

// Oauth のユーザーを取得
func GetOauthUser(provider string,providerid string) (User, error) {
	// ユーザーデータを格納する変数
	UserData := User{}

	// ユーザーを取得
	result := dbconn.Where(&User{
		Provider: provider,
		ProviderID: providerid,
	}).First(&UserData)

	// エラー処理
	if result.Error != nil {
		return User{}, result.Error
	}

	return UserData, nil
}

// Oauth のユーザーを作成
func CreateOauthUser(UserName string,ProviderUserid string, Labels []UserLabel, email string, provider string) (string, error) {
	//Oauth ユーザーを取得
	ouser,err := GetOauthUser(provider,ProviderUserid)

	//エラー処理
	if err == nil {
		return ouser.UserID, nil
	}

	//ユーザーID生成
	UserID := utils.GenID()

	//ユーザーデータ作成
	UserData := User{
		UserID:   UserID,
		UserName: UserName,
		Labels:   Labels,
		AuthType: AuthType_O,
		Email:    email,
		ProviderID: ProviderUserid,
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

	//ファイルをコピーする

	// アイコンをコピーする
	err = utils.CopyFile(DefaultUserIcon,path.Join(UserIconDir,UserID + ".jpeg"))

	return UserID, err
}

// パスワード認証ユーザーを作成
func CreateUser(UserName string, Labels []UserLabel, Password string,DoHash bool) (string, error) {
	//ユーザーID生成
	UserID := utils.GenID()

	// パスワード
	passwd := Password

	// パスワードをハッシュ化するか
	if DoHash {
	//パスワードをハッシュ化
		hashed, err := utils.HashPassword(Password)	

		//エラー処理
		if err != nil {
			return "", err
		}

		// パスワードを設定
		passwd = hashed
	}

	//ユーザーデータ作成
	UserData := User{
		UserID:   UserID,
		UserName: UserName,
		Labels:   Labels,
		AuthType: AuthType_B,
		Email:    "",
		ProviderID: "",
		Provider: "",
		Password: passwd,
		HashPassword: DoHash,
	}

	//データベースに保存
	result := dbconn.Save(&UserData)

	//エラー処理
	if result.Error != nil {
		log.Println("failed to create user : " + result.Error.Error())
		return "", result.Error
	}

	// アイコンをコピーする
	err := utils.CopyFile(DefaultUserIcon,path.Join(UserIconDir,UserID + ".jpeg"))

	return UserID, err
}

// ユーザーIDでユーザを取得する
func GetUserByID(userid string) (User,error) {
	// バリデーション
	if userid == "" {
		return User{},errors.New("userid is empty")
	}

	// ユーザーデータを格納する変数
	UserData := User{}

	// ユーザーデータ取得
	result := dbconn.Where(&User{UserID: userid}).First(&UserData)

	// エラー処理
	if result.Error != nil {
		return UserData,result.Error
	}

	return UserData,nil
}

//ラベルを取得する関数
func (usr User) GetLabels() ([]UserLabel,error) {
	//ラベルを格納する変数
	var labels []UserLabel

	// ラベルを取得
	err := dbconn.Model(usr).Association("Labels").Find(&labels)
	return labels,err
}

// ユーザーを削除する
func (usr User) Delete() (error) {
	// ラベルを削除する
	result := dbconn.Where(&UserLabel{
		UID: usr.UserID,
	}).Unscoped().Delete(&UserLabel{})

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	// ユーザーアイコンを削除する
	err := os.Remove(path.Join(UserIconDir,usr.UserID + ".jpeg"))

	// エラー処理
	if err != nil {
		return err
	}

	// ユーザーを削除する
	result = dbconn.Where(&User{
		UserID: usr.UserID,
	}).Unscoped().Delete(&User{})

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (usr User) ValidatePassword(Password string) (bool) {
	//ハッシュ化されたパスワードと入力されたパスワードを比較
	return bcrypt.CompareHashAndPassword([]byte(usr.Password),[]byte(Password)) == nil
}

func GetAllUser() ([]User,error) {
	var users []User

	// 全ユーザー取得
	result := dbconn.Find(&users)

	// エラー処理
	if result.Error != nil {
		return []User{},result.Error
	}

	return users,nil
}