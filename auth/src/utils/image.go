package utils

import (
	"image"
	"os"

	_ "image/gif"
	"image/jpeg"
	_ "image/png"

	"golang.org/x/image/draw"
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

func ResizeImage(img image.Image, width, height int) image.Image {
	// 欲しいサイズの画像を新しく作る
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// サイズを変更しながら画像をコピーする
	draw.BiLinear.Scale(newImage, newImage.Bounds(), img, img.Bounds(), draw.Over, nil)

	return newImage
}

// 画像を保存する関数。
// 保存先のパスと画像データを渡すと保存してくれる。
func SaveImage(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, img, &jpeg.Options{
		Quality: 50, // JPEGのクオリティ設定。省略するとjpeg.DefaultQualityの値（75）が使われる。
	})
	return err
}
