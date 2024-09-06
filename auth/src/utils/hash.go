package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string,error) {
	//パスワードのバイナリ生成
	passwd_bin := []byte(password)

	//ハッシュ化
	hashed,err := bcrypt.GenerateFromPassword(passwd_bin,10)

	//エラー処理
	if err != nil {
		return "",err
	}

	return string(hashed),nil
}