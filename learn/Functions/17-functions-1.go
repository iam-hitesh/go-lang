package main

import (
    "fmt"
)

func add(x int, y int) int {
  return x + y
}

func muliplication(a, b int) int {
  return a * b;
}

func main() {
  fmt.Println(add(2, 3))
  fmt.Println(muliplication(2, 3))
}


/**
// Output

5

**/
