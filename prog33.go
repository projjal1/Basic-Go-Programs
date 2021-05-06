//Program in Go to implement interface 

package main
import "fmt"

//Declaring Shape interface
type shape interface{
	area() float64
}

//Declaring struct of Circle, Rectangle, and Mutlishape
type MultiShape struct{
	shapes []shape
}

type Circle struct{
	r float64
}

type Rectangle struct{
	l,b float64
}

//Circle area function
func (c Circle) area() float64{
	return c.r*c.r*3.14
}

//Rectangle area function
func (r Rectangle) area() float64{
	return r.l*r.b
}

//Multishape function
func (m *MultiShape) area() float64{
	var total float64
	total=0
	for _, s := range m.shapes{
		total+=s.area()
	}
	return total
}

func main(){
	//Instantiating Multishape comprising Circle and Rectangle
	multishape := MultiShape{
		shapes: []shape{
			Circle{5},
			Rectangle{10, 10},
		},
	}

	fmt.Println("The total area including Circle and Rectangle is:")
	fmt.Println(multishape.area())
}