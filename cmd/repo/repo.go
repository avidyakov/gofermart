package repo

import "errors"

type Repo interface {
	CreateUser(login, password string) error
}

var ErrUserExists = errors.New("user already exists")
