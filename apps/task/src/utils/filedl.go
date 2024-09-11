package utils

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string) error {
	// URL からデータを取得
    resp, err := http.Get(url)

	// エラー処理
    if err != nil {
        return err
    }

	// 自動で閉じるようにする
    defer resp.Body.Close()

	// ファイルを作成する
    outfile, err := os.Create(filepath)

	// エラー処理
    if err != nil {
        return err
    }

	// 自動で閉じるようにする
    defer outfile.Close()

	// バッファをコピー
    _, err = io.Copy(outfile, resp.Body)
    return err
}