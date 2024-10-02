package main

import "testing"

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
