package groupsdk

import (
	"context"
	"errors"
	"task/sdks/groupsdk/protoc"
	"log"
)

func GetMember(memberid string) (protoc.MemberData, error) {
	// 初期化済みでない場合 panic
	if !isInit {
		log.Fatalln("Not initialized")
	}

	// コンテキスト生成
	ctx := context.Background()

	// メンバーを取得する
	result, err := gclient.GetMember(ctx, &protoc.GetMemberRequest{
		Memberid: memberid,
	})

	// エラー処理
	if err != nil {
		return protoc.MemberData{}, err
	}

	// 成功したか
	if result.Success {
		// 成功した場合
		return *result.Data, nil
	}

	return protoc.MemberData{}, errors.New("failed to get member")
}
