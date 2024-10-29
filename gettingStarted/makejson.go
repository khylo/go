package main

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
	"strings"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address. 
Your program should create a map and add the name and address to the map using the keys
 “name” and “address”, respectively. 
 Your program should use Marshal() to create a JSON object from the map, and then your program 
 should print the JSON object.
 */

 func main(){
	fmt.Println("Enter a name")
	reader := bufio.NewReader(os.Stdin)
	name,_ := reader.ReadString('\n')
	fmt.Println("Enter address")
	address,_ := reader.ReadString('\n')
	data := map[string]string{
		"name":strings.TrimSpace(name),
		"address":strings.TrimSpace(address),
	}
	jsonStr,_ := json.Marshal(data)
	fmt.Printf("Json = %v", string(jsonStr))
 }