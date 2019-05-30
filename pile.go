package deckofcards

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
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
		if len(draw.Cards) >= len(cards) && draw.AreAllCardsInThisDraw(cards) {
			z.cards = append(z.cards, cards...)
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
	printString = append(printString, z.cards.String())

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
			for i, pileCard := range z.cards {
				if card.Equals(pileCard) {
					tempCards = append(tempCards, pileCard)
					z.cards = append(z.cards[:i], z.cards[i+1:]...)
					z.Remaining--
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

//UnmarshalPile unmarshals a byte array into a pointer to a Pile for internal use.
func UnmarshalPile(data []byte) (*Pile, error) {
	var r Pile
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(data, &r)
	return &r, err
}

//Marshal marshals a pointer to a Pile into a byte array for transmission.
func (z *Pile) Marshal() ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(z)
}
