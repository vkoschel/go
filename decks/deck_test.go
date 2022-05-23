package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("unexpected deck size")
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("deck in wrong order")
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("deck in wrong order")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("unexpected deck size")
	}

	os.Remove("_decktesting")

}
