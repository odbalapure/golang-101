package main

import (
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := createDeck()

	// Success scenario
	// if len(d) == 28 {
	// 	t.Log("Deck has correct no. of cards!")
	// }

	if len(d) != 28 {
		t.Error("Expected deck length of 28, but got", len(d))
	}

	if d[0] != "Ace of Spades" {
		// "f" is required for formatted string
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Joker of Clubs" {
		t.Errorf("Expected last card to be Joker of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileNewDeckToFile(t *testing.T) {
	// Clean up files created by previous tests
	os.Remove("_decktesting")

	deck := createDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(deck) != len(loadedDeck) {
		t.Errorf("Expected deck length of %v, but got %v", len(deck), len(loadedDeck))
	}

	for i, card := range deck {
		if card != loadedDeck[i] {
			t.Errorf("Expected card at position %d to be %v, but got %v", i, card, loadedDeck[i])
		}
	}

	os.Remove("_decktesting")
}
