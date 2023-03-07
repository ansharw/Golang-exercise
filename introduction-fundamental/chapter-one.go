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
	// menampilkan nilai 21
	fmt.Printf("%v\n", i)
	// menampilkan tipe data dari variabel i
	fmt.Printf("%T\n",i)
	// menampilkan tanda %
	fmt.Printf("%%\n")
	// menampilkan nilai boolean j : true
	fmt.Printf("%t\n", j)
	// menampilkan nilai boolean j : true
	fmt.Printf("%t\n", j)
	// menampilkan unicode russia : Я (ya)
	fmt.Printf("%c\n", r)
	// menampilkan nilai base 10 : 21
	fmt.Printf("%d\n", i)
	// menampilkan nilai base 8 :25 
	fmt.Printf("%o\n", i)
	// menampilkan nilai base 16 : f 
    fmt.Printf("%x\n", k)
	// menampilkan nilai base 16 : F 
    fmt.Printf("%X\n", k)
	// menampilkan unicode karakter Я : U+042F
	fmt.Printf("%U\n", r)
	// var k float64 = 123.456; menampilkan float : 123.456000
	fmt.Printf("%f\n", f)
	// menampilkan float scientific : 1.234560E+02
	fmt.Printf("%e\n", f)
}
