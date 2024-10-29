package main

import "fmt"

func main() {
        str := "Hello, 世界!"

        for i := 0; i < len(str)-1; i++ {
                fmt.Printf("%c, next char is %c\n", str[i], str[i+1])
        }
		fmt.Println()
		fmt.Println("Using Runes")
		var prev rune
		for i, char := range str {
			if i > 0 {
				fmt.Printf("%c, next char is %c\n", prev, char)
			}
			prev = char
		}
}