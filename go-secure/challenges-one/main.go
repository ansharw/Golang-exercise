package main

import "challenges-one/routers"

func main() {
	route := routers.Router()
	route.Run(":8080")
}