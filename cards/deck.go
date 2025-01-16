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

func createDeck() deck {
	cards := deck{}
	cardSuits := deck{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := deck{"Ace", "Two", "Three", "Four", "King", "Queen", "Joker"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
