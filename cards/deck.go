package main

import "fmt"

type deck []string

// Every variable of "type deck" has access to "print" method
// Using "d" is a convention, not a hard and fast rule
// This is similar to "this" keyword in Javascript
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
