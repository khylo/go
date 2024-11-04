/*
Read.go
Uses ReadFile from ioutil.. Not good for big files sine it reads it all to memory
 */
 package main

 import (
	"fmt"
	"io/ioutil"
	"strings"
 )

 type nameStruct struct{
	 fname string
	 lname string
 }

func main(){
	var filename string
	fmt.Println("Enter name of file")
	fmt.Scan(&filename)
	data,err := ioutil.ReadFile(filename)
	if(err != nil){
		fmt.Printf("Error reading %s. Error : %v",filename, err)
		return
	}
	//Split into lines
	lines := strings.Split(string(data),"\n")	
	nameSlice := make([]nameStruct, 0, len(lines))
	for _,line := range lines{
		names := strings.Split(line, " ")
		nameInst := nameStruct{fname: names[0], lname: names[1]}
		nameSlice = append(nameSlice, nameInst)
		//nameSlice.append(nameStruct{fname: names[0], lname: names[1]})
	}
	fmt.Printf("Read %d names into Slice of Structs\n",len(nameSlice))
	for _,name := range nameSlice{
		fmt.Printf("FirstName: %s  LastName: %s\n",name.fname, name.lname)
	}

	
}