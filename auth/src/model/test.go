package model

import (
	"log"
)

func Test() {
	// ファイルを削除する
	// os.Remove("./auth.db")

	// ユーザーを作成する
	uid1,err := CreateOauthUser("wao",[]UserLabel{
		CreateLabel("hello"),
		CreateLabel("world"),
	},"admin@example.com","discord")

	if err != nil {
		log.Println(err)
	}

	log.Println(uid1)
}