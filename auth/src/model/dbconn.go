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
)

func Init() {
	//データベース接続を開く
	db,err := gorm.Open(sqlite.Open(os.Getenv("DBPATH")),&gorm.Config{})

	// エラー処理
	if err != nil {
		log.Fatalln("failed to Init Database : " + err.Error())
	}

	//マイグレーション
	db.AutoMigrate(User{})
	db.AutoMigrate(UserLabel{})

	//グローバル変数に格納
	dbconn = db
}