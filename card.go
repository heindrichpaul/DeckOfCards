package deckofcards

import (
	"fmt"
	"regexp"
	"strings"
)

//Card is a type that implements the structure of a Card.
type Card struct {
	Value  string `json:"value"`
	Suit   string `json:"suit"`
	Code   string `json:"code"`
	DeckID string `json:"deckId"`
	drawn  bool
}

type cardError struct {
	err   string
	value string
	suit  string
}

func newCard(deckID, value, suit string) (card *Card, err error) {

	if !strings.EqualFold(deckID, "") {

		card = &Card{
			DeckID: deckID,
			Code:   "",
			Value:  "",
			Suit:   "",
			drawn:  false,
		}

		card.Value, err = getValue(value)
		if err != nil {
			return card, err
		}
		card.Suit, err = getSuit(suit)
		if err != nil {
			return card, err
		}

		card.Code = fmt.Sprintf("%s%s", value, suit)
	}

	return
}

func getValue(value string) (result string, err error) {

	values := regexp.MustCompile(`[2-9]|0|A|K|Q|J|\*`)
	if !values.MatchString(value) {
		return "", &cardError{"invalid value.", value, ""}
	}

	switch value {
	case "A":
		result = "ACE"
	case "K":
		result = "KING"
	case "Q":
		result = "QUEEN"
	case "J":
		result = "JACK"
	case "0":
		result = "10"
	case "*":
		result = "JOKER"
	default:
		result = value
	}

	return
}

func getSuit(suit string) (result string, err error) {

	suites := regexp.MustCompile(`S|D|C|H|\*`)
	if !suites.MatchString(suit) {
		return "", &cardError{"invalid suit.", "", suit}
	}

	switch suit {
	case "S":
		result = "SPADES"
	case "D":
		result = "DIAMONDS"
	case "C":
		result = "CLUBS"
	case "H":
		result = "HEARTS"
	default:
		result = "NONE"
	}

	return
}

//String function serializes the Card struct into a representable string output.
func (z *Card) String() string {
	return fmt.Sprintf("DeckID: %s\n%s - %s", z.DeckID, z.Suit, z.Value)
}

func (z *Card) draw() *Card {

	z.drawn = true
	card := z.cloneCard()
	return card
}

func (z *Card) cloneCard() *Card {
	card := &Card{
		DeckID: z.DeckID,
		Code:   z.Code,
		Value:  z.Value,
		Suit:   z.Suit,
		drawn:  z.drawn,
	}

	return card
}

//Equals function compares two cards with each other.
func (z *Card) Equals(card *Card) bool {
	if z.Code == card.Code && z.Value == card.Value && z.Suit == card.Suit && z.drawn == card.drawn {
		return true
	}
	return false
}

func (z *cardError) Error() string {
	return fmt.Sprintf(`Card suit (%s), value (%s): %s`, z.suit, z.value, z.err)
}
