//Implementing function to return multiple values in Go

package main
import "fmt"

func min_max(x1 int,x2 int) (int,int){
	if(x1>x2){
		return x2,x1
	}
	return x1,x2
}

func main(){
	x,y:=min_max(56,51)
	fmt.Println("Max is ",x)
	fmt.Println("Min is ",y)
}