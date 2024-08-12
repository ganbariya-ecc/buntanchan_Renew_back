package main

import (
	"auth/controller"
	"auth/model"
	"auth/router"
	"auth/utils"
	"os"
)

func main() {
	// //初期化
	// Init()

	// テスト
	Test()

	//サーバー起動
	// ServerMain()
}

func ServerMain() {
	// ルーター初期化
	router := router.InitRouter()

	// Start server
	router.Logger.Fatal(router.Start("0.0.0.0:3001"))
}

//初期化
func Init() {
	// env ファイル読み込む
	utils.LoadEnv()

	//コントローラー初期化
	controller.Init()

	//モデル初期化
	model.Init()
}

func Test() {
	// env ファイル読み込む
	utils.LoadEnv()

	//コントローラー初期化
	controller.Init()
	
	// ファイル削除
	os.Remove("./auth.db")

	//モデル初期化
	model.Init()

	//データーベーステスト
	model.Test()
}