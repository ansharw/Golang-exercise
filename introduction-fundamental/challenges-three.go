package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input kalimat
	kalimat := "selamat malam"

	// Pecah kalimat menjadi array kata
	kata := strings.Split(kalimat, " ")

	// Looping kata dan print character
	for _, k := range kata {
		for _, h := range k {
			fmt.Printf("%c\n", h)
		}
		fmt.Printf("\n")
	}

	// input ke map[string]int
	hitung := make(map[string]int)
	
	for _, char := range kalimat {
		hitung[string(char)]++
	}
	
	fmt.Println(hitung)
}