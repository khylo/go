package main

import "fmt"

func main() {
	fmt.Println("Enter a FP number?")
	var fp float64
	num, err := fmt.Scan(&fp)
	if err != nil {
		fmt.Printf("Error occurred in %d entries = %s", num, err)
	} else {
		var trunc = int64(fp)
		fmt.Printf("Truncated number is %d", trunc)
	}
}
