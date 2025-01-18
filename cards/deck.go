package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

// Every variable of "type deck" has access to "print" method
// Using "d" is a convention, not a hard and fast rule
// This is similar to "this" keyword in Javascript
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
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

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// NOTE: ReadFile returns a byte slice and an error
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading from : %s", filename)
		os.Exit(1)
	}

	// Alternate syntax
	// s := strings.Split(string(bytes), ",")
	// return deck(s)

	return strings.Split(string(bytes), ",")
}

func (d deck) shuffle() deck {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

	return d
}
