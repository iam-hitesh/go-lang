// The range form of the for loop iterates over a slice or map.

// When ranging over a slice, two values are returned for each iteration. 
// The first is the index, and the second is a copy of the element at that index. 

// https://tour.golang.org/moretypes/16

package main

import "fmt"

var pow = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
  for index, value := range pow {
    fmt.Println(index, value)
  }
  
  pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
