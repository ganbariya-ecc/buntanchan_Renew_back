package sdk_client

import (
	"log"

	"task/sdks/sdk_client/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	// GRPC 接続
	gconn *grpc.ClientConn = nil

	// GRPC クライアント
	gclient protoc.TemplateServiceClient = nil

	// 初期化済みか
	isInit = false

	// SDK を使うためのキー
	sdkkey = ""
)

func Init(ServerAddr, SDKKEY string) {
	addr := ServerAddr
	// TLS認証を追加
	creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	if err != nil {
		log.Fatal(err)
	}

	// 接続作成
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	// conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// クライアント作成
	aclient := protoc.NewTemplateServiceClient(conn)

	// グローバル変数に格納
	gclient = aclient
	gconn = conn
	sdkkey = SDKKEY

	isInit = true
}

func Disconnect() error {
	// 初期化済みでない場合 panic
	if !isInit {
		log.Fatalln("Not initialized")
	}

	// 接続を切る
	return gconn.Close()
}
