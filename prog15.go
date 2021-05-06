//Program to check numbers between 1 and 10 as 
//Odd or Even 

package main
import "fmt"

func main(){
    for i:=1;i<=10;i++ {
        if i%2==0 {
            fmt.Println(i," is even")
		} else{
            fmt.Println(i," is odd")
		}
    }
}