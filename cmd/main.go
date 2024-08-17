package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
	// not implemented yet.
}

// エラーハンドリング用のInterceptor
func errorHandlingUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
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
