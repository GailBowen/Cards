package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	cards := newDeck()

	fmt.Println(cards)

	//Convert deck to string.
	cardsAsString := cards.toString()

	//Convert string to byte slice.
	cardsAsByteSlice := []byte(cardsAsString)

	err := os.WriteFile("Chickens", cardsAsByteSlice, 0644)

	if err != nil {
		log.Fatal(err)
	}

	bytesFromFile, readError := os.ReadFile("Chickens")

	if readError != nil {
		log.Fatal(readError)
	}

	//Convert byte slice to string
	stringFromBytes := string(bytesFromFile)

	//Convert string to string array
	stringArray := strings.Split(stringFromBytes, ",")

	//Convert string array to deck
	myDeck := deck(stringArray)

	fmt.Println("-------------------------------------------------")

	fmt.Println(myDeck)

}
