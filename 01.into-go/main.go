package main

// func generateCard() string {
// 	return "Ace of spades"
// }

func main() {
	// cards := []string{"Ace of diamonds", generateCard()}
	// Using a custom type
	// NOTE: Make sure to run "main.go" & "deck.go" together
	// cards := deck{"Ace of diamonds", generateCard()}
	// cards = append(cards, "King of hearts")

	// For loop
	// for _, card := range cards {
	// 	fmt.Println(card)
	// }
	// The "print" method now comes from "deck.go"
	// cards.print()

	newDeck := createDeck()

	// Getting multiple values in "return" from a function
	// hand, remainingDeck := deal(newDeck, 5)
	// hand.print()
	// remainingDeck.print()

	// fmt.Println(newDeck.toString())

	// Write to file
	// newDeck.saveToFile("deck.txt")

	// Read from file
	// newDeckFromFile("deck.txt").print()

	newDeck.shuffle().print()
}
