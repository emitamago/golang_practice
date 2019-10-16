package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 48 {
		t.Errorf("Expected deck of 48 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card will be Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clovers" {
		t.Errorf("Expected last card will be King of Clovers but  got  %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 48 {
		t.Errorf("Expected deck of 48 but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
