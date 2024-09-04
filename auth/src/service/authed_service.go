package service

import (
	"auth/model"
	"auth/utils"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenJWT(userid string) (string, *utils.HttpError) {
	// JWT の有効期限取得
	ExpirySec := os.Getenv("EXPIRYSEC")

	// 数字に変換
	ExpirySecInt, err := strconv.Atoi(ExpirySec)

	// エラー処理
	if err != nil {
		return "", utils.NewHttpError(http.StatusInternalServerError, "failed to convert EXPIRYSEC")
	}

	// JWT作成
	token := jwt.NewWithClaims(SignMethod, jwt.MapClaims{
		"userid": userid,
		"exp":    utils.NowTime().Add(time.Second * time.Duration(ExpirySecInt)).Unix(),
	})

	// jwt 署名
	tokenString, err := token.SignedString([]byte(JWT_KEY))

	// エラー処理
	if err != nil {
		return "", utils.NewHttpError(http.StatusInternalServerError, "failed to sign token")
	}

	return tokenString, nil
}

// JWT を検証してユーザーを返す
func ValidateJwt(tokenString string) (model.User, error) {
	// トークンを検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWT_KEY), nil
	})

	// エラー処理
	if err != nil {
		return model.User{}, err
	}

	// 検証できたか
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//検証できた場合
		userid := claims["userid"]

		// ユーザーを取得
		return model.GetUserByID(userid.(string))
	} else {
		return model.User{}, err
	}
}

func GetUserInfo(userid string) (model.User,error) {
	// ユーザー情報取得
	userData,err := model.GetUserByID(userid)

	// エラー処理
	if err != nil {
		return model.User{},err
	}

	return userData,nil
}