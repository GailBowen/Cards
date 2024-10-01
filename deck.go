package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handsize int) (deck, deck) {

	d = append(d, "cake plz")
	return d[:handsize], d[handsize:]

}

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}

func (d deck) toString() string {

	deskAsString := []string(d)

	return strings.Join(deskAsString, ",")

}

func (d deck) SaveToFile(fileName string) {
	cardsAsString := d.toString()
	cardsAsByteSlice := []byte(cardsAsString)

	err := os.WriteFile(fileName, cardsAsByteSlice, 0644)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func ReadFromFile(fileName string) deck {

	bytesFromFile, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}

	stringFromBytes := string(bytesFromFile)

	stringArray := strings.Split(stringFromBytes, ",")

	return deck(stringArray)
}
