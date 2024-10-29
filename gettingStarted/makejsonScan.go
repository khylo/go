package main

import (
	"fmt"
	"encoding/json"
)

/*
Problem with this was Scan breaks on a space so doesn't really work
Write a program which prompts the user to first enter a name, and then enter an address. 
Your program should create a map and add the name and address to the map using the keys
 “name” and “address”, respectively. 
 Your program should use Marshal() to create a JSON object from the map, and then your program 
 should print the JSON object.
 */

 func main(){
	var name, address string
	fmt.Println("Enter a name")
	fmt.Scan(&name)
	fmt.Println("Enter address")
	fmt.Scan(&address)
	data := map[string]string{
		"name":name,
		"address":address,
	}
	byteArray,_ := json.Marshal(data)
	fmt.Printf("Json = %v", string(byteArray))
 }