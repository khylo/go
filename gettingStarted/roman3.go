package main

import "fmt"
// update 1 of  orig now using trick I spotted in other solns
func main() {
    doIt("I" )
	doIt("II")
	doIt("V" )
	doIt("IV")
	doIt("CC")
	doIt("DD")
	doIt("XL")
	doIt("XC")
	doIt("CM")
	doIt("M" )
}

func doIt(val string){
	fmt.Printf("%v = %v\n", val, romanToInt(val))
}


func romanToInt(s string) int {
	var count, last int
    //for i, char := range s {
	for i := 0; i < len(s); i++ {	
		switch s[i] {
		case 'M':
			count+=1000
		case 'D':
			count+=500
		case 'C':
			count+=100
		case 'L':
			count+=50;
		case 'X':
			count+=10
		case 'V':
			count+=5;
		case 'I':
			count+=1
			
		}
		if(i>0 && count-last){
			i=i+1
			count+=400
		}else if(next(s,i)=="M"){
			i=i+1
			count+=900
		}else{
	}
	return count
}

func next(s string, i int) string{
	if(i+1>=len(s)){
		return "";
	}
	return string(s[i+1])
}