//  An array has a fixed size. A slice, on the other hand, is a dynamically-sized, 
// flexible view into the elements of an array. In practice, slices are much more 
// common than arrays. 

// https://tour.golang.org/moretypes/7


package main

import (
  "fmt"
)

func main() {
  s := make([]int, 3)
  fmt.Println(s)

  s[0] = 1
  s[1] = 2
  s[2] = 3
  fmt.Println(s)
  fmt.Println(s[2])

  // len is used for getting length of slice
  fmt.Println("lenght of slice s is ", len(s))

  s = append(s, 4)
  s = append(s, 5)
  fmt.Println("lenght of slice s is ", len(s))

  c := make([]int, len(s))
  copy(c, s)
  fmt.Println("cpy:", c)

  l := s[2:5]
  fmt.Println("sl1:", l)

  l = s[:5]
  fmt.Println("sl2:", l)

  l = s[2:]
  fmt.Println("sl3:", l)

  t := []string{"g", "h", "i"}
  fmt.Println("dcl:", t)

  twoD := make([][]int, 3)
  for i := 0; i < 3; i++ {
      innerLen := i + 1
      twoD[i] = make([]int, innerLen)
      for j := 0; j < innerLen; j++ {
          twoD[i][j] = i + j
      }
  }
  fmt.Println("2d: ", twoD)
}
