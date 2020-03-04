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
	return p.name
}

func (p *Person) getAge() int {
	return p.age
}

func main(){
	p1 := Person{"Hitesh Yadav", 20, 170}
	p1.getPersonName()
	fmt.Println(personName(p1))
	fmt.Println("Age:", p1.getAge())
}