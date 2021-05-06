//Program in Go to demostrate Embedded Type relationship

package main
import "fmt"

type Person struct{
	name string
	age int32
	salary float64
}

type Designate struct{
	person Person
	designation string
}

func main(){
	//Instantiate the Person class
	person:=Person{name:"Projjal Gop",age:23,salary:50000}

	//Instantiate the Designate class 
	designate:=Designate{person:person,designation:"Data Analyst"}

	//--Details--
	fmt.Println("--Details about person--")
	fmt.Println("Name:",designate.person.name)
	fmt.Println("Age:",designate.person.age)
	fmt.Println("Salary:",designate.person.salary,"\n")

	fmt.Println("--Designation--")
	fmt.Println("Designation:",designate.designation)
}