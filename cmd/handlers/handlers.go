package handlers

import (
	"gophermart/cmd/config"
	"gophermart/cmd/loyalty"
	"gophermart/cmd/repo"
)

type Handlers struct {
	repo          repo.Repo
	conf          *config.Config
	accrualSystem loyalty.LoyaltySystem
}

func New(repo repo.Repo, conf *config.Config, accrual loyalty.LoyaltySystem) *Handlers {
	return &Handlers{
		repo:          repo,
		conf:          conf,
		accrualSystem: accrual,
	}
}
