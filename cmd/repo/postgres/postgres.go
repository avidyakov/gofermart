package postgres

import (
	"gophermart/cmd/config"
	"gophermart/cmd/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(config *config.Config) repo.Repo {
	db, err := gorm.Open(postgres.Open(config.DatabaseURI), &gorm.Config{})
	handleError(err)

	err = db.AutoMigrate(&User{})
	handleError(err)

	return &Repo{
		db: db,
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
