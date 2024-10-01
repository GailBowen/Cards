package main

func main() {

	cards := newDeck()

	cards.SaveToFile("camel")

	moreCards := ReadFromFile("camel")

	moreCards.print()
}
