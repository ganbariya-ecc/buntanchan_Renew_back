package sdks

import (
	"auth/model"
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

// GetUserAll implements protoc.AuthServiceServer.
func (aservice *AuthService) GetUserAll(ctx context.Context,data *protoc.GetUserAllRequest) (*protoc.GetUserAllResponse, error) {
	// SDKKEY 認証
	if data.SDKKEY != os.Getenv("SDK_KEY") {
		return &protoc.GetUserAllResponse{
			Success: false,
			User: &protoc.UserAllData{},
		}, errors.New("SDKKEY does not match")
	}

	// ユーザーIDがないとき
	if data.UserID == "" {
		return &protoc.GetUserAllResponse{
			Success: false,
			User: &protoc.UserAllData{},
		}, errors.New("UserID not specified")
	}

	// ユーザーを取得
	user,err := model.GetUserByID(data.UserID)

	// エラー処理
	if err != nil {
		// ユーザーの取得に失敗した場合
		log.Println("Failed to get user from SDK : " + err.Error())
		
		return &protoc.GetUserAllResponse{
			Success: false,
			User: &protoc.UserAllData{},
		}, errors.New("Failed to get user")
	}

	return &protoc.GetUserAllResponse{
		Success: true,
		User: &protoc.UserAllData{
			UserID: user.UserID,
			UserName: user.UserName,
			AuthType: string(user.AuthType),
			Email: user.Email,
			CreatedAt: user.CreatedAt.Unix(),
			Password: user.Password,
			IsHashed: user.HashPassword,
		},
	}, nil
}

// Create implements protoc.AuthServiceServer.
func (aservice *AuthService) Create(ctx context.Context, data *protoc.CreateData) (*protoc.CreateResponse, error) {
	// SDKKEY 認証
	if data.SDKKEY != os.Getenv("SDK_KEY") {
		return &protoc.CreateResponse{
			Success: false,
			Userid:  "",
		}, errors.New("SDKKEY does not match")
	}

	// バリデーション
	if data.UserName == "" {
		return &protoc.CreateResponse{
			Success: false,
			Userid:  "",
		}, errors.New("Username not specified")
	}

	// パスワード検証
	if data.Password == "" {
		return &protoc.CreateResponse{
			Success: false,
			Userid:  "",
		}, errors.New("Password not specified")
	}

	// ユーザーを作成する
	userid, err := model.CreateUser(data.UserName, []model.UserLabel{}, data.Password, false)

	// エラー処理
	if err != nil {
		log.Println("Failed to create user from SDK : " + err.Error())
		return &protoc.CreateResponse{
			Success: false,
			Userid:  "",
		}, err
	}

	return &protoc.CreateResponse{
		Success: true,
		Userid:  userid,
	}, nil
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
