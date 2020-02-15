// A map maps keys to values.

// The zero value of a map is nil. A nil map has no keys, nor can keys be added.

// The make function returns a map of the given type, initialized and ready for use. 

package main

import "fmt"

func main(){
  m := make(map[string] int)

  // Only ""(double quotes) allowed
  m["k1"] = 1
  m["k2"] = 1
  fmt.Println("map:", m)

  v1 := m["k1"]
  fmt.Println("v1: ", v1)

  fmt.Println("len:", len(m))

  delete(m, "k2")
  fmt.Println("map:", m)

  _, prs := m["k2"]
  fmt.Println("prs:", prs)

  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)
}
