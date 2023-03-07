package main

import (
	"fmt"
)

func main() {
	i := 21
	j := true
	k := 15
	var r rune = 'Я'
	var f float64 = 123.456
	fmt.Printf("%v\n", i)
	fmt.Printf("%T\n",i)
	fmt.Printf("%%\n")
	fmt.Printf("%t\n", j)
	fmt.Printf("%t\n", j)
	fmt.Printf("%c\n", r)
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
    fmt.Printf("%x\n", k)
    fmt.Printf("%X\n", k)
	fmt.Printf("%U\n", r)
	fmt.Printf("%f\n", f)
	fmt.Printf("%e\n", f)
}
