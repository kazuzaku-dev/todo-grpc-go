//go:build wireinject
// +build wireinject

package repository_impl

import (
	"github.com/google/wire"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/domain/repository"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/infrastructure/database"
)

var UserRepositorySet = wire.NewSet(
	database.NewDB,
	NewUserRepositoryImpl,
	wire.Bind(new(repository.UserRepository), new(*UserRepositoryImpl)),
)
