package entity

import (
	"github.com/kazuzaku-dev/todo-grpc-go/pkg/util"
)

type User struct {
	userID   string
	name     string
	email    string
	password string
}

// ユーザーエンティティを生成する
func NewUser(name, email, password string) *User {
	// validationなどのビジネスロジック、ルールをここに書く
	// 今回は省略

	// userIDをULIDで生成する
	userID := util.NewULID()

	return &User{
		userID:   userID,
		name:     name,
		email:    email,
		password: password,
	}
}

func (u *User) GetUserID() string {
	return u.userID
}

func (u *User) copy() *User {
	return &User{
		userID:   u.userID,
		name:     u.name,
		email:    u.email,
		password: u.password,
	}
}

func (u *User) Update(name, email, password string) *User {
	// validationなどのビジネスロジック、ルールをここに書く
	// 今回は省略

	// userIDは変更しない
	updatedUser := u.copy()
	updatedUser.name = name
	updatedUser.email = email
	updatedUser.password = password

	return updatedUser
}
