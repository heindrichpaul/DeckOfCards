package deckofcards

import (
	"fmt"
	"strings"

	"github.com/twinj/uuid"
)

//Pile is a type that implements the structure of a Draw.
type Pile struct {
	cards     Cards
	Remaining int
	PileID    string
}

//AddCardsToPile adds all the cards in the cards parameter that are present in the Draw to the pile.
func (z *Pile) AddCardsToPile(draw *Draw, cards Cards) {

	if draw != nil && draw.Success && len(draw.Cards) != 0 {
		if len(draw.Cards) >= len(cards) {

			for _, card := range cards {
				found := false
				for _, f := range draw.Cards {
					if f.Value == card.Value && f.Suit == card.Suit {
						found = true
					}
				}
				if found {
					fmt.Println(card.DeckID)
					z.cards = append(z.cards, card.cloneCard())
				}
			}
		}
	}

	z.Remaining = len(z.cards)
}

//NewPile creates a new Pile instance an returns a pointer to it.
func NewPile() *Pile {
	return &Pile{
		PileID:    uuid.NewV4().String(),
		cards:     make(Cards, 0),
		Remaining: 0,
	}
}

func (z *Pile) String() string {
	var printString []string
	printString = append(printString, fmt.Sprintf("PileID: %s", z.PileID))

	for _, pileCard := range z.cards {
		printString = append(printString, pileCard.String())
	}

	return strings.Join(printString, "\n")
}

//RetrieveCardsInPile returns a copy of all cards in the pile
func (z *Pile) RetrieveCardsInPile() (cards Cards) {
	for _, pileCard := range z.cards {
		cards = append(cards, pileCard.cloneCard())
	}

	return
}

//PickAmountOfCardsFromBottomOfPile returns a Draw of the amount selected from the bottom of the pile.
func (z *Pile) PickAmountOfCardsFromBottomOfPile(amount int) *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make(Cards, 0),
		Remaining: 0,
	}

	if amount <= 0 {
		return draw
	}

	if amount >= z.Remaining {
		amount = z.Remaining
	}

	sliceOfCards := z.cards[z.Remaining-amount:]
	draw.Cards = append(draw.Cards, sliceOfCards...)
	z.cards = z.cards[:z.Remaining-amount]
	z.Remaining = len(z.cards)

	draw.Remaining = len(draw.Cards)
	draw.Success = true

	return draw
}

//PickAmountOfCardsFromTopOfPile returns a Draw of the amount selected from the top of the pile.
func (z *Pile) PickAmountOfCardsFromTopOfPile(amount int) *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make(Cards, 0),
		Remaining: 0,
	}

	if amount <= 0 {
		return draw
	}

	if amount >= z.Remaining {
		amount = z.Remaining
	}

	sliceOfCards := z.cards[:amount]
	draw.Cards = append(draw.Cards, sliceOfCards...)
	z.cards = z.cards[amount:]
	z.Remaining = len(z.cards)

	draw.Remaining = len(draw.Cards)
	draw.Success = true

	return draw
}

//PickAllCardsFromPile returns all the cards in the pile as a Draw and clears the pile.
func (z *Pile) PickAllCardsFromPile() *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make(Cards, 0),
		Remaining: 0,
	}
	if len(z.cards) > 0 {
		draw.Success = true
		draw.Remaining = len(z.cards)
		draw.Cards = append(draw.Cards, z.cards...)
	}
	z.cards = make(Cards, 0)
	draw.Success = true
	return draw
}

//GetCardsFromPile returns the specified cards from the pile as a Draw.
func (z *Pile) GetCardsFromPile(cards Cards) *Draw {
	draw := &Draw{
		Success:   false,
		Cards:     make(Cards, 0),
		Remaining: 0,
	}
	if len(z.cards) > 0 && len(cards) <= len(z.cards) {
		var tempCards Cards
		for _, card := range cards {
			for _, pileCard := range z.cards {
				if strings.EqualFold(pileCard.Suit, card.Suit) && strings.EqualFold(pileCard.Value, card.Value) {
					tempCards = append(tempCards, pileCard)
				}
			}
		}
		for _, card := range cards {
			for i, pileCard := range z.cards {
				if strings.EqualFold(pileCard.Suit, card.Suit) && strings.EqualFold(pileCard.Value, card.Value) {
					z.cards = append(z.cards[:i], z.cards[i+1:]...)
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

//ShufflePile returns a shuffled pile
func ShufflePile(pile *Pile) *Pile {
	pile.cards = shuffle(pile.cards)
	return pile
}

//GetCardAtID returns the card at the id from the pile.
func (z *Pile) GetCardAtID(index int) (*Draw, error) {
	if index > len(z.cards)-1 || index < 0 {
		return nil, fmt.Errorf("Index out of bounds")
	}
	draw := &Draw{
		Success:   false,
		Cards:     make(Cards, 0),
		Remaining: 0,
	}
	draw.Cards = append(draw.Cards, z.cards[index])
	z.cards = append(z.cards[:index], z.cards[index+1:]...)
	z.Remaining = len(z.cards)
	draw.Remaining = len(draw.Cards)
	draw.Success = true
	return draw, nil
}
