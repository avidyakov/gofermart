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

	autoMigrateModels(db, &User{}, &Order{}, &Transaction{})

	return &Repo{
		db: db,
	}
}

func autoMigrateModels(db *gorm.DB, models ...interface{}) {
	for _, model := range models {
		err := db.AutoMigrate(model)
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
