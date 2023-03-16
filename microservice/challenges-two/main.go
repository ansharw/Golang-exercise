package main

import "challenges-two/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}