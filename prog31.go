//Program to demonstrate the use of function calling as struct methods explicitly

package main 
import "fmt"

type Circle struct{
	r float64
}

func (c *Circle) perimeter() float64{
	return 2*3.14*c.r
}

func (c *Circle) area() float64{
	return 3.14*c.r*c.r
}

func main(){
	//Instantiation
	obj:=Circle{r:5.6}

	fmt.Println("Area of circle is ")
	fmt.Println(obj.area())

	fmt.Println("Perimeter of circle is  ")
	fmt.Println(obj.perimeter())
}