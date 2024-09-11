package main

import (
	"log"
	"os"
	"task/model"
	"task/router"

	"task/sdks/authsdk"
	"task/sdks/groupsdk"

	"github.com/joho/godotenv"

	"task/sdks/sdk_server"
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

	// グループSDK 初期化
	groupsdk.Init(os.Getenv("GROUP_ADDR"))

	// モデル初期化
	model.Init(os.Getenv("DBPATH"))

	// サーバー開始
	go sdk_server.StartServer(os.Getenv("GRPC_ADDR"))

	// log.Println("サーバー起動")

	// ルーター初期化
	router := router.InitRouter()

	// Start server
	router.Logger.Fatal(router.Start("0.0.0.0:3000"))
}
