package main

import ("fmt"
	"sort"
	"strconv")

func main(){
	var input string
	var slice = make([]int,0,3);
	for  {
		fmt.Println("Enter an Integer. (Enter X to finish)");
		fmt.Scan(&input)
		if(input=="X"){
			break;
		}
		i,err := strconv.Atoi(input)
		if err != nil{
			fmt.Printf("Skipping non integer %s\n",input)
		}else{
			slice = append(slice, i)
			sort.Ints(slice)	
		}
		fmt.Printf("Slice now = %v\n", slice)	
	}
	fmt.Printf("Finished. Final Slice = %v\n", slice)
}