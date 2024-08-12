package model

import (
	"log"
	"os"
)

func Test() {
	// ファイルを削除する
	// os.Remove("./auth.db")

	// ユーザーアイコンフォルダを削除
	os.RemoveAll(UserIconDir)

	// アイコンフォルダを作成
	os.MkdirAll(UserIconDir,0644)

	// Oauth ユーザーを作成する
	uid1, err := CreateOauthUser("wao", []UserLabel{
		CreateLabel("hello"),
		CreateLabel("world"),
	}, "admin@example.com", "discord")

	if err != nil {
		log.Println(err)
	}

	log.Println("uid1 : " + uid1)

	// 一般ユーザーを作成する
	uid2, err := CreateUser("wao", []UserLabel{
		CreateLabel("hellob"),
		CreateLabel("worldb"),
	}, "password")

	if err != nil {
		log.Println(err)
	}

	log.Println("uid2 : " + uid2)

	// 一般ユーザー3を作成する
	uid3, err := CreateUser("wao3", []UserLabel{
		CreateLabel("hellob3"),
		CreateLabel("worldb3"),
	}, "password3")

	if err != nil {
		log.Println(err)
	}

	log.Println("uid3 : " + uid3)

	//ユーザー1 を取得する
	user1,err := GetUserByID(uid1)

	// エラー処理
	if err != nil {
		log.Fatalln("failed to get user1 : " + err.Error())
	}

	log.Println("user1")
	log.Println(user1)

	// ユーザー1 のラベルを取得
	usr1_labels,err := user1.GetLabels()

	// エラー処理
	if err != nil {
		log.Fatalln("failed to get user1 labels : " + err.Error())
	}

	log.Println("user1 labels")
	log.Println(usr1_labels)

	//ユーザー2を取得する
	user2,err := GetUserByID(uid2)

	// エラー処理
	if err != nil {
		log.Fatalln("failed to get user2 : " + err.Error())
	}

	log.Println("user2")
	log.Println(user2)

	// ユーザー1 のラベルを取得
	usr2_labels,err := user2.GetLabels()

	// エラー処理
	if err != nil {
		log.Fatalln("failed to get user2 labels : " + err.Error())
	}

	log.Println("user2 labels")
	log.Println(usr2_labels)

	//ユーザー3 を取得する
	user3,err := GetUserByID(uid3)

	// エラー処理
	if err != nil {
		log.Fatalln("failed to get user3 : " + err.Error())
	}

	// ユーザー3 を削除する
	log.Println(user3.Delete())
}
