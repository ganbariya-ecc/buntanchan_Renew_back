package main

import (
	"template/router"

)

func main() {
	// ルーター初期化
	router := router.InitRouter()

	// Start server
	router.Logger.Fatal(router.Start("0.0.0.0:3000"))
}
