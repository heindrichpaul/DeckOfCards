package deckofcards

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

const TestDECKID = "TEST_ID"

func TestNewCard(t *testing.T) {

	var suits = [...]string{"S", "D", "C", "H"}
	for _, suit := range suits {
		for i := 2; i < 10; i++ {
			cardCreatorHelper(TestDECKID, suit, strconv.Itoa(i), t)
		}

		cardCreatorHelper(TestDECKID, suit, "0", t)
		cardCreatorHelper(TestDECKID, suit, "A", t)
		cardCreatorHelper(TestDECKID, suit, "K", t)
		cardCreatorHelper(TestDECKID, suit, "Q", t)
		cardCreatorHelper(TestDECKID, suit, "J", t)
	}

	cardCreatorHelper(TestDECKID, "*", "*", t)
}

func cardCreatorHelper(deckID, suit, value string, t *testing.T) {
	t.Log("Now running " + fmt.Sprintf("%s%s", value, suit) + ": " + time.Now().String())
	card, err := newCard(deckID, value, suit)
	if err != nil {
		t.Logf("Failed to create card for: %s%s\n", value, suit)
		t.FailNow()
	}
	if !strings.EqualFold(card.Code, fmt.Sprintf("%s%s", value, suit)) {
		t.Logf("Failed to verify card code for: %s%s expected: %s\n", value, suit, card.Code)
		t.FailNow()
	}
	CheckSuit(t, card, value, suit)
	CheckValue(t, card, value, suit)

	if card.drawn {
		t.Logf("Failed to verify card drawn flag for: %s%s expected: false but received: %t\n", value, suit, card.drawn)
		t.FailNow()
	}

	if !strings.EqualFold(card.DeckID, TestDECKID) {
		t.Logf("The DeckID is not correctly stored on the creation of a new card.\n")
		t.FailNow()

	}

	t.Log("Finished running " + fmt.Sprintf("%s%s", value, suit) + ": " + time.Now().String())
}

func CheckSuit(t *testing.T, card *Card, value, suit string) {
	switch suit {
	case "S":
		if !strings.EqualFold(card.Suit, "SPADES") {
			t.Logf("Failed to verify card suit for: %s%s expected: SPADES\n", value, suit)
			t.FailNow()
		}
	case "D":
		if !strings.EqualFold(card.Suit, "DIAMONDS") {
			t.Logf("Failed to verify card suit for: %s%s expected: DIAMONDS\n", value, suit)
			t.FailNow()
		}
	case "C":
		if !strings.EqualFold(card.Suit, "CLUBS") {
			t.Logf("Failed to verify card suit for: %s%s expected: CLUBS\n", value, suit)
			t.FailNow()
		}
	case "H":
		if !strings.EqualFold(card.Suit, "HEARTS") {
			t.Logf("Failed to verify card suit for: %s%s expected: HEARTS\n", value, suit)
			t.FailNow()
		}
	case "*":
		if !strings.EqualFold(card.Suit, "NONE") {
			t.Logf("Failed to verify card suit for: %s%s expected: NONE\n", value, suit)
			t.FailNow()
		}
	}
}

func CheckValue(t *testing.T, card *Card, value, suit string) {
	switch value {
	case "A":
		if !strings.EqualFold(card.Value, "ACE") {
			t.Logf("Failed to verify card code for: %s%s expected: ACE but received: %s\n", value, suit, card.Value)
			t.FailNow()
		}
	case "K":
		if !strings.EqualFold(card.Value, "KING") {
			t.Logf("Failed to verify card value for: %s%s expected: KING but received: %s\n", value, suit, card.Value)
			t.FailNow()
		}
	case "Q":
		if !strings.EqualFold(card.Value, "QUEEN") {
			t.Logf("Failed to verify card value for: %s%s expected: QUEEN but received: %s\n", value, suit, card.Value)
			t.FailNow()
		}
	case "J":
		if !strings.EqualFold(card.Value, "JACK") {
			t.Logf("Failed to verify card value for: %s%s expected: JACK but received: %s\n", value, suit, card.Value)
			t.FailNow()
		}
	case "0":
		if !strings.EqualFold(card.Value, "10") {
			t.Logf("Failed to verify card value for: %s%s expected: 10 but received: %s\n", value, suit, card.Value)
			t.FailNow()
		}
	case "*":
		if !strings.EqualFold(card.Value, "JOKER") {
			t.Logf("Failed to verify card value for: %s%s expected: JOKER but received %s\n", value, suit, card.Value)
			t.FailNow()
		}
	default:
		if !strings.EqualFold(card.Value, value) {
			t.Logf("Failed to verify card code for: %s%s expected: %s but received: %s\n", value, suit, value, card.Value)
			t.FailNow()
		}
	}
}
func TestNewCardWithInvalidSuit(t *testing.T) {
	suit := ""
	value := "0"
	expectedError := fmt.Sprintf("Card suit (%s), value (%s): invalid suit.", suit, value)
	card, err := newCard(TestDECKID, value, suit)
	if card == nil {
		if !strings.EqualFold(err.Error(), expectedError) {
			t.Logf("expected:[%s] but received:[%s]\n", expectedError, err.Error())
			t.FailNow()
		}
	}
}

func TestNewCardWithInvalidValue(t *testing.T) {
	suit := "S"
	value := "^"
	expectedError := fmt.Sprintf("Card suit (%s), value (%s): invalid value.", suit, value)
	card, err := newCard(TestDECKID, value, suit)
	if card == nil {
		if !strings.Contains(err.Error(), expectedError) {
			t.Logf("expected:[%s] but received:[%s]\n", expectedError, err.Error())
			t.FailNow()
		}
	}
}

func TestCardString(t *testing.T) {
	card, err := newCard(TestDECKID, "*", "*")
	if err != nil {
		t.Logf("Failed to create card: %s\n", err.Error())
		t.FailNow()
	}
	actualString := card.String()
	expectedString := fmt.Sprintf("%s - %s", card.Suit, card.Value)
	if !strings.EqualFold(actualString, expectedString) {
		t.Logf("expected:[%s] but received:[%s]\n", expectedString, actualString)
	}
}

func TestDraw(t *testing.T) {
	card, err := newCard(TestDECKID, "*", "*")
	if err != nil {
		t.Logf("Failed to create card: %s\n", err.Error())
		t.FailNow()
	}

	if card.drawn {
		t.Logf("Card not properly initialized. Expected drawn property to be false after creation.\n")
		t.FailNow()
	}

	drawnCard := card.draw()
	if !drawnCard.drawn {
		t.Logf("Card not properly drawn. Expected drawn property to be true after a draw.\n")
		t.FailNow()
	}
	if !strings.EqualFold(drawnCard.Code, card.Code) {
		t.Logf("expected:[%s] but received:[%s]\n", card.Code, drawnCard.Code)
		t.FailNow()
	}
	if !strings.EqualFold(drawnCard.Value, card.Value) {
		t.Logf("expected:[%s] but received:[%s]\n", card.Value, drawnCard.Value)
		t.FailNow()
	}
	if !strings.EqualFold(drawnCard.Suit, card.Suit) {
		t.Logf("expected:[%s] but received:[%s]\n", card.Suit, drawnCard.Suit)
		t.FailNow()
	}
}

func TestClone(t *testing.T) {
	card, err := newCard(TestDECKID, "*", "*")
	if err != nil {
		t.Logf("Failed to create card: %s\n", err.Error())
		t.FailNow()
	}
	clone := card.cloneCard()

	if !card.Equals(clone) {
		t.Logf("The two cards do not match after a clone.\n")
		t.FailNow()
	}

}

func TestEquals(t *testing.T) {
	card, err := newCard(TestDECKID, "*", "*")
	if err != nil {
		t.Logf("Failed to create card: %s\n", err.Error())
		t.FailNow()
	}
	clone := card.cloneCard()

	if !card.Equals(clone) {
		t.Logf("The two cards do not match after a clone.\n")
		t.FailNow()
	}

	clone.Code = "1"

	if card.Equals(clone) {
		t.Logf("The two cards do match after altering the clone.\n")
		t.FailNow()
	}

	clone = card.cloneCard()

	clone = card.cloneCard()
	clone.Suit = "3"
	if card.Equals(clone) {
		t.Logf("The two cards do match after altering the clone.\n")
		t.FailNow()
	}

	clone = card.cloneCard()
	clone.Value = "4"
	if card.Equals(clone) {
		t.Logf("The two cards do match after altering the clone.\n")
		t.FailNow()
	}

	clone = card.cloneCard()
	clone.drawn = !card.drawn
	if card.Equals(clone) {
		t.Logf("The two cards do match after altering the clone.\n")
		t.FailNow()
	}

}
