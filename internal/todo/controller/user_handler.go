package controller

import (
	"context"

	pb "github.com/kazuzaku-dev/todo-grpc-go/grpc"
)

type UserServiceHandler struct {
	pb.UnimplementedUserServiceServer
}

var _ pb.UserServiceServer = (*UserServiceHandler)(nil)

func NewUserServiceHandler() *UserServiceHandler {
	return &UserServiceHandler{}
}

func (handler *UserServiceHandler) Create(ctx context.Context, req *pb.User) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}

func (handler *UserServiceHandler) Update(context.Context, *pb.User) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{}, nil
}

func (handler *UserServiceHandler) Delete(context.Context, *pb.User) (*pb.DeleteUserResponse, error) {
	return &pb.DeleteUserResponse{}, nil
}

func (handler *UserServiceHandler) Get(context.Context, *pb.GetUserRequest) (*pb.User, error) {
	return &pb.User{}, nil
}

func (handler *UserServiceHandler) List(context.Context, *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	return &pb.GetUserListResponse{}, nil
}
