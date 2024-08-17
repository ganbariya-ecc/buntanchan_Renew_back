package sdks

import (
	"auth/sdks/protoc"
	"auth/service"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func StartServer(bindAddr string) error {
	log.Print("main start")

	// 9000番ポートでクライアントからのリクエストを受け付けるようにする
	listen, err := net.Listen("tcp",bindAddr)
	if err != nil {
		return err
	}

	// GRPC サーバー起動
	grpcServer := grpc.NewServer()

	// AuthService構造体のアドレスを渡すことで、クライアントからAuthリクエストされると
	// Authメソッドが呼ばれるようになる
	protoc.RegisterAuthServiceServer(grpcServer, &AuthService{})

	// 以下でリッスンし続ける
	if err := grpcServer.Serve(listen); err != nil {
		return err
	}

	return nil
}

type AuthService struct {}

func (aservice *AuthService) Auth(
	ctx context.Context,
	data *protoc.AuthData,
) (*protoc.UserResult, error) {
	// トークンを検証してユーザーを取得する
	user,err := service.ValidateJwt(data.Token)

	// エラー処理
	if err != nil {
		log.Println("GRPC authentication error : " + err.Error())
		return &protoc.UserResult{
			Success: false,
		},err
	}

	// ユーザー情報生成
	userData := protoc.User{
		UserID: user.UserID,
		UserName: user.UserName,
		Email: user.Email,
		AuthType: string(user.AuthType),
		CreatedAt: user.CreatedAt.Unix(),
	}

	return &protoc.UserResult{
		Success: true,
		User: &userData,
	},nil
}

