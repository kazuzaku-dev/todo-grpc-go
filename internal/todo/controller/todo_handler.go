package controller

import (
	"context"

	pb "github.com/kazuzaku-dev/todo-grpc-go/grpc"
)

type ToDoServiceHandler struct {
	pb.UnimplementedToDoServieServer
}

var _ pb.ToDoServieServer = (*ToDoServiceHandler)(nil)

func NewTodoServiceHandler() *ToDoServiceHandler {
	return &ToDoServiceHandler{}
}

func (t *ToDoServiceHandler) Create(context.Context, *pb.ToDo) (*pb.CreateToDoResponse, error) {
	panic("unimplemented")
}

func (t *ToDoServiceHandler) Update(context.Context, *pb.ToDo) (*pb.UpdateToDoResponse, error) {
	panic("unimplemented")
}

func (t *ToDoServiceHandler) Delete(context.Context, *pb.ToDo) (*pb.DeleteToDoResponse, error) {
	panic("unimplemented")
}

func (t *ToDoServiceHandler) Get(context.Context, *pb.GetToDoRequest) (*pb.ToDo, error) {
	panic("unimplemented")
}

func (t *ToDoServiceHandler) List(context.Context, *pb.GetToDoListRequest) (*pb.GetToDoListResponse, error) {
	panic("unimplemented")
}
