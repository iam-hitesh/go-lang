package main

import (
  "fmt"
)

func main() {
  var a[10] int
  fmt.Println(a)

  a[4] = 1
  fmt.Println(a)
  fmt.Println(a[4])

  b := [5]int{1,2,3,4,5}
  fmt.Println(b)

  var twoD [2][3]int
  for i := 0; i < 2; i++ {
      for j := 0; j < 3; j++ {
          twoD[i][j] = i + j
      }
  }
  fmt.Println("2d: ", twoD)
}
