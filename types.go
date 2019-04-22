package deckofcards

//Cards A type that represents a slice of cards to make the code easier to read and easier to use in other type definitions
type Cards []*Card

func (z Cards) swapCards(i, j int) {
	z[j], z[i] = z[i], z[j]
}
