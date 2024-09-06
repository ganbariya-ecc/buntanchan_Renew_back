package authsdk

import "log"

func Test() {
	// ユーザーを作成
	userid, err := CreateUser("wao", "password")

	// エラー処理
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(userid)
}