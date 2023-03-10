package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Printf("nilai i = %d\n", i)
    }

    for j := 0; j < 5; j++ {
        fmt.Printf("nilai j = %d\n", j)
    }

    var characters = []rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041E'}
    for j := 0; j <= 10; j++ {
        if j == 5 {
            for p, c := range characters {
                fmt.Printf("character %U '%c' starts at byte position %d\n", c, c, p*2)
            }
        } else {
            fmt.Printf("nilai j = %d\n", j)
        }
    }
}
