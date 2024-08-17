package sdk

import (
	"context"
	"log"
	"test/sdk/protoc"
)

func Auth(token string) {
	// 初期化済みでない場合 panic
	if !isInit {
		log.Fatalln("Not initialized")
	}

	ctx := context.Background()

	result, err := gaclient.Auth(ctx, &protoc.AuthData{
		Token: token,
	})
	if err != nil {
		log.Println("could not greet: %v", err)
	}

	log.Println(result)
}