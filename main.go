package main

func main() {

	//fmt.Println(rand.Intn(54))

	cards := newDeck()

	cards.shuffle()

	cards.print()

}
