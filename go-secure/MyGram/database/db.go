package database

import (
	"fmt"
	"log"

	"MyGram/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "my_gram"
	// db       *gorm.DB
	// err      error
)

func GetConnection() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Error to connection database", err)
	}

	fmt.Println("Successfully connected to database")

	db.Debug().AutoMigrate(model.User{}, model.Photo{}, model.Comment{}, model.SocialMedia{})
	return db
}
