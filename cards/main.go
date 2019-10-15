package main

import "fmt"

func main() {
	deckOfCards := deck{"Ace of Diamonds", newCard()}
	deckOfCards = append(deckOfCards, "King of Spades")
	for i, card := range deckOfCards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamond"
}
