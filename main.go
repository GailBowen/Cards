package main

func main() {

	cards := deck{"Ace of Diamonds", newCard()}

	cards = append(cards, "Jack of Diamonds")

	cards.print()

}

func newCard() string {
	return "Five of Spades"
}
