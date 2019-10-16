package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// take deck and turn into string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save deck to hard drive with permission of 0666 anyone can read and write
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// read from file and save value to deck
func newDeckFromFile(filename string) deck {

	// reading from file
	bs, err := ioutil.ReadFile(filename)

	// check err and if there is, exit program
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// turn bs into slice of strings
	s := strings.Split(string(bs), ",")

	// turn s into deck
	return deck(s)
}

func (d deck) shuffle() {
	// create new source for random number seed generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		// generate ramdom number 0 to length of d -1
		newPosition := r.Intn(len(d) - 1)

		// swap element at i and at new position
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
