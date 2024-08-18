package controller

import (
	"context"
	"log"

	pb "github.com/kazuzaku-dev/todo-grpc-go/grpc"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/usecase/command"
)

type UserServiceHandler struct {
	pb.UnimplementedUserServiceServer
	createCommand *command.CreateUserCommandHandler
}

var _ pb.UserServiceServer = (*UserServiceHandler)(nil)

func NewUserServiceHandler(
	createCommand *command.CreateUserCommandHandler,
) *UserServiceHandler {
	return &UserServiceHandler{
		createCommand: createCommand,
	}
}

func (handler *UserServiceHandler) Create(ctx context.Context, req *pb.User) (*pb.CreateUserResponse, error) {
	id, error := handler.createCommand.Handle(ctx, &command.CreateUserCommand{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if error != nil {
		log.Fatalf("failed to create user: %v", error)
		return nil, error
	}
	return &pb.CreateUserResponse{
		Id: id,
	}, nil
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
