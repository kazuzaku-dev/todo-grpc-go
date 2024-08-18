package command

import (
	"context"

	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/domain/repository"
)

type UpdateUserCommand struct {
	UserID   string
	Name     string
	Email    string
	Password string
}

type UpdateUserCommandHandler struct {
	repository repository.UserRepository
}

func NewUpdateUserCommandHandler(
	repository repository.UserRepository,
) *UpdateUserCommandHandler {
	return &UpdateUserCommandHandler{
		repository: repository,
	}
}

func (handler *UpdateUserCommandHandler) Handle(ctx context.Context, command *UpdateUserCommand) (string, error) {
	user, err := handler.repository.FindByID(ctx, command.UserID)
	if err != nil {
		return "", err
	}

	newUser := user.Update(
		command.Name,
		command.Email,
		command.Password,
	)

	userID, error := handler.repository.Update(ctx, newUser)
	if error != nil {
		return "", error
	}

	return userID, nil
}
