package repo

type Repo interface {
	CreateUser(username, password string) error
}
