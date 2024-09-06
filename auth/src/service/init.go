package service

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var (
	// JWT 署名メソッド
	SignMethod = jwt.SigningMethodHS512

	// JWT 署名用のキー
	JWT_KEY = os.Getenv("JWT_KEY")
)
