package main

import "fmt"

func changeVal(ival *int, value int) {
	*ival = value
}

func main(){
	ival := 3
	fmt.Println("initial:", ival)
	
	changeVal(&ival, 20)
	fmt.Println("After value update:", ival)
	
	changeVal(&ival, 10)
	fmt.Println("After change:", ival)
}
