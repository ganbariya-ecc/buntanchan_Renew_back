package controller

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/line"
	"github.com/markbates/goth/providers/google"

	"github.com/gorilla/sessions"
)

func Init() {
	key := "lfaKhdetDvdiUeJHDP6CBujepYrkyQL8gUKOlATKvdvKsEYYJ2xm3iWia9j8qKkb"            // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30 // 30 days
	isProd := true      // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	// Oauth2 プロバイダー初期化
	goth.UseProviders(
		discord.New(os.Getenv("Discord_ClientID"), os.Getenv("Discord_Secret"), os.Getenv("Discord_Callback")),
		google.New(os.Getenv("Google_KEY"), os.Getenv("Google_SECRET"), os.Getenv("Google_Callback")),
		line.New(os.Getenv("Line_KEY"), os.Getenv("Line_SECRET"), os.Getenv("Line_Callback"), "profile"),
	)
}
