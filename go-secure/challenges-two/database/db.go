package database

import (
	"challenges-two/models"
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
	dbname   = "simple_api"
	db       *gorm.DB
	err      error
)

// func GetConnection() *gorm.DB {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Error to connection database", err)
// 	}

// 	fmt.Println("Successfully connected to database")

// 	db.Debug().AutoMigrate(models.User{}, models.Product{})
// 	return db
// }

func StartDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Error to connection database", err)
	}

	fmt.Println("Successfully connected to database")

	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
