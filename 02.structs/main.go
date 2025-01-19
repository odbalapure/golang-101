package main

import (
	"fmt"
)

type contact struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	age       int
	// contact   contact
	// No need to explicitly name the type
	contact
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updatName(name string) {
	// (*p).firstName = name
	// Dereferencing done automatically for receivers
	// No need for explicit dereferencing
	p.firstName = name
}

func main() {
	// 1. intializing a "struct"
	// om := person{firstName: "Om", lastName: "Balapure", age: 29}
	// Relying upon the order of definition
	om := person{
		firstName: "Om",
		lastName:  "Balapure",
		age:       29,
		contact: contact{
			email: "ombalapure@gmail.com",
			zip:   442005,
		},
	}
	// fmt.Printf("%v %v is %v years old", om.firstName, om.lastName, om.age)
	// fmt.Printf("%+v", om)

	// 2. Another way of intializing "struct"
	// var hari person
	// hari.firstName = "Harish"
	// hari.lastName = "Balapure"
	// hari.age = 32

	// 3. Using pointers with "struct"
	// Gives memory address of the value, the variable is pointing at
	// NOTE: Go is a "pass by value" language, which *always* copies arguments that are passed to a function

	// omPointer := &om
	// omPointer.updatName("Omi")
	// omPointer.print()
	// NOTE: A shorter way to do this would be
	om.updatName("Omi")
	om.print()
}
