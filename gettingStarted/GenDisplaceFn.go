package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
/*
assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a * t^2 + vo*t + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code.


*/

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func ParseFloat(s string) float64{
	ret,err := strconv.ParseFloat( strings.TrimSpace(s),64)
	if err!=nil{
		fmt.Printf("Cannot parse input parameter [%v] into a float: Error: %v\n",s, err)
		os.Exit(1)
	}
	return ret
}

func main() {
	fmt.Println("Please enter values (space seperated) for Acceleration, Initial Velocity, Initial Displacement")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	inputArray := strings.Split(strings.TrimSpace(input), " ")
	a := ParseFloat(inputArray[0])
	v0 := ParseFloat(inputArray[1])
	s0 := ParseFloat(inputArray[2])
	fun := GenDisplaceFn(a,v0,s0)
	for {
		fmt.Println("Please enter a time in seconds (-1 to finish)")
		scanner.Scan()
		input = scanner.Text()
		t := ParseFloat(input)
		if(t!=-1){
			fmt.Printf("You will have travelled %v units after %v seconds\n",fun(t),t)
		} else {
			break;
		}
	}
}