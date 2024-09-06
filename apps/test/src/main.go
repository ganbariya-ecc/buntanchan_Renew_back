package main

import (
	"log"
	"os"
	"template/model"
	"template/router"

	"template/sdks/authsdk"

	"github.com/joho/godotenv"

	"template/sdks/sdk_template"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//Env 読み込み
	LoadEnv()

	//認証sdk 初期化
	authsdk.Init(os.Getenv("AUTH_ADDR"), os.Getenv("AUTH_SDKKEY"))

	// モデル初期化
	model.Init(os.Getenv("DBPATH"))

	// サーバー開始
	sdk_template.StartServer(os.Getenv("GRPC_ADDR"))

	// ルーター初期化
	router := router.InitRouter()

	// Start server
	router.Logger.Fatal(router.Start("0.0.0.0:3000"))
}
