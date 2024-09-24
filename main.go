package main

import "fmt"

func main() {

	cards := newDeck()

	handsize := 3

	myHand, myLeftOver := deal(cards, handsize)

	myHand.print()

	fmt.Println("That is hand printed")

	myLeftOver.print()

	c := color("purple")

	fmt.Println(c.describe("is amazing"))

	cards.print()

}
