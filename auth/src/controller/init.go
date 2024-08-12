package controller

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/discord"
)

func Init() {
	// Oauth2 プロバイダー初期化
	goth.UseProviders(
		discord.New(os.Getenv("Discord_ClientID"),os.Getenv("Discord_Secret"),os.Getenv("Discord_Callback")),
	)
}