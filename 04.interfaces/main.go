package main

/*
1. Interfaces are not generic types here like Java, C#
2. They are implicit, no need to manually tell that our custom type satisfies some interface
3. A contract to help us manage types but also remember:
	- Garbage IN -> Garbage OUT
	- If custom type implementation is broken then interfaces won't help us
*/

import "fmt"

// a struct "implements" an interface automatically if it provides all the methods declared in that interface.
// As long as the struct/name (a type) has the methods required by the interface, it satisfies the interface.
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}
type name string

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	om := name("Om")
	printGreeting(om)
}

// NOTE: Function overloading not supported in "Go"
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Remove the name inside the receiver if not used
// If the receiver won't be used atleast mention the type for attaching the interface method to this type
func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func (n name) getGreeting() string {
	return "My name is " + string(n)
}
