// Creates an excutable
package main

// Import package
import "fmt"

// var name string = "Om"
// var nums = []int{1, 2, 3}
// var strs = []string{"Om", "Hari"}

func main() {
	// 1. Creating variables
	// var card string = "Ace of spades"
	// Type cohersion is being done here
	// NOTE: Reassignment of card = "" won't work
	// card := "Ace of spades"
	// fmt.Println(card, name)

	// 2. Calling a function
	// newCard := generateCard()
	// fmt.Println(newCard)

	// 3. Creating slice and looping
	// var cards = []string{"Ace of diamonds", generateCard()}

	cards := []string{"Ace of diamonds", generateCard()}
	cards = append(cards, "King of hearts")

	for _, card := range cards {
		fmt.Println(card)
	}
}

func generateCard() string {
	return "Ace of spades"
}
