package postgres

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"gophermart/cmd/repo"
)

const (
	PgUniqueViolationErrorCode = "23505"
)

func (r *PostgresRepo) CreateUser(login, password string) error {
	user := User{
		Username: login,
	}
	user.setPassword(password)
	dbc := r.db.Create(&user)

	if dbc.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(dbc.Error, &pgErr) {
			if pgErr.Code == PgUniqueViolationErrorCode {
				return fmt.Errorf("%w", repo.ErrUserExists)
			}
			return pgErr
		}
	}
	return nil
}
