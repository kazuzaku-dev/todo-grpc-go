//go:build wireinject
// +build wireinject

package command

import (
	"github.com/google/wire"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/infrastructure/repository_impl"
)

var UserCommandSet = wire.NewSet(
	repository_impl.UserRepositorySet,
	NewCreateUserCommandHandler,
)
