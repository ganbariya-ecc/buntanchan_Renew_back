package model

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	//データベース接続
	dbconn *gorm.DB = nil

	// ユーザーアイコンのフォルダ
	UserIconDir = "./assets/UserIcons"

	// Adminアイコンのフォルダ
	AdminIconDir = "./assets/AdminIcons"

	// デフォルトアイコン
	DefaultUserIcon = "./assets/defaults/usericon.jpg"
)

func Init(dbpath string) {
	log.Println("dbpath : " + dbpath)

	//データベース接続を開く
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})

	// エラー処理
	if err != nil {
		log.Fatalln("failed to Init Database : " + err.Error())
	}

	//マイグレーション
	db.AutoMigrate(User{})
	db.AutoMigrate(UserLabel{})
	db.AutoMigrate(AdminUser{})
	db.AutoMigrate(AdminLabel{})

	//グローバル変数に格納
	dbconn = db

	// ユーザーアイコンフォルダを作成する
	os.MkdirAll(UserIconDir, 0644)

	// Adminアイコンフォルダを作成する
	os.MkdirAll(AdminIconDir, 0644)

	// デフォルトアイコンが存在するか
	if _, err := os.Stat(DefaultUserIcon); err != nil {
		log.Println(err)
		log.Fatalln("User default icon does not exist")
	}
}
