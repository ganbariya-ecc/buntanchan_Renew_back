package model

import (
	"log"
	"os"
	// "os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	//データベース接続
	dbconn *gorm.DB = nil

	// 画像フォルダのパス
	TaskImageDir = "./assets/TaskImages"
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
	db.AutoMigrate(Task{})

	//グローバル変数に格納
	dbconn = db

	// タスク画像フォルダを作成する
	os.MkdirAll(TaskImageDir, 0644)

	// // デフォルトアイコンが存在するか
	// if _, err := os.Stat(DefaultUserIcon); err != nil {
	// 	log.Println(err)
	// 	log.Fatalln("User default icon does not exist")
	// }
}
