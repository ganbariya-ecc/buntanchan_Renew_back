package main

import (
	"auth/controller"
	"auth/model"
	"auth/router"
	"auth/sdks"
	"auth/utils"
	"log"
	"os"
)

func main() {
	//初期化
	Init()

	// テスト
	// Test()

	// 実行開始ログ
	log.Println("Server execution starts")

	// サーバー起動
	ServerMain()
}

func ServerMain() {
	// マルチスレッドで開始
	go func ()  {
		// GRPC サーバー起動
		err := sdks.StartServer(os.Getenv("GRPC_BindAddr"))

		// エラー処理
		if err != nil {
			log.Fatalln("faild to start sdk server : " + err.Error())
		}
	}()

	// ルーター初期化
	router := router.InitRouter()

	// Start server
	router.Logger.Fatal(router.Start(os.Getenv("BindAddr")))
}

// 初期化 
func Init() {
	// env ファイル読み込む
	utils.LoadEnv()

	//コントローラー初期化
	controller.Init()

	//モデル初期化
	model.Init(os.Getenv("DBPATH"))
}

func Test() {
	// env ファイル読み込む
	utils.LoadEnv()

	//コントローラー初期化
	controller.Init()

	// ファイル削除
	os.Remove("./test.db")

	//モデル初期化
	model.Init("./test.db")

	//データーベーステスト
	model.Test()
}
