//  An interface type is defined as a set of method signatures.
// Type defined in Golang program can have methods associated with it.

package main

import "fmt"

type Person struct{
	name string
	age int
	height int
}

type PersonInterface interface {
	getPersonName()
}

// Methods are bound to receiver’s base type.
func (p Person) getPersonName() {
	fmt.Println(p.name)
}

func (p *Person) getAge() int {
	return p.age
}

func main(){
	var p1 PersonInterface = Person{"Hitesh Yadav", 20, 170}
	p1.getPersonName()
	fmt.Println("Age:", p1.getAge())
}