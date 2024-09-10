package sdk_server

import (
	"context"
	"group/model"
	"group/sdks/sdk_server/protoc"

	// "auth/service"
	"log"
	"net"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

func StartServer(bindAddr string) error {
	log.Print("main start")

	// 9000番ポートでクライアントからのリクエストを受け付けるようにする
	listen, err := net.Listen("tcp", bindAddr)
	if err != nil {
		return err
	}

	// 鍵読み込み
	cred, err := credentials.NewServerTLSFromFile("./sdks/sdk_server/server.crt", "./sdks/sdk_server/server.key")
	if err != nil {
		log.Fatal(err)
	}

	// GRPC サーバー起動
	grpcServer := grpc.NewServer(grpc.Creds(cred))

	// 構造体を登録する
	protoc.RegisterGroupSdkServiceServer(grpcServer, &GroupSdkService{})

	// 以下でリッスンし続ける
	if err := grpcServer.Serve(listen); err != nil {
		return err
	}

	return nil
}

type GroupSdkService struct{}

// GetMember implements protoc.GroupSdkServiceServer.
func (groupS *GroupSdkService) GetMember(ctx context.Context, req *protoc.GetMemberRequest) (*protoc.GetMemberResponse, error) {
	// メンバーを取得
	member,err := model.GetMember(req.Memberid)

	// エラー処理
	if err != nil {
		log.Println("failed to get member from sdk : " + err.Error())
		return &protoc.GetMemberResponse{
			Success: false,
		},err
	}

	return &protoc.GetMemberResponse{
		Success: true,
		Data: &protoc.MemberData{
			Memberid: member.MemberID,
			MemberName: member.Name,
			GroupID: member.GroupID,
			MemberRole: string(member.MemberRole),
		},
	},nil
}
