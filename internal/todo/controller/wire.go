//go:build wireinject
// +build wireinject

package controller

import (
	"github.com/google/wire"
)

var UserServiceHandlerSet = wire.NewSet(
	NewUserServiceHandler,
)

var ToDoServieHandlerSet = wire.NewSet(
	NewTodoServiceHandler,
)
