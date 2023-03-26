package main

import (
	"challenges-three/routers"
)

func main() {
	r := routers.StartApp()
	r.Run(":8080")
}
