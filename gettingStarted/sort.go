package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	intArr := GetNumbers()
	intArr = SortData(intArr)
	fmt.Printf("Results %v\n", intArr)
}

func GetNumbers() []int {
	fmt.Println("Enter numbers to sort.")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	handleError(err)
	fmt.Printf("Read data: %v\n", input)
	strArr := strings.Split(strings.TrimSpace(input), " ")
	fmt.Printf("Split into %v\n", strArr)
	intArr := make([]int, 0, len(strArr))
	fmt.Printf("Making intSlice: %v\n", len(strArr))
	for _, num := range strArr {
		n, err := strconv.Atoi(strings.TrimSpace(num))
		handleError(err)
		intArr = append(intArr, n)
	}
	return intArr
}

func SortData(intArr []int) []int {
	var wg sync.WaitGroup
	rng := len(intArr)/4
	if len(intArr)%4 !=0
	    rng++
	var arrayOfSlices [4][]int
	for i := 0; i < 4; i++ {
		top := (i + 1) * rng
		if top >= len(intArr) {
			top = len(intArr)
		}
		fmt.Printf("Slice iteration %d from %d to %d\n", i, (i * rng), top)
		arrayOfSlices[i] = intArr[i*rng : top]
		wg.Add(1)
		go sortData(arrayOfSlices[i], &wg)
	}
	wg.Wait()
	ret := make([]int, 0, len(intArr))
	for i := 0; i < 4; i++ {
		ret = append(ret, arrayOfSlices[i]...)
	}
	sort.Ints(ret)
	return ret
}

func sortData(intArr []int, wg *sync.WaitGroup) {
	fmt.Printf("  goroutine sorting %v\n", intArr)
	sort.Ints(intArr)
	wg.Done()
}

func handleError(err interface{}) {
	if err != nil {
		fmt.Printf("Problem reading data: \n", err)
		os.Exit(1)
	}
}
