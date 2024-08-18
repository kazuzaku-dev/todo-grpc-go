package command

import (
	"context"

	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/domain/entity"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/domain/repository"
)

type CreateUserCommand struct {
	Name     string
	Email    string
	Password string
}

type CreateUserCommandHandler struct {
	repository repository.UserRepository
}

func NewCreateUserCommandHandler(
	repository repository.UserRepository,
) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		repository: repository,
	}
}

func (handler *CreateUserCommandHandler) Handle(ctx context.Context, command *CreateUserCommand) (string, error) {
	// ユーザーエンティティを生成する
	user := entity.NewUser(command.Name, command.Email, command.Password)

	// ユーザーを作成する（永続化）
	userID, error := handler.repository.Create(ctx, user)
	if error != nil {
		return "", error
	}

	return userID, nil
}
