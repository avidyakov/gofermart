package handlers

import (
	"gophermart/cmd/config"
	"gophermart/cmd/repo"
)

type Handlers struct {
	repo repo.Repo
	conf *config.Config
}

func New(repo repo.Repo, conf *config.Config) *Handlers {
	return &Handlers{
		repo: repo,
		conf: conf,
	}
}
