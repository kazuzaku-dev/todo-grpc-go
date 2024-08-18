package repository

import (
	"context"

	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (string, error)
	Update(ctx context.Context, user *entity.User) (string, error)
	Delete(ctx context.Context, userID string) (string, error)
	FindByID(ctx context.Context, userID string) (*entity.User, error)
}
