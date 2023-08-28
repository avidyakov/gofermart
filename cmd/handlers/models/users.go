package models

import (
	"encoding/json"
	"errors"
	"io"
)

type Model interface {
	Validate() error
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewUser(rawUser io.Reader) (*User, error) {
	var user User
	decoder := json.NewDecoder(rawUser)

	decodeErr := decoder.Decode(&user)
	if decodeErr != nil {
		return nil, decodeErr
	}

	validationErr := user.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	return &user, nil
}

func (u *User) Validate() error {
	if u.Login == "" {
		return errors.New("login is required")
	} else if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
