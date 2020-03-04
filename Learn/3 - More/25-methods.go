// Type defined in Golang program can have methods associated with it, Go does not have classes but we can
// define methods on types.
// Remember: a method is just a function with a receiver argument.
// A good read to methods & interfaces https://medium.com/golangspec/methods-in-go-part-i-a4e575dff860

package main

import "fmt"

type Person struct{
	name string
	age int
	height int
}

// Methods are bound to receiver’s base type.
//  In this example, the getPersonName method has a receiver of type Person named p.
func (p Person) getPersonName() {
	fmt.Println(p.name)
}

// A simple function or we can say method without any base type we can pass any Person type.
func personName(p Person) string {
	p.name = "Hello World"
	return p.name
}

//  method with pointer receiver.
func (p *Person) getAge() int {
	p.age = 50
	return p.age
}

func main(){
	p1 := Person{"Hitesh Yadav", 20, 170}
	p1.getPersonName()
	fmt.Println(personName(p1))
	fmt.Println("Age:", p1.getAge())

	// Pointer indirection in methods
	p2 := &Person{"Some Name", 30, 175}
	p2.getPersonName()
	// Output will be 50 as we are updated the value in getAge() method
	fmt.Println("Age:", p2.getAge())
}