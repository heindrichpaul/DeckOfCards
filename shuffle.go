package deckofcards

import "math/rand"

func shuffle(z Cards) Cards {
	for i := 1; i < len(z); i++ {
		// Create a random int up to the number of cards
		r := rand.Intn(i + 1)

		// If the the current card doesn't match the random
		// int we generated then we'll switch them out
		if i != r {
			z.swapCards(i, r)
		}
	}
	return z
}
