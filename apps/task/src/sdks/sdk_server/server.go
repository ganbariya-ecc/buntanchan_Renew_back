package sdk_server

import (
	"context"
	"task/sdks/sdk_server/protoc"

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
	cred, err := credentials.NewServerTLSFromFile("./server.crt", "./server.key")
	if err != nil {
		log.Fatal(err)
	}

	// GRPC サーバー起動
	grpcServer := grpc.NewServer(grpc.Creds(cred))

	// 構造体を登録する
	protoc.RegisterTemplateServiceServer(grpcServer, &TemplateService{})

	// 以下でリッスンし続ける
	if err := grpcServer.Serve(listen); err != nil {
		return err
	}

	return nil
}

type TemplateService struct{}

// Test implements protoc.TemplateServiceServer.
func (tservice *TemplateService) Test(context.Context, *protoc.TemplateData) (*protoc.TemplateResult, error) {
	log.Println("test")

	return &protoc.TemplateResult{
		Success: true,
	}, nil
}
