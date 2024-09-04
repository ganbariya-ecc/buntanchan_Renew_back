package main

import (
	"group/model"
	"group/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Env 読み込み
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	LoadEnv()

	// モデル初期化
	model.Init(os.Getenv("DBPATH"))

	// ルーター初期化
	router := router.InitRouter()

	// Start server
	router.Logger.Fatal(router.Start("0.0.0.0:3000"))
}
