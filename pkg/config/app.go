package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	//#TODO: env var to get db infos

	database, err := gorm.Open("postgres", "host=localhost port=5433 user=books_db dbname=books_db password=12345678! sslmode=disable")

	if err != nil {
		panic(err)
	}

	db = database
}

func GetDB() *gorm.DB {
	return db
}
