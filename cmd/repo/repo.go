package repo

import "errors"

type Repo interface {
	CreateUser(login, password string) (userID uint, err error)
	GetUser(login string) (userID uint, err error)
	CheckPassword(login, password string) (bool, error)

	CreateOrder(number string, userID uint) (orderID uint, err error)
}

var ErrUserExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user not found")
