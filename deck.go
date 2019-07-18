package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	//returns elements from 0(inclusive) to the handSize(exclusive)
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//[]string(d) converts deck d into a slice of strings 
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error  {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}