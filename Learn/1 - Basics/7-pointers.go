package main

import "fmt"

func main(){
  var p *int
  
  i := 43
  p = &i
  
  fmt.Println(*p)
  
  j := 44
  p = &j
  fmt.Println(*p)

  *p = 45
  fmt.Println(*p)
}
