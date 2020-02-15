// Go’s structs are typed collections of fields. 
// They’re useful for grouping data together to form records.

package main

import "fmt"

type person struct{
  name string
  age int
}

func NewPerson (name string) *person{
  p := person{name: name}
  p.age = 42

  return &p
}

func main() {
  fmt.Println(person{"Bob", 20})
  fmt.Println(person{name: "Yadav", age:22})
  
  fmt.Println(&person{name: "Ann", age: 40})
  fmt.Println(NewPerson("Hitesh"))

  s := person{name: "Sean", age: 50}
  fmt.Println(s.name)

  sp := &s
  fmt.Println(sp.age)

  sp.age = 51
  fmt.Println(sp.age)
}
