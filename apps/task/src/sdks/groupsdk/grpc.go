package groupsdk

import (
	"log"

	"task/sdks/groupsdk/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	// GRPC 接続
	gconn *grpc.ClientConn = nil

	// GRPC クライアント
	gclient protoc.GroupSdkServiceClient = nil

	// 初期化済みか
	isInit = false
)

func Init(ServerAddr string) {
	addr := ServerAddr
	// TLS認証を追加
	creds, err := credentials.NewClientTLSFromFile("./sdks/groupsdk/server.crt", "")
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
	aclient := protoc.NewGroupSdkServiceClient(conn)

	// グローバル変数に格納
	gclient = aclient
	gconn = conn

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
