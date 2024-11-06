go code,
coursera https://www.coursera.org/learn/golang-functions-methods/lecture/kNIkB/m2-1-1-first-class-values

Note
 slices are passed by refence (pointr is passed by value)

## Functions are 1st class
functional prog in go
* variables can be declared with a fn type
* Can be created dynamically (variables and fns)

### Variables as functions
```
var f func(int) int 
func inc(x int) int {
    return x+1
}

func main(){
    f =  inc // note no ().. means assigning not calling
    fmt.Printf("ans = %d",f(2))

}
```

### functions as args
fns can be passed as args
```
func applyIt(f func(int) int, val int) int{
    	return f(val)
}
```
### Anonymous
```
func main(){
    v := applyIt(func(x int) int{
        return x-1
    },2)
    fmt.Printf("ans = %d",v)
}
```
### returning fns
might create a fn with configurable params
One powerful use case for returning functions in Go is to create flexible and customizable logging functions.

Basic Logging Function:
```

func Logger(prefix string) func(string) {
    return func(msg string) {
        fmt.Printf("[%s] %s\n", prefix, msg)
    }
}
// to use
// Create a logger for error messages
errorLogger := Logger("ERROR")

// Create a logger for info messages
infoLogger := Logger("INFO")

// Use the loggers
errorLogger("Something went wrong!")
infoLogger("Process completed successfully.")
```

### Receiving Type
Implicit method Arguement... The obj is the arguement. passed by Val (maybe slow)
unless we use pointer receiver. no need to derefence.. compiler does it automatically\

rec all or none use pointer receivers..
 (mi *SmiInt )
```
type MyInt int

// Here we are adding a receiving type to our MyIn class. Note the (mi MyInt)  which differs from normal function naming.tells what class we are adding to      
// Double is the name of the function we are using on MyInt
func (mi MyInt) Double () int {

...
{
    v := MyInt(3)
    v.Double()
}
}
```