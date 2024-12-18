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
	addressMap := make(map[string]string)
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a name")
	reader.Scan()
	addressMap["name"] = strings.TrimSpace(reader.Text())

	fmt.Println("Enter address")
	reader.Scan()
	addressMap["address"] = strings.TrimSpace(reader.Text())
	
	jsonStr,_ := json.Marshal(addressMap)
	fmt.Printf("Json = %v", string(jsonStr))
 }