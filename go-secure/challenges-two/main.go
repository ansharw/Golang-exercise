package main

import (
	"challenges-two/routers"
	"challenges-two/database")

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}