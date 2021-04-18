//Program to implement array in Go (Part 3)

package main
import "fmt"

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64 = 0


	// _ means we won't use this variable in our script
	for _,val := range x {
		total += val
	}
	
	fmt.Println(total / 5)
}