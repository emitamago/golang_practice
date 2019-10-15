package main

func main() {
	// create new deck of cards
	deckOfCards := newDeck()

	// create hands and remaining deck of cards
	hand, remainingDeck := deal(deckOfCards, 5)

	hand.print()
	remainingDeck.print()
}
