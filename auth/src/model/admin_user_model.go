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

type AdminUser struct {
	//ユーザーの情報
	UserID   string      `gorm:"primaryKey"` //ユーザーID
	UserName string      //ユーザー名
	Labels   []AdminLabel `gorm:"foreignKey:UID"` //ラベル

	// 認証関連
	Password string //パスワードハッシュ

	//その他
	CreatedAt time.Time //作成日時
	UpdatedAt time.Time //更新日時
}

type AdminLabel struct {
	// LabelID string `gorm:"primaryKey"`
	UID     string	
	Name    string `gorm:"primaryKey"`
}

func CreateAdminLabel(name string) AdminLabel {
	// 管理者ラベルを生成
	return AdminLabel{
		Name:    name,
	}
}

// Admin ユーザーを作成
func CreateAdminUser(UserName string, Labels []AdminLabel, Password string) (string, error) {
	//ユーザーID生成
	UserID := utils.GenID()

	//パスワードをハッシュ化
	hashed, err := utils.HashPassword(Password)

	//エラー処理
	if err != nil {
		return "", err
	}

	//ユーザーデータ作成
	UserData := AdminUser{
		UserID:   UserID,
		UserName: UserName,
		Labels:   Labels,
		Password: hashed,
	}

	//データベースに保存
	result := dbconn.Save(&UserData)

	//エラー処理
	if result.Error != nil {
		log.Println("failed to create user : " + result.Error.Error())
		return "", result.Error
	}

	// アイコンをコピーする
	err = utils.CopyFile(DefaultUserIcon,path.Join(AdminIconDir,UserID + ".jpeg"))

	return UserID, err
}

// ユーザーIDでユーザを取得する
func GetAdmin(userName string) (AdminUser,error) {
	// バリデーション
	if userName == "" {
		return AdminUser{},errors.New("userName is empty")
	}

	// Adminデータを格納する変数
	UserData := AdminUser{}

	// Adminデータ取得
	result := dbconn.Where(&AdminUser{UserName: userName}).First(&UserData)

	// エラー処理
	if result.Error != nil {
		return UserData,result.Error
	}

	return UserData,nil
}

// ユーザーIDでユーザを取得する
func GetAdminByID(userid string) (AdminUser,error) {
	// バリデーション
	if userid == "" {
		return AdminUser{},errors.New("userid is empty")
	}

	// Adminデータを格納する変数
	UserData := AdminUser{}

	// Adminデータ取得
	result := dbconn.Where(&AdminUser{UserID: userid}).First(&UserData)

	// エラー処理
	if result.Error != nil {
		return UserData,result.Error
	}

	return UserData,nil
}

//ラベルを取得する関数
func (usr AdminUser) GetLabels() ([]AdminLabel,error) {
	//ラベルを格納する変数
	var labels []AdminLabel

	// ラベルを取得
	err := dbconn.Model(usr).Association("Labels").Find(&labels)
	return labels,err
}

// ユーザーを削除する
func (usr AdminUser) Delete() (error) {
	// ラベルを削除する
	result := dbconn.Where(&AdminLabel{
		UID: usr.UserID,
	}).Unscoped().Delete(&AdminLabel{})

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	// ユーザーアイコンを削除する
	err := os.Remove(path.Join(AdminIconDir,usr.UserID + ".jpeg"))

	// エラー処理
	if err != nil {
		return err
	}

	// ユーザーを削除する
	result = dbconn.Where(&AdminUser{
		UserID: usr.UserID,
	}).Unscoped().Delete(&AdminUser{})

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (usr AdminUser) ValidatePassword(Password string) (bool) {
	//ハッシュ化されたパスワードと入力されたパスワードを比較
	return bcrypt.CompareHashAndPassword([]byte(usr.Password),[]byte(Password)) == nil
}

func GetAdmins() ([]AdminUser,error) {
	// Admin を格納する変数
	var admins []AdminUser

	// すべてのユーザーを取得する
	result := dbconn.Find(&admins)

	// エラー処理
	if result.Error != nil {
		return []AdminUser{},result.Error
	}

	return admins,nil
}