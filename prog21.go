//Program to create slice and copy content of one in another 

package main
import "fmt"

func main() {
 	slice1 := []int{1,2,3}
	//Creating a blank slice
 	slice2 := make([]int, 2)
	//Copy one slice content into another
 	copy(slice2, slice1)
 	fmt.Println(slice1, slice2)
}