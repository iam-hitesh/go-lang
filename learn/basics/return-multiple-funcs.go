package main

import (
    "fmt"
)

func calculation(a, b int) (int, int) {
  return a * b, a + b;
}

func main() {
  a, b := calculation(2, 3)
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(calculation(2, 3))
}
