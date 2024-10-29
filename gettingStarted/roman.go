package main

import "fmt"
// my orig = Go 0 ms 4.5 MB ..  100.00%  Memory beats 14.8%
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
	var count int
    //for i, char := range s {
	for i := 0; i < len(s); i++ {	
		switch s[i] {
		case 'M':
			count+=1000
		case 'D':
			count+=500
		case 'C':
			if(next(s,i)=="D"){
				i=i+1
				count+=400
			}else if(next(s,i)=="M"){
				i=i+1
				count+=900
			}else{
				count+=100
			}
		case 'L':
			count+=50;
		case 'X':
			if(next(s,i)=="L"){
				i=i+1
				count+=40
			}else if(next(s,i)=="C"){
				i=i+1
				count+=90
			}else{
				count+=10
			}
		case 'V':
			count+=5;
		case 'I':
			if(next(s,i)=="V"){
				i=i+1
				count+=4
			}else if(next(s,i)=="X"){
				i=i+1
				count+=9
			}else{
				count+=1
			}
		}
	}
	return count
}

func next(s string, i int) string{
	if(i+1>=len(s)){
		return "";
	}
	return string(s[i+1])
}