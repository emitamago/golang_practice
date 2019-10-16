package main

func main() {
	// create new deck of cards
	deckOfCards := newDeckFromFile("my_first_deck")
	deckOfCards.shuffle()

	// // create hands and remaining deck of cards
	hand, _ := deal(deckOfCards, 5)

	hand.print()
}
