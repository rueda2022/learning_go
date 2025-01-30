package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := new_deck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
}

func TestSaveTodeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := new_deck()
	deck.saveToFile("_decktesting")

	loadedDeck := new_deck_from_file("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
