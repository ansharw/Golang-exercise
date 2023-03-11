package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Data3 struct {
	Value int
}

type Data4 struct {
	Value string
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 8; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			var data interface{}
			if i%2 == 0 {
				time.Sleep(2 * time.Second)
				data = Data3{Value: rand.Intn(100)}
			} else {
				time.Sleep(2 * time.Second)
				data = Data4{Value: fmt.Sprintf("Data %d", rand.Intn(100))}
			}

			mu.Lock()
			defer mu.Unlock()
			fmt.Println(data)
		}(i)
	}

	wg.Wait()
}
