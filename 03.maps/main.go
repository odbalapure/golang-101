package main

import "fmt"

func main() {
	// Declaration #1
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// Declaration #2
	// var colors map[string]string

	// Declaration #3
	// colors := make(map[string]string)

	colors := map[int]string{
		0: "#ff0000",
		1: "#4bf745",
	}
	// Adding a key-value pair
	colors[2] = "#ffffff"
	// Deleting a key-value pair
	delete(colors, 2)

	fmt.Println(colors)
	printMap(colors)

	// When to use maps vs structs
	// 1. Map is a reference type
	// 2. Keys can be added/removed in a Map
	// 3. In a Map, all the keys and all the values are of the same type
	// 4. We can iterate over map as the keys are indexed
}

func printMap(mapper map[int]string) {
	for key, value := range mapper {
		fmt.Println(key, "->", value)
	}
}
