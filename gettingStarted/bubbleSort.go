package main
import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)
func main(){
	data := make([]int,0, 10)
	fmt.Println("Enter a space seperated list of integers")
	strData,_ := bufio.NewReader(os.Stdin).ReadBytes('\n')
	strDataArr := strings.Split(strings.TrimSpace(string(strData)), " ")
	for i:=0; i<len(strDataArr); i++{		
		num, _ := strconv.Atoi(strings.TrimSpace(strDataArr[i]))
		data = append(data, num)
	}
    BubbleSort(data)
	fmt.Printf("Sorted data = %v\n", data)
}

func BubbleSort(data []int){
    for i := 0; i < len(data); i++{
        for j:= 0;j<len(data)-1;j++{
            if data[j] > data[j+1]{
                Swap(data, j)
            }
        }
    }
}

func Swap(data []int, i int){
    temp := data[i]
    data[i] = data[i+1]
    data[i+1] = temp
}