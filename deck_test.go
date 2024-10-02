package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Error: expected 52 cards but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Error: expected Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Error: expected King of Clubs but got %v", d[0])
	}
}

func TestSaveToFileThenLoad(t *testing.T) {
	//delete all instances of TestFile
	os.Remove("_testFile")

	d := newDeck()

	d.saveToFile("_testFile")

	nd := readFromFile("_testFile")

	if len(nd) != 52 {
		t.Errorf("Deck should have been 52 but was %v", len(nd))
	}

	os.Remove("_testFile")

}
