package postgres

import (
	"gophermart/cmd/config"
	"gophermart/cmd/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo(config *config.Config) repo.Repo {
	db, err := gorm.Open(postgres.Open(config.DatabaseURI), &gorm.Config{})
	handleError(err)

	err = db.AutoMigrate(&User{})
	handleError(err)

	return &PostgresRepo{
		db: db,
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
