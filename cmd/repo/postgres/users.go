package postgres

func (r *PostgresRepo) CreateUser(username, password string) error {
	user := User{
		Username: username,
	}
	user.setPassword(password)

	err := r.db.Create(&user).Error
	return err
}
