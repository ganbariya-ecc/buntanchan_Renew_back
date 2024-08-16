package service

import (
	"auth/utils"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenJWT(userid string) (string,*utils.HttpError) {
	// JWT の有効期限取得
	ExpirySec := os.Getenv("EXPIRYSEC")

	// 数字に変換
	ExpirySecInt,err := strconv.Atoi(ExpirySec)

	// エラー処理
	if err != nil {
		return "",utils.NewHttpError(http.StatusInternalServerError,"failed to convert EXPIRYSEC")
	}

	// JWT作成
	token := jwt.NewWithClaims(SignMethod,jwt.MapClaims{
		"userid" : userid,
		"exp" : utils.NowTime().Add(time.Second * time.Duration(ExpirySecInt)).Unix(),
	})

	// jwt 署名
	tokenString,err := token.SignedString([]byte(JWT_KEY))

	// エラー処理
	if err != nil {
		return "",utils.NewHttpError(http.StatusInternalServerError,"failed to sign token")
	}

	return tokenString,nil
}