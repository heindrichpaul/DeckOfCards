package deckOfCards

import (
	"fmt"
	"strings"

	"github.com/twinj/uuid"
)

type Pile struct {
	stack  []*pileObject
	PileID string
}

type pileObject struct {
	deckID string
	card   *Card
}

func (z *pileObject) String() string {
	var printString []string
	printString = append(printString, fmt.Sprintf("DeckID: %s", z.deckID))
	printString = append(printString, fmt.Sprintf("%s", z.card))

	return strings.Join(printString, "\n")
}

func (z *Pile) AddCardsToPile(draw *Draw, cards []*Card) {

	if draw != nil && draw.Success && len(draw.Cards) != 0 {
		if len(draw.Cards) >= len(cards) {

			for _, card := range cards {
				found := false
				for i, f := range draw.Cards {
					if f.Value == card.Value && f.Suit == card.Suit {
						draw.Cards = append(draw.Cards[:i], draw.Cards[i+1:]...)
						found = true
					}
				}
				if found {
					p := &pileObject{
						deckID: card.DeckID,
						card:   card,
					}
					z.stack = append(z.stack, p)
				}
			}
		}
	} else {
	}
}

func NewPile() *Pile {
	return &Pile{
		PileID: uuid.NewV4().String(),
		stack:  make([]*pileObject, 0),
	}
}

func (z *Pile) ListCardsInPile() string {
	var printString []string
	printString = append(printString, fmt.Sprintf("PileID: %s", z.PileID))

	for _, stackObject := range z.stack {
		printString = append(printString, fmt.Sprintf("%s", stackObject))
	}

	return strings.Join(printString, "\n")
}

func (z *Pile) RetrieveCardsInPile() (cards []*Card) {
	for _, stackObject := range z.stack {
		cards = append(cards, stackObject.card)
	}

	return
}

func (z *Pile) PickAmountOfCardsFromBottomOfPile(amount int) *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make([]*Card, 0),
		Remaining: 0,
	}

	if amount <= 0 {
		return draw
	}

	if amount >= len(z.stack) {
		amount = len(z.stack)
	}

	for i, stackObject := range z.stack[len(z.stack)-amount-1:] {
		draw.Cards = append(draw.Cards, stackObject.card)
		z.stack = append(z.stack[:i], z.stack[i+1:]...)
	}

	draw.Remaining = len(draw.Cards)

	return draw
}

func (z *Pile) PickAmountOfCardsFromTopOfPile(amount int) *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make([]*Card, 0),
		Remaining: 0,
	}

	if amount <= 0 {
		return draw
	}

	if amount >= len(z.stack) {
		amount = len(z.stack)
	}

	for i, stackObject := range z.stack {
		draw.Cards = append(draw.Cards, stackObject.card)
		z.stack = append(z.stack[:i], z.stack[i+1:]...)
	}

	draw.Remaining = len(draw.Cards)

	return draw
}

func (z *Pile) PickAllCardsFromPile() *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make([]*Card, 0),
		Remaining: 0,
	}
	return
}

func (z *Pile) GetCardsFromPile(cards []*Card) *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make([]*Card, 0),
		Remaining: 0,
	}
	return nil
}
