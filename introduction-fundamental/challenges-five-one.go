package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Data1 struct {
	Value int
}

type Data2 struct {
	Value string
}

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func() {
		for i := 0; i < 4; i++ {
			time.Sleep(time.Duration(1))
			ch1 <- Data1{Value: rand.Intn(100)}
		}
	}()

	go func() {
		for i := 0; i < 4; i++ {
			time.Sleep(time.Duration(1))
			ch2 <- Data2{Value: fmt.Sprintf("Angka %d", rand.Intn(100))}
		}
	}()

	for i := 0; i < 8; i++ {
		select {
		case d1 := <-ch1:
			fmt.Println("Data 1:", d1)
		case d2 := <-ch2:
			fmt.Println("Data 2:", d2)
		}
	}
}
