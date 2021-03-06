package deckofcards

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck(1)
	if deck == nil {
		t.Logf("Failed to create deck\n")
		t.FailNow()
	}
	if !deck.Success {
		t.Logf("Deck not properly initialized. Expected a success on a successful creation\n")
		t.FailNow()
	}
	if deck.Remaining != len(deck.cards) {
		t.Logf("Deck not properly initialized. Expected amount of cards remaining and the length of the cards slice to be equavalent after a new deck is created.\n")
		t.FailNow()
	}
	if strings.EqualFold(deck.DeckID, "\n") {
		t.Logf("Deck not properly initialized. Expected a non empty DeckID\n")
		t.FailNow()
	}
	if deck.Shuffled {
		t.Logf("Deck not properly initialized. Expected an unshuffled deck\n")
		t.FailNow()
	}
	for _, card := range deck.cards {
		if !strings.EqualFold(deck.DeckID, card.DeckID) {
			t.Logf("The card's DeckID differs from the deck it belongs to.\n")
			t.FailNow()
		}
	}
}

func TestNewDeckWithNegativeNumber(t *testing.T) {
	deck := NewDeck(-1)
	if deck == nil {
		t.Logf("Failed to create deck\n")
		t.FailNow()
	}
	if deck.Success {
		t.Logf("Deck not properly initialized failed deck. Expected a false on a unsuccessful creation\n")
		t.FailNow()
	}
	if deck.Remaining > 0 {
		t.Logf("Deck not properly initialized. Expected zero cards as it failed do create a deck\n")
		t.FailNow()
	}
	if strings.EqualFold(deck.DeckID, "") {
		t.Logf("Deck not properly initialized. Expected a non empty DeckID\n")
		t.FailNow()
	}
	if deck.Shuffled {
		t.Logf("Deck not properly initialized. Expected an unshuffled deck\n")
		t.FailNow()
	}
}

func TestInjectionOfUnsupportedSuit(t *testing.T) {
	suits[0] = "F"
	deck := NewDeck(1)
	if deck == nil {
		t.Logf("Failed to create deck\n")
		t.FailNow()
	}
	if deck.Success {
		t.Logf("Deck not properly initialized failed deck. Expected a false on a unsuccessful creation\n")
		t.FailNow()
	}
	if deck.Remaining > 0 {
		t.Logf("Deck not properly initialized. Expected zero cards as it failed do create a deck\n")
		t.FailNow()
	}
	if !strings.EqualFold(deck.DeckID, "") {
		t.Logf("Deck not properly initialized. Expected an empty DeckID\n")
		t.FailNow()
	}
	if deck.Shuffled {
		t.Logf("Deck not properly initialized. Expected an unshuffled deck\n")
		t.FailNow()
	}
	suits[0] = "S"
}

func TestNewDeckWithJokers(t *testing.T) {
	deck := NewDeckWithJokers(1)
	if deck == nil {
		t.Logf("Failed to create deck\n")
		t.FailNow()
	}
	if !deck.Success {
		t.Logf("Deck not properly initialized. Expected a success on a successful creation\n")
		t.FailNow()
	}
	if deck.Remaining != len(deck.cards) && len(deck.cards) != 54 {
		t.Logf("Deck not properly initialized. Expected amount of cards remaining and the length of the cards slice to be equavalent after a new deck is created.\n")
		t.FailNow()
	}
	if strings.EqualFold(deck.DeckID, "\n") {
		t.Logf("Deck not properly initialized. Expected a non empty DeckID\n")
		t.FailNow()
	}
	if deck.Shuffled {
		t.Logf("Deck not properly initialized. Expected an unshuffled deck\n")
		t.FailNow()
	}
	if !strings.EqualFold(deck.cards[53].Value, "JOKER\n") && !strings.EqualFold(deck.cards[53].Suit, "NONE") && !strings.EqualFold(deck.cards[52].Value, "JOKER\n") && !strings.EqualFold(deck.cards[52].Suit, "NONE") {
		t.Logf("Deck not properly initialized. Expected last two cards on an unshuffled deck to be JOKERS.\n")
		t.FailNow()
	}
}

func TestNewDeckWithJokerWithNegativeNumber(t *testing.T) {
	deck := NewDeckWithJokers(-1)
	if deck == nil {
		t.Logf("Failed to create deck\n")
		t.FailNow()
	}
	if deck.Success {
		t.Logf("Deck not properly initialized failed deck. Expected a false on a unsuccessful creation\n")
		t.FailNow()
	}
	if deck.Remaining > 0 {
		t.Logf("Deck not properly initialized. Expected zero cards as it failed do create a deck\n")
		t.FailNow()
	}
	if strings.EqualFold(deck.DeckID, "") {
		t.Logf("Deck not properly initialized. Expected a non empty DeckID\n")
		t.FailNow()
	}
	if deck.Shuffled {
		t.Logf("Deck not properly initialized. Expected an unshuffled deck\n")
		t.FailNow()
	}
}

func TestInjectionOfUnsupportedSuitWhileRunningNewDeckWithJoker(t *testing.T) {
	suits[0] = "F"
	deck := NewDeckWithJokers(1)
	if deck == nil {
		t.Logf("Failed to create deck\n")
		t.FailNow()
	}
	if deck.Success {
		t.Logf("Deck not properly initialized failed deck. Expected a false on a unsuccessful creation\n")
		t.FailNow()
	}
	if deck.Remaining > 0 {
		t.Logf("Deck not properly initialized. Expected zero cards as it failed do create a deck\n")
		t.FailNow()
	}
	if !strings.EqualFold(deck.DeckID, "") {
		t.Logf("Deck not properly initialized. Expected an empty DeckID\n")
		t.FailNow()
	}
	if deck.Shuffled {
		t.Logf("Deck not properly initialized. Expected an unshuffled deck\n")
		t.FailNow()
	}
	suits[0] = "S"
}

func TestShuffleDeck(t *testing.T) {
	deck := NewDeckWithJokers(1)
	t.Logf("Deck is being shuffled\n")
	deck = ShuffleDeck(deck)
	if !deck.Shuffled {
		t.Logf("Deck not properly shuffled. Expected an shuffled deck\n")
		t.FailNow()
	}
	if strings.EqualFold(deck.cards[53].Value, "JOKER\n") && strings.EqualFold(deck.cards[53].Suit, "NONE") && strings.EqualFold(deck.cards[52].Value, "JOKER") && strings.EqualFold(deck.cards[52].Suit, "NONE") {
		t.Logf("Deck not properly shuffled. Expected last two cards on an shuffled deck to not be JOKERS.\n")
		t.FailNow()
	}
}

func TestDrawWithValidNumber(t *testing.T) {
	deck := NewDeckWithJokers(1)
	remaining := deck.Remaining
	drawAmount := 2
	codeOfCard1, codeOfCard2 := deck.cards[0].Code, deck.cards[1].Code
	draw := deck.Draw(drawAmount)
	if deck.Remaining != remaining-drawAmount {
		t.Logf("Draw did not reduce the number of remaining cards\n")
		t.FailNow()
	}
	if len(draw.Cards) != drawAmount {
		t.Logf("The length of the drawn cards is not the same as the amount we drew.\n")
		t.FailNow()
	}
	if !deck.cards[0].drawn && !deck.cards[1].drawn {
		t.Logf("The cards that should have been drawn is not\n")
		t.FailNow()
	}

	if !strings.EqualFold(draw.Cards[0].Code, codeOfCard1) && !strings.EqualFold(draw.Cards[1].Code, codeOfCard2) {
		t.Logf("The draw did not have the expected cards from the top of the deck\n")
		t.FailNow()
	}

	if !draw.Success {
		t.Logf("The draw reports it was not successful\n")
		t.FailNow()
	}
	if draw.Remaining != 2 {
		t.Logf("The draw's Remaining and the deck's Remaining cards does not match.\n")
		t.FailNow()
	}
}

func TestDrawWithMoreThanRemainingNumberOfCards(t *testing.T) {
	deck := NewDeckWithJokers(1)
	deck.Remaining = 2
	remaining := deck.Remaining
	drawAmount := remaining + 2
	draw := deck.Draw(drawAmount)
	if deck.Remaining != 0 {
		t.Logf("Draw reduced the number of remaining cards by more than possible.\n")
		t.FailNow()
	}
	if len(draw.Cards) != remaining {
		t.Logf("The length of the drawn cards is not the same as the amount of remaining cards.\n")
		t.FailNow()
	}
	if !draw.Success {
		t.Logf("The draw reports it was not successful\n")
		t.FailNow()
	}
	if draw.Remaining != remaining {
		t.Logf("The draw's Remaining and the deck's Remaining cards does not match.\n")
		t.FailNow()
	}
}

func TestDrawWithNoMoreCardsRemaining(t *testing.T) {
	deck := NewDeckWithJokers(1)
	deck.Remaining = 0
	remaining := deck.Remaining
	drawAmount := remaining + 2
	draw := deck.Draw(drawAmount)
	if deck.Remaining != 0 {
		t.Logf("Draw reduced the number of remaining cards by more than possible.\n")
		t.FailNow()
	}
	if len(draw.Cards) != remaining {
		t.Logf("The length of the drawn cards is not the same as the amount of remaining cards.\n")
		t.FailNow()
	}
	if draw.Success {
		t.Logf("The draw reports it was successful but it should be unsuccessful.\n")
		t.FailNow()
	}
	if draw.Remaining != deck.Remaining {
		t.Logf("The draw's Remaining and the deck's Remaining cards does not match.\n")
		t.FailNow()
	}
}
func TestDrawWithInvalidNumber(t *testing.T) {
	deck := NewDeckWithJokers(1)
	remaining := deck.Remaining
	drawAmount := -1
	draw := deck.Draw(drawAmount)
	if deck.Remaining != remaining {
		t.Logf("Drawing an invalid amount did reduce the number of remaining cards\n")
		t.FailNow()
	}
	if len(draw.Cards) != 0 {
		t.Logf("The draw should have returned no cards for a invalid draw\n")
		t.FailNow()
	}
}

func TestDeckString(t *testing.T) {
	deck := NewDeck(1)
	id := fmt.Sprintf("DeckID: %s", deck.DeckID)
	remaining := fmt.Sprintf("Remaining: %d", deck.Remaining)
	success := fmt.Sprintf("Success: %t", deck.Success)
	shuffled := fmt.Sprintf("Shuffled: %t", deck.Shuffled)

	deckString := deck.String()

	if !strings.Contains(deckString, id) {
		t.Logf("The string needs to contain the DeckID\n")
		t.FailNow()
	}
	if !strings.Contains(deckString, success) {
		t.Logf("The string needs to contain a success string\n")
		t.FailNow()
	}
	if !strings.Contains(deckString, shuffled) {
		t.Logf("The string needs to contain a shuffled string\n")
		t.FailNow()
	}
	if !strings.Contains(deckString, remaining) {
		t.Logf("The string needs to contain a remaining string that contains the exact amount of cards remaining.\n")
		t.FailNow()
	}
}

func TestUnmarshal(t *testing.T) {
	deck := NewDeck(1)
	marshalDeck, err := deck.Marshal()
	if err != nil {
		t.Logf("There was an error marshaling the deck: %s\n", err.Error())
		t.FailNow()
	}
	udeck, err := UnmarshalDeck(marshalDeck)
	if err != nil {
		t.Logf("There was an error unmarshaling the deck: %s\n", err.Error())
		t.FailNow()
	}
	if deck.DeckID != udeck.DeckID {
		t.Logf("The DeckID's do not match\n")
		t.FailNow()
	}
	if deck.Remaining != udeck.Remaining {
		t.Logf("The Remaining cards do not match\n")
		t.FailNow()
	}
	if deck.Success != udeck.Success {
		t.Logf("The Success property do not match\n")
		t.FailNow()
	}
	if deck.Shuffled != udeck.Shuffled {
		t.Logf("The Shuffled property do not match\n")
		t.FailNow()
	}
}
