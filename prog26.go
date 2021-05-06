//Program to implement variadic parameter function in Go

package main 
import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
	total += v
	}
	return total
}

func main(){
	fmt.Println(add(3,4,5,6,7,8))
}