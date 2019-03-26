package deckofcards

import (
	"fmt"
	"strings"

	"github.com/twinj/uuid"
)

//Pile is a type that implements the structure of a Draw.
type Pile struct {
	stack     []*pileObject
	Remaining int
	PileID    string
}

type pileObject struct {
	deckID string
	card   *Card
}

//String function serializes the Pile struct into a representable string output.
func (z *pileObject) String() string {
	var printString []string
	printString = append(printString, fmt.Sprintf("%s", z.card))

	return strings.Join(printString, "\n")
}

//AddCardsToPile adds all the cards in the cards parameter that are present in the Draw to the pile.
func (z *Pile) AddCardsToPile(draw *Draw, cards []*Card) {

	if draw != nil && draw.Success && len(draw.Cards) != 0 {
		if len(draw.Cards) >= len(cards) {

			for _, card := range cards {
				found := false
				for _, f := range draw.Cards {
					if f.Value == card.Value && f.Suit == card.Suit {
						//						draw.Cards = append(draw.Cards[:i], draw.Cards[i+1:]...)
						found = true
					}
				}
				if found {
					fmt.Println(card.DeckID)
					p := &pileObject{
						deckID: card.DeckID,
						card:   card,
					}
					z.stack = append(z.stack, p)
				}
			}
		}
	}

	z.Remaining = len(z.stack)
}

//NewPile creates a new Pile instance an returns a pointer to it.
func NewPile() *Pile {
	return &Pile{
		PileID:    uuid.NewV4().String(),
		stack:     make([]*pileObject, 0),
		Remaining: 0,
	}
}

func (z *Pile) String() string {
	var printString []string
	printString = append(printString, fmt.Sprintf("PileID: %s", z.PileID))

	for _, stackObject := range z.stack {
		printString = append(printString, stackObject.String())
	}

	return strings.Join(printString, "\n")
}

//RetrieveCardsInPile returns a copy of all cards in the pile
func (z *Pile) RetrieveCardsInPile() (cards []*Card) {
	for _, stackObject := range z.stack {
		cards = append(cards, stackObject.card.cloneCard())
	}

	return
}

//PickAmountOfCardsFromBottomOfPile returns a Draw of the amount selected from the bottom of the pile.
func (z *Pile) PickAmountOfCardsFromBottomOfPile(amount int) *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make([]*Card, 0),
		Remaining: 0,
	}

	if amount <= 0 {
		return draw
	}

	if amount >= z.Remaining {
		amount = z.Remaining
	}

	sliceOfStack := z.stack[z.Remaining-amount:]
	for _, stackObject := range sliceOfStack {
		draw.Cards = append(draw.Cards, stackObject.card)
	}
	z.stack = z.stack[:z.Remaining-amount]
	z.Remaining = len(z.stack)

	draw.Remaining = len(draw.Cards)
	draw.Success = true

	return draw
}

//PickAmountOfCardsFromTopOfPile returns a Draw of the amount selected from the top of the pile.
func (z *Pile) PickAmountOfCardsFromTopOfPile(amount int) *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make([]*Card, 0),
		Remaining: 0,
	}

	if amount <= 0 {
		return draw
	}

	if amount >= z.Remaining {
		amount = z.Remaining
	}

	sliceOfStack := z.stack[:amount]
	for _, stackObject := range sliceOfStack {
		draw.Cards = append(draw.Cards, stackObject.card)
	}
	z.stack = z.stack[amount:]
	z.Remaining = len(z.stack)

	draw.Remaining = len(draw.Cards)
	draw.Success = true

	return draw
}

//PickAllCardsFromPile returns all the cards in the pile as a Draw.
func (z *Pile) PickAllCardsFromPile() *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make([]*Card, 0),
		Remaining: 0,
	}
	if len(z.stack) > 0 {
		draw.Success = true
		draw.Remaining = len(z.stack)
		for _, stackObject := range z.stack {
			draw.Cards = append(draw.Cards, stackObject.card)
		}
	}
	z.stack = make([]*pileObject, 0)
	draw.Success = true
	return draw
}

//GetCardsFromPile returns the specified cards from the pile as a Draw.
func (z *Pile) GetCardsFromPile(cards []*Card) *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make([]*Card, 0),
		Remaining: 0,
	}
	if len(z.stack) > 0 && len(cards) <= len(z.stack) {
		var tempCards []*Card
		for _, card := range cards {
			for _, stackObject := range z.stack {
				if strings.EqualFold(stackObject.card.Suit, card.Suit) && strings.EqualFold(stackObject.card.Value, card.Value) {
					tempCards = append(tempCards, stackObject.card)
				}
			}
		}
		for _, card := range cards {
			for i, stackObject := range z.stack {
				if strings.EqualFold(stackObject.card.Suit, card.Suit) && strings.EqualFold(stackObject.card.Value, card.Value) {
					z.stack = append(z.stack[:i], z.stack[i+1:]...)
				}
			}
		}
		if len(cards) == len(tempCards) {
			draw.Success = true
			draw.Remaining = len(cards)
			draw.Cards = tempCards
		}

	}
	return draw
}
