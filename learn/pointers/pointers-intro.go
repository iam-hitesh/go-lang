package main

import "fmt"

func changeVal(ival *int, value int) {
	*ival = 0
}

func main(){
	ival := 3
	fmt.Println("initial:", ival)
	
	changeVal(&ival, 0)
	fmt.Println("After value update:", ival)
	
	changeVal(&ival, 1)
	fmt.Println("After chnage:", ival)
}
