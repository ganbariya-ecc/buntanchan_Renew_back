package authsdk

import (
	"context"
	"errors"
	"log"
	"template/sdks/authsdk/protoc"
)

func Auth(token string) (protoc.User, error) {
	// 初期化済みでない場合 panic
	if !isInit {
		log.Fatalln("Not initialized")
	}

	// コンテキスト生成
	ctx := context.Background()

	// トークンを渡してユーザーを取得する
	result, err := gaclient.Auth(ctx, &protoc.AuthData{
		Token: token,
	})

	// エラー処理
	if err != nil {
		return protoc.User{}, err
	}

	// 成功したか
	if result.Success {
		// 成功した場合
		return *result.User, nil
	}

	return protoc.User{}, errors.New("User authentication failed")
}


