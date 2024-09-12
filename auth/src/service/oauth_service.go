package service

import (
	"auth/model"
	"auth/utils"
	"errors"
	"log"
	"os"
	"path"

	"github.com/markbates/goth"
)

// Oauth の完了
func CallbackOauth(user goth.User) (string,error) {
	userName := ""

	log.Println(user)

	// ユーザー名が空かどうか
	if user.Name != "" {
		// ユーザー名
		userName = user.Name
	} else if user.FirstName != "" {
		// ファーストネーム設定
		userName := user.FirstName

		// ラスト名設定
		if user.LastName != "" {
			//ラスト名設定
			userName += user.LastName
		}
	} else {
		//ニックネームに設定
		userName = user.NickName
	}

	// ユーザー作成
	userid,err := model.CreateOauthUser(userName,user.UserID,[]model.UserLabel{},user.Email,user.Provider)

	// エラー処理
	if err != nil {
		return "",err
	}

	// 一時ファイルのパス
	tempPath := path.Join(model.UserIconDir,userid + ".tmp")

	// 一時ファイルのパス
	iconPath := path.Join(model.UserIconDir,userid + ".jpeg")

	// アイコンをダウンロード
	err = utils.DownloadFile(tempPath,user.AvatarURL)

	// 成功したとき
	if err == nil {

		// 画像をロード
		imgData,err := utils.LoadImage(tempPath)

		// エラー処理
		if err != nil {
			return "",errors.New("Failed to load icon")
		}

		// 画像をリサイズ
		resizedImg := utils.ResizeImage(imgData.Src,256,256)

		// 画像を保存する
		err = utils.SaveImage(iconPath,resizedImg)

		// エラー処理
		if err != nil {
			return "",errors.New("Failed to save icon")
		}

		// 一時ファイルを削除
		err = os.Remove(tempPath)

		if err != nil {
			log.Println("failed to delete temp file")
		}
	}

	return userid,nil
}