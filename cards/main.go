package main

func generateCard() string {
	return "Ace of spades"
}

func main() {
	// cards := []string{"Ace of diamonds", generateCard()}
	// Using a custom type
	// NOTE: Make sure to run "main.go" & "deck.go" together
	cards := deck{"Ace of diamonds", generateCard()}
	cards = append(cards, "King of hearts")

	// for _, card := range cards {
	// 	fmt.Println(card)
	// }
	// The "print" method now comes from "deck.go"
	// cards.print()

	newDeck := createDeck()

	hand, remainingDeck := deal(newDeck, 5)
	hand.print()
	remainingDeck.print()
}
