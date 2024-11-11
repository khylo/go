/*
Write a program to sort an array of integers. The program should partition the 
array into 4 parts, each of which is sorted by a different goroutine. 
Each partition should be of approximately equal size. Then the main goroutine 
should merge the 4 sorted subarrays into one large sorted array. 

The program should prompt the user to input a series of integers. Each goroutine 
which sorts ¼ of the array should print the subarray that it will sort. When 
sorting is complete, the main goroutine should print the entire sorted list.

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
}
/*
	var input string
	fmt.Println("Enter a series of integers separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	inputSlice := strings.Split(input, " ")
	var intSlice []int
	for _, i := range inputSlice {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, j)
	}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
}*/