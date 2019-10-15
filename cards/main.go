package main

import "fmt"

func main() {
	// direct way to declare variable
	// var card string = "Ace of Spades"

	// newer and easier way
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamond"
}
