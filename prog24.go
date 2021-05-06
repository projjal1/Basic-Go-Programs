//Program to implement a function to calculate average of float numbers

package main
import "fmt"

func average(xs []float64) float64{
	total:=0.0
	for _,val := range xs{
		total+=val
	}
	return total/float64(len(xs))
}

func main(){
	xs:= []float64{4.5,6.2,5.6,1.2,6.7,7,8.9};
	fmt.Println(average(xs))
}