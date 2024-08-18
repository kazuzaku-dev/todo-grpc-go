//go:build wireinject
// +build wireinject

package todo

import (
	"github.com/google/wire"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/controller"
)

func InitUserService() (*controller.UserServiceHandler, error) {
	wire.Build(controller.UserServiceHandlerSet)
	return &controller.UserServiceHandler{}, nil
}

func InitTodoService() (*controller.ToDoServiceHandler, error) {
	wire.Build(controller.ToDoServieHandlerSet)
	return &controller.ToDoServiceHandler{}, nil
}
