package main

import "fmt"

func main() {
        str := "Hello, World!"

        // Loop through the string using a for loop
        for i := 0; i < len(str); i++ {
                fmt.Printf("%c ", str[i])
        }
        fmt.Println()

        // Loop through the string using a for-range loop
        for _, char := range str {
                fmt.Printf("%c ", char)
        }
        fmt.Println()
}