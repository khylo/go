/*
Read.go
Uses ReadFile from ioutil.. Not good for big files sine it reads it all to memory


Write a program which reads information from a file and represents it in a slice of structs. 
Assume that there is a text file which contains a series of names. 
Each line of the text file has a first name and a last name, in that order, separated by a 
single space on the line.

Your program will define a name struct which has two fields,
 fname for the first name, and
 lname for the last name. 
Each field will be a string of size 20 (characters).


Your program should prompt the user for the name of the text file. 
Your program will successively read each line of the text file and create a struct which contains 
the first and last names found in the file. Each struct created
 will be added to a slice, and after all lines have been read from the file, 
 your program will have a slice containing one struct for each line in the file. 
 After reading all lines from the file, your program should iterate through your slice of structs and
 print the first and last names found in each struct.
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
	data,_ := ioutil.ReadFile(filename)
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