//Program to enter temperature in deg C and 
//convert to deg F

package main
import "fmt"

func main(){
	fmt.Print("Enter deg C temperature:")
	var temp float64
	fmt.Scanf("%f",&temp)
	output:=temp*9/5+32
	fmt.Println("Temperature in deg F:",output)
}