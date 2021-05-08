//Program to write file contents in Go  (IO part 2)
package main

import (
"os"
)

func main() {
file, err := os.Create("write-test.txt")
if err != nil {
// handle the error here
return
}
defer file.Close()
file.WriteString("Testing how to write to file after creation.")
}