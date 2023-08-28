package postgres

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"gophermart/cmd/repo"
	"gorm.io/gorm"
)

const (
	PgUniqueViolationErrorCode = "23505"
)

func (r *Repo) CreateUser(login, password string) (uint, error) {
	user := User{
		Username: login,
	}
	user.setPassword(password)
	dbc := r.db.Create(&user)

	if dbc.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(dbc.Error, &pgErr) {
			if pgErr.Code == PgUniqueViolationErrorCode {
				return 0, fmt.Errorf("%w", repo.ErrUserExists)
			}
			return 0, pgErr
		}
	}
	return user.ID, nil
}

func (r *Repo) CheckPassword(login, password string) (bool, error) {
	var user User
	dbc := r.db.Where("username = ?", login).First(&user)
	if dbc.Error != nil {
		if errors.Is(dbc.Error, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf("%w", repo.ErrUserNotFound)
		}
		return false, dbc.Error
	}
	return user.checkPassword(password), nil
}
