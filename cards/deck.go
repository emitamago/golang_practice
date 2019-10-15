package main

import "fmt"

// Create a new type of 'deck'
//  it is slice of strings

type deck []string

// Create new deck and return it
func newDeck() deck {
	cards := deck{}

	// slices of suites and values to create string for an individual card
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clovers"}
	cardValue := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Jack",
		"Queen",
		"King"}

	// Iterate over the slices and append to 52 invidual card strings
	for _, suites := range cardSuites {
		for _, value := range cardValue {
			cards = append(cards, suites+" of "+value)
		}
	}
	return cards
}

// Print out individual card and index in deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Return 2 slices --- one with handsize element for beginning and
// the other with remaining cards
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
