package main

import "belajar-gin/routers"

func main() {
	var PORT = ":8081"

	routers.StartServer().Run(PORT)
}