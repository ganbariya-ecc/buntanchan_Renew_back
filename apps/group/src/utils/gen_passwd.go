package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strconv"
)

func GenPasswd(length int) (string,error) {
	total_str := ""

	// 指定回数回す
	for i := 0; i < length; i++ {
		// 一桁を生成する
		passwd, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "",errors.New("Failed to generate random number")
		}

		// 文字列にして結合
		total_str += strconv.Itoa(int(passwd.Int64()))
	}

	return total_str,nil
}