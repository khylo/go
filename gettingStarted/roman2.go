package main

import "fmt"
// Solution from leet.. 
// my orig = Go 0 ms 4.5 MB ..  100.00%  Memory beats 14.8%

// this = 4ms  4.6MB
func main() {
    doIt("I" )
	doIt("II")
	doIt("V" )
	doIt("IV")
	doIt("CC")
	doIt("DD")
	doIt("XL")
	doIt("XC")
	doIt("CM")
	doIt("M" )
}

func doIt(val string){
	fmt.Printf("%v = %v\n", val, romanToInt(val))
}
func romanToInt(s string) int {
	sum := 0

	rim := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i, v := range s {
		sum += rim[string(v)]
		if i != 0 {
			if rim[string(s[i-1])] < rim[string(v)] {
				sum -= 2 * rim[string(s[i-1])]
			}
		}
	}

	return sum
}