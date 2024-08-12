package utils

import (
	"github.com/google/uuid"
)

func GenID() (string) {
	//ID生成
	genid, err := uuid.NewRandom()

	//エラー処理
    if err != nil {
        panic(err)
    }

	return genid.String()
}