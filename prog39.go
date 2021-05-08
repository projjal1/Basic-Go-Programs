//Program to read a file contents in Go  (IO part 1)
package main

import (
"fmt"
"os"
)

func main() {
//Open file pointer
file, err := os.Open("test.txt")
if err != nil {
// handle the error here
return
}
defer file.Close()

// get the file size
stat, err := file.Stat()
if err != nil {
return
}
fmt.Println("The file size (in bytes) is:",stat.Size());

// read the file
bs := make([]byte, stat.Size())
_, err = file.Read(bs)
if err != nil {
return
}

//Read the string
str := string(bs)
fmt.Println(str)
}