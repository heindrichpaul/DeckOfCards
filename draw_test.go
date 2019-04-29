package deckofcards

import (
	"fmt"
	"strings"
	"testing"
)

func TestMarshallingOfDraw(t *testing.T) {
	deck := NewDeck(1)
	if !deck.Success {
		t.Logf("Deck not successfully created.\n")
		t.FailNow()
	}
	draw := deck.Draw(2)
	if !draw.Success {
		t.Logf("Draw was not successful\n")
		t.FailNow()
	}

	marshallDraw, err := draw.Marshal()
	if err != nil {
		t.Logf("There was an error during marshaling the draw: %s\n", err.Error())
		t.FailNow()
	}

	newDraw, err := UnmarshalDraw(marshallDraw)
	if err != nil {
		t.Logf("There was an error during unmarshaling the draw: %s\n", err.Error())
		t.FailNow()
	}

	if newDraw.Success != draw.Success {
		t.Logf("The success property of the draws do not match.\n")
		t.FailNow()
	}

	if newDraw.Remaining != draw.Remaining {
		t.Logf("The Remaining property of the draws do not match.\n")
		t.FailNow()

	}

}

func TestDrawString(t *testing.T) {
	deck := NewDeckWithJokers(1)
	draw := deck.Draw(deck.Remaining)

	actualString := draw.String()
	expectedString := fmt.Sprintf("Success: %t\nRemaining: %d\n%s", draw.Success, draw.Remaining, draw.Cards.String())
	if !strings.EqualFold(actualString, expectedString) {
		t.Logf("expected:[%s] but received:[%s]\n", expectedString, actualString)
		t.FailNow()
	}
}

func TestAreAllCardsInThisDraw(t *testing.T) {
	deck := NewDeckWithJokers(1)
	draw := deck.Draw(deck.Remaining)
	if !draw.AreAllCardsInThisDraw(draw.Cards) {
		t.Logf("All cards in a draw should be present in itself")
		t.FailNow()
	}
}

func TestAreAllCardsInThisDrawNegativeTest(t *testing.T) {
	deck := NewDeck(1)
	draw := deck.Draw(deck.Remaining)
	card, err := newCard(deck.DeckID, "*", "*")
	if err != nil {
		t.Logf("Failed to create joker card")
		t.FailNow()
	}
	var cards Cards
	cards = append(cards, card)
	if draw.AreAllCardsInThisDraw(cards) {
		t.Logf("All cards from a deck without jockers should not contain a jocker")
		t.FailNow()
	}
}
