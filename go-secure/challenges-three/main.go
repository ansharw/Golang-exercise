package main

import (
	"challenges-three/database"
	"challenges-three/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
