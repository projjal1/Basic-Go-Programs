//Program to demonstrate use of constructor in struct in Go

package main
import "fmt"

type Const struct{
	//Class variables initialized with 0 (default values)
	a,b,c float64
	d int64
}

func main(){
	//one way of instantiation
	c1:=new(Const)
	//other way of instantiation
	c2:=Const{a:1,b:2,c:3,d:5}

	//--C1 object--
	fmt.Println("Printing paramters for c1 object")
	fmt.Println(c1.a,c1.b,c1.c,c1.d)


	//--C2 object--
	fmt.Println("Printing paramters for c2 object")
	fmt.Println(c2.a,c2.b,c2.c,c2.d)
}
