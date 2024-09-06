package sdks

import (
	"auth/sdks/protoc"
	"auth/service"
	"context"
	"errors"
	"log"
	"net"
	"os"

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

	// AuthService構造体のアドレスを渡すことで、クライアントからAuthリクエストされると
	// Authメソッドが呼ばれるようになる
	protoc.RegisterAuthServiceServer(grpcServer, &AuthService{})

	// 以下でリッスンし続ける
	if err := grpcServer.Serve(listen); err != nil {
		return err
	}

	return nil
}

type AuthService struct{}

// Create implements protoc.AuthServiceServer.
func (aservice *AuthService) Create(ctx context.Context, data *protoc.CreateData) (*protoc.CreateResponse, error) {
	// SDKKEY 認証
	if data.SDKKEY != os.Getenv("SDK_KEY") {
		return &protoc.CreateResponse{
			Success: false,
			Userid: "",
		},errors.New("Username not specified")
	}

	// バリデーション
	if data.UserName == "" {
		return &protoc.CreateResponse{
			Success: false,
			Userid: "",
		},errors.New("Username not specified")
	}

	// パスワード検証
	if data.Password == "" {
		return &protoc.CreateResponse{
			Success: false,
			Userid: "",
		},errors.New("Password not specified")
	}

	// ユーザーを作成する
	user_data,err := service.BasicSignup(data.UserName,data.Password)

	// エラー処理
	if err != nil {
		log.Println("Failed to create user from SDK : " + err.Error())
		return &protoc.CreateResponse{
			Success: false,
			Userid: "",
		},err
	}

	return &protoc.CreateResponse{
		Success: true,
		Userid: user_data.UserID,
	},nil
}

func (aservice *AuthService) Auth(
	ctx context.Context,
	data *protoc.AuthData,
) (*protoc.UserResult, error) {
	// トークンを検証してユーザーを取得する
	user, err := service.ValidateJwt(data.Token)

	// エラー処理
	if err != nil {
		log.Println("GRPC authentication error : " + err.Error())
		return &protoc.UserResult{
			Success: false,
		}, err
	}

	// ユーザー情報生成
	userData := protoc.User{
		UserID:    user.UserID,
		UserName:  user.UserName,
		Email:     user.Email,
		AuthType:  string(user.AuthType),
		CreatedAt: user.CreatedAt.Unix(),
	}

	return &protoc.UserResult{
		Success: true,
		User:    &userData,
	}, nil
}
