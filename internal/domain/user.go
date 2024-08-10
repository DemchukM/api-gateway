package domain

import "errors"

type User struct {
	ID    int
	login string
}

type UserRepository interface {
	GetById(id int) (*User, error)
	Create(user *User) (int, error)
}

var ErrUserNotFound = errors.New("user not found")
