package repository_impl

import (
	"context"

	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/domain/entity"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/domain/repository"
	"gorm.io/gorm"
)

// UserRepositoryImpl ...
type UserRepositoryImpl struct {
	db *gorm.DB
}

var _ repository.UserRepository = (*UserRepositoryImpl)(nil)

func NewUserRepositoryImpl(
	db *gorm.DB,
) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) (string, error) {
	r.db.
	return "", nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) (string, error) {
	return "", nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, userID string) (string, error) {
	return "", nil
}

func (r *UserRepositoryImpl) Get(ctx context.Context, userID string) (*entity.User, error) {
	return nil, nil
}
