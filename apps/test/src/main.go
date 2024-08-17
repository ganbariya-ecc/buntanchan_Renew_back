package main

import (
	"test/router"

	"test/sdk"
)

func main() {
	// ルーター初期化
	router := router.InitRouter()

	//sdk 初期化
	sdk.Init()

	// Start server
	router.Logger.Fatal(router.Start("0.0.0.0:3000"))
}
