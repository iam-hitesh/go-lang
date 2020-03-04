//fmt.Println prints the passed in variables’ values and appends a newline. 
// fmt.Printf is used when you want to print one or multiple values using a defined format specifier.

package main

import "fmt"

func main(){
  name := "Hitesh"
  aka := fmt.Sprintf("Number %d", 6)
  fmt.Printf("%s is also known as %s",name, aka)
}
