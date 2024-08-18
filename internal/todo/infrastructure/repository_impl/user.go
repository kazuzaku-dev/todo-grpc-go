package repository_impl

import (
	"context"

	"gorm.io/gorm"

	"github.com/kazuzaku-dev/todo-grpc-go/internal/generated/gorm/dao"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/generated/gorm/model"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/domain/entity"
	"github.com/kazuzaku-dev/todo-grpc-go/internal/todo/domain/repository"
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
	gUser := &model.GUser{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
	}

	err := r.db.Transaction(func(tx *gorm.DB) (string error) {
		userDao := dao.Use(r.db).GUser
		error := userDao.Create(gUser)
		if error != nil {
			return error
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return gUser.UserID, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) (string, error) {
	gUser := &model.GUser{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
	}

	userDao := dao.Use(r.db).GUser
	error := userDao.Save(gUser)
	if error != nil {
		return "", error
	}
	return gUser.UserID, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, userID string) (string, error) {
	userDao := dao.Use(r.db).GUser
	_, error := userDao.Delete(&model.GUser{
		UserID: userID,
	})
	if error != nil {
		return "", error
	}
	return "Success", nil
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, userID string) (*entity.User, error) {
	userDao := dao.Use(r.db).GUser
	gUser, error := userDao.Where(userDao.UserID.Eq(userID)).First()
	if error != nil {
		return nil, error
	}

	user := &entity.User{
		UserID: gUser.UserID,
		Name:   gUser.Name,
		Email:  gUser.Email,
	}
	return user, nil
}
