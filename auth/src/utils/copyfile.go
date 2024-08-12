package utils

import (
	"io"
	"os"
)

func CopyFile(src,dst string) error {
	// ファイルを作成
	dst_file,err := os.Create(dst)

	// エラー処理
	if err != nil {
		return err
	}

	// ファイルを開く
	src_file,err := os.Open(src)

	// エラー処理
	if err != nil {
		return err
	}

	// ファイルをコピー
	io.Copy(dst_file,src_file)

	return nil
}