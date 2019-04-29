package deckofcards

import (
	"strings"
)

//Cards A type that represents a slice of cards to make the code easier to read and easier to use in other type definitions
type Cards []*Card

func (z Cards) swapCards(i, j int) {
	z[j], z[i] = z[i], z[j]
}

func (z Cards) String() string {
	var printString []string
	for _, card := range z {
		if !card.drawn {
			printString = append(printString, card.String())
		}
	}
	return strings.Join(printString, "\n")
}
