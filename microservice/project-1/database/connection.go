package database

import (
	"fmt"
	"log"
	"project-1/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "books_new"
)

func GetConnection() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Error to connection database", err)
	}

	fmt.Println("Successfully connected to database")

	db.Debug().AutoMigrate(model.Book{})
	return db
}
