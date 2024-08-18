//go:build wireinject
// +build wireinject

package controller

import (
	"github.com/google/wire"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/usecase/command"
)

var UserServiceHandlerSet = wire.NewSet(
	command.UserCommandSet,
	NewUserServiceHandler,
)

var ToDoServieHandlerSet = wire.NewSet(
	NewTodoServiceHandler,
)
