/*
If we ran this sequentially the value of count would be 0, and the program would 
block until all loops are complete

However this program has a race condition when executed concurrently.
Both increment and decrement loops are 'racing' against each other to alter the 
value of the 'count' variable.
The order of execution is not guaranteed, and so if we stop exution after 1 second
the value of count will be a random/ non deterministic  number.

This is a race condition.
*/
package main

import (
	"fmt"
	"time"
)

var count int

func main() {
 count=0
	for i := 0; i < 1000; i++ {
		go increment()
		go decrement()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}

func increment() {
	for i := 0; i < 1000; i++ {
		count++
	}
}
func decrement() {
	for i := 0; i < 1000; i++ {
		count--
	}
}