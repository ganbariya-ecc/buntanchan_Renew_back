package sdk_client

import (
	"context"
	"errors"
	"log"
	"template/sdks/sdk_client/protoc"
)

func Auth() (protoc.TemplateResult, error) {
	// 初期化済みでない場合 panic
	if !isInit {
		log.Fatalln("Not initialized")
	}

	// コンテキスト生成
	ctx := context.Background()

	// テストを実行する
	result, err := gclient.Test(ctx, &protoc.TemplateData{
		Msg: "hello world",
	})

	// エラー処理
	if err != nil {
		return protoc.TemplateResult{}, err
	}

	// 成功したか
	if result.Success {
		// 成功した場合
		return *result, nil
	}

	return protoc.TemplateResult{}, errors.New("Template failure")
}


