package utils

import (
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"
)

type Img struct {
	Src    image.Image // 画像データ
	Path   string      // 画像のパス
	Width  int         // 横幅
	height int         //立幅
}

func LoadImage(path string) (Img, error) {
	// ファイルを開く
	file, err := os.Open(path)

	// エラー処理a-
	if err != nil {
		return Img{},err
	}

	// ファイルを自動で閉じるようにする
	defer file.Close()

	// 画像をデコードする
	src, _, err := image.Decode(file)

	// エラー処理
	if err != nil {
		return Img{},err
	}

	// 画像のサイズ取得
	size := src.Bounds().Size()
	width, height := size.X, size.Y

	// 返す構造体生成
	imgData := Img{
		Src: src,
		Path: path,
		Width: width,
		height: height,
	}

	return imgData,nil
}
