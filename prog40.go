//Program to read a file contents in Go  (IO part 2)
package main

import (
"fmt"
"io/ioutil"
)

func main() {
//Open file pointer
bs, err := ioutil.ReadFile("test.txt")
if err != nil {
return
}

//Read the string content
str := string(bs)
fmt.Println(str)
}