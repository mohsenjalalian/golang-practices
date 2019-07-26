package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T)  {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected len 16 but got %v", len(d))
	}

	if (d[0] != "ace of spades") {
		t.Errorf("Expected ace of spades but got %v", d[0])
	}

	if (d[len(d) - 1] != "four of clubs") {
		t.Errorf("Expected four of clubs but got %v", d[len(d) - 1	])
	}
}

func TestSaveToDeckAndnewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_deckstring")

	loadedDeck := newDeckFromFile("_deckstring")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck %v", len(loadedDeck))
	}

	os.Remove("_deckstring")
}