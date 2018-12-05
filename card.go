package deckofcards

import (
	"fmt"
	"regexp"
	"strings"
)

//Card is a type that implements the structure of a Card.
type Card struct {
	Image  string `json:"image"`
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
	values := regexp.MustCompile(`[2-9]|0|A|K|Q|J|\*`)
	suites := regexp.MustCompile(`S|D|C|H|\*`)

	if !strings.EqualFold(deckID, "") {

		card = &Card{
			DeckID: deckID,
			Code:   "",
			Image:  "",
			Value:  "",
			Suit:   "",
			drawn:  false,
		}

		if !suites.MatchString(suit) {
			return nil, &cardError{"invalid suit.", value, suit}
		}

		switch suit {
		case "S":
			card.Suit = "SPADES"
		case "D":
			card.Suit = "DIAMONDS"
		case "C":
			card.Suit = "CLUBS"
		case "H":
			card.Suit = "HEARTS"
		default:
			card.Suit = "NONE"
		}

		if !values.MatchString(value) {
			return nil, &cardError{"invalid value.", value, suit}
		}
		switch value {
		case "A":
			card.Value = "ACE"
		case "K":
			card.Value = "KING"
		case "Q":
			card.Value = "QUEEN"
		case "J":
			card.Value = "JACK"
		case "0":
			card.Value = "10"
		case "*":
			card.Value = "JOCKER"
		default:
			card.Value = value
		}

		card.Code = fmt.Sprintf("%s%s", value, suit)
		if !strings.EqualFold("*", value) && !strings.EqualFold("*", suit) {
			card.Image = fmt.Sprintf("https://deckofcardsapi.com/static/img/%s.png", card.Code)
		} else {
			card.Image = ""
		}
	}

	return
}

//String function serializes the Card struct into a representable string output.
func (z *Card) String() string {
	return fmt.Sprintf("DeckID: %s\n%s - %s", z.DeckID, z.Suit, z.Value)
}

func (z *Card) draw() *Card {

	z.drawn = true
	card := &Card{
		Code:  z.Code,
		Image: z.Image,
		Value: z.Value,
		Suit:  z.Suit,
		drawn: z.drawn,
	}

	return card
}

func (z *Card) cloneCard() *Card {
	card := &Card{
		Code:  z.Code,
		Image: z.Image,
		Value: z.Value,
		Suit:  z.Suit,
		drawn: z.drawn,
	}

	return card
}

func (z *cardError) Error() string {
	return fmt.Sprintf(`Card suit (%s), value (%s): %s`, z.suit, z.value, z.err)
}
