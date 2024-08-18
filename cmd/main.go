package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/kazuzaku-dev/todo-grpc-go/grpc"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(errorHandlingUnaryInterceptor),
		grpc.StreamInterceptor(errorHandlingStreamInterceptor),
	}

	// gRPCサーバーを作成する
	server := grpc.NewServer(opts...)

	// サーバーにリフレクションを登録する
	reflection.Register(server)

	// サーバーにサービスを登録する
	registerService(server)

	// サーバーを起動する
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// サーバーにgRPCサービスを登録する
func registerService(server *grpc.Server) {
	userServie, err := todo.InitUserService()
	if err != nil {
		log.Fatalf("failed to create user service: %v", err)
	}
	pb.RegisterUserServiceServer(server, userServie)

	todoServie, err := todo.InitTodoService()
	if err != nil {
		log.Fatalf("failed to create todo service: %v", err)
	}
	pb.RegisterToDoServieServer(server, todoServie)
}

// エラーハンドリング用のInterceptor
func errorHandlingUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		// panicが発生した場合の処理
		if r := recover(); r != nil {
			// スタックトレースを取得してログに出力
			buf := make([]byte, 1<<16)
			stackSize := runtime.Stack(buf, false)
			fmt.Printf("panic recovered: %s\nStack trace:\n%s\n", r, buf[:stackSize])

			// クライアントにエラーレスポンスを返す
			err = status.Errorf(codes.Internal, "Internal server error")
		}
	}()
	resp, err = handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error: %v", err)
		return nil, status.Errorf(status.Code(err), "RPC error: %v", err)
	}
	return resp, nil
}

// エラーハンドリング用のInterceptor（Stream用）
func errorHandlingStreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	err := handler(srv, ss)
	if err != nil {
		log.Printf("RPC failed with error: %v", err)
		return status.Errorf(status.Code(err), "RPC error: %v", err)
	}
	return nil
}
