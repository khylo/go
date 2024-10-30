package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Enter a string to search :")
	reader := bufio.NewReader(os.Stdin)
	input,err := reader.ReadString('\n')
	if err!=nil {
		fmt.Printf("Error occurred %s\n",err)
		return;
	}
	input = strings.TrimSpace(input)
	input=strings.ToLower(input)
	if(strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a")){
		fmt.Printf("Found!")
	}else{		
		fmt.Printf("Not Found!")
	}
}