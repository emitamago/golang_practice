package main

func main() {
	// create new deck of cards
	deckOfCards := newDeckFromFile("my_first_deck")
	deckOfCards.print()

	// // create hands and remaining deck of cards
	// hand, remainingDeck := deal(deckOfCards, 5)

	// hand.print()
	// remainingDeck.print()
}
