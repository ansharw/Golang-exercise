package database

import (
	"MyGram/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "my_gram"
)

func GetConnection() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Error to connection database", err)
	}

	log.Println("Successfully connected to database")

	db.Debug().AutoMigrate(model.User{}, model.Photo{}, model.Comments{}, model.SocialMedia{})
	return db
}
