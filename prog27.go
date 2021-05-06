//Program to implement Closure(Lambda function) in Go (PART 1)

package main
import "fmt"

func main() {
 	x := 0
 	increment := func() int {
		x++
 		return x
 	}
 	fmt.Println(increment())
 	fmt.Println(increment())
}