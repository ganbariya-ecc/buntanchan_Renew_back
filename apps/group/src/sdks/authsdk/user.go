package authsdk

import (
	"context"
	"errors"
	"log"
	"template/sdks/authsdk/protoc"
)

func CreateUser(UserName,Password string) (string, error) {
	// 初期化済みでない場合 panic
	if !isInit {
		log.Fatalln("Not initialized")
	}

	// コンテキスト生成
	ctx := context.Background()

	// トークンを渡してユーザーを取得する
	result, err := gaclient.Create(ctx, &protoc.CreateData{
		SDKKEY: sdkkey,
		UserName: UserName,
		Password: Password,
	})

	// エラー処理
	if err != nil {
		return "", err
	}

	// 成功したか
	if result.Success {
		// 成功した場合
		return *&result.Userid, nil
	}

	return "", errors.New("User creation failed")
}
