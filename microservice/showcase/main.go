package main

import (
	"showcase/database"
	"showcase/routers"
)

func main() {
	db := database.GetConnection()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	route := routers.Route(db)
	route.Run(":8080")
}