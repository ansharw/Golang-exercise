package main

import "MyGram/routers"

func main() {
	r := routers.StartApp()
	r.Run(":8080")
}
