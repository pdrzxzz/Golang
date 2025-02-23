package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type deck
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil { //if there's an error
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//
	return deck(strings.Split(string(bs), ","))
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert a deck into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
