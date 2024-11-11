/*
This prog checks the balance and withdraws money if there is enough money in the account,
However there is a delay after checking the balance so a race condition can occur and 
more can be withdrawn
*/
package main

import (
	"fmt"
	"time"
)

var x int
var balance float64 = 1000

func Balance() float64{
	return balance
}

func Withdraw(amount float64) float64{
	if Balance() < amount {
		fmt.Printf("Withdraw of %f failed: Insufficient balance: %f\n", amount, Balance())
		return 0
	}
	time.Sleep(100* time.Millisecond)
	balance = Balance() - amount
	fmt.Printf("Withdrawn %f\n", amount)
	return amount
}

func main() {
	fmt.Printf("Opening Balance = %f\n",Balance())
	for i:=0;i<3;i++ {
		go Withdraw(800)
	}
	time.Sleep(500* time.Millisecond)
	fmt.Printf("Closing Balance = %f", Balance());
}