package entity

import (
	"github.com/kazuzaku-dev/todo-grpc-go/pkg/util"
)

type User struct {
	UserID   string
	Name     string
	Email    string
	Password string
}

// ユーザーエンティティを生成する
func NewUser(name, email, password string) *User {
	// validationなどのビジネスロジック、ルールをここに書く
	// 今回は省略

	// userIDをULIDで生成する
	userID := util.NewULID()

	return &User{
		UserID:   userID,
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (u *User) Update(name, email, password string) *User {
	// validationなどのビジネスロジック、ルールをここに書く
	// 今回は省略

	// userIDは変更しない
	updatedUser := &User{
		UserID:   u.UserID,
		Name:     name,
		Email:    email,
		Password: password,
	}

	return updatedUser
}
