package deckofcards

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/twinj/uuid"
)

var suits = [...]string{"S", "D", "C", "H"}

func UnmarshalDeck(data []byte) (*Deck, error) {
	var r Deck
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (z *Deck) Marshal() ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(z)
}

type Deck struct {
	Remaining int    `json:"remaining"`
	DeckID    string `json:"deckId"`
	Success   bool   `json:"success"`
	Shuffled  bool   `json:"shuffled"`
	cards     []*Card
}

func NewDeck(amount int) *Deck {
	deck, err := newDeck(amount, false)
	if err != nil {
		log.Printf("Failed to create the deck: %s", err.Error())
		deck = &Deck{
			DeckID:    "",
			Success:   false,
			Shuffled:  false,
			Remaining: 0,
			cards:     make([]*Card, 0),
		}

	}

	return deck
}

func NewDeckWithJockers(amount int) *Deck {
	deck, err := newDeck(amount, true)
	if err != nil {
		log.Printf("Failed to create the deck: %s", err.Error())
		deck = &Deck{
			DeckID:    "",
			Success:   false,
			Shuffled:  false,
			Remaining: 0,
			cards:     make([]*Card, 0),
		}
	}

	return deck
}

func newDeck(amount int, jockers bool) (deck *Deck, err error) {
	deck = &Deck{
		DeckID:    uuid.NewV4().String(),
		Success:   false,
		Shuffled:  false,
		Remaining: 0,
		cards:     make([]*Card, 0),
	}

	for deckNum := 0; deckNum < amount; deckNum++ {
		for _, suit := range suits {
			//ACE
			card, err := newCard(deck.DeckID, "A", suit)
			if err != nil {
				return nil, err
			}
			deck.cards = append(deck.cards, card)
			deck.Remaining++
			//NUMERICAL CARDS
			for i := 2; i < 10; i++ {
				card, err = newCard(deck.DeckID, strconv.Itoa(i), suit)
				if err == nil {
					deck.cards = append(deck.cards, card)
					deck.Remaining++
				}
			}
			//TEN
			card, err = newCard(deck.DeckID, "0", suit)
			if err == nil {
				deck.cards = append(deck.cards, card)
				deck.Remaining++
			}
			//JACK
			card, err = newCard(deck.DeckID, "J", suit)
			if err == nil {
				deck.cards = append(deck.cards, card)
				deck.Remaining++
			}
			//QUEEN
			card, err = newCard(deck.DeckID, "Q", suit)
			if err == nil {
				deck.cards = append(deck.cards, card)
				deck.Remaining++
			}
			//KING
			card, err = newCard(deck.DeckID, "K", suit)
			if err == nil {
				deck.cards = append(deck.cards, card)
				deck.Remaining++
			}
		}
		if jockers {
			card, err := newCard(deck.DeckID, "*", "*")
			if err == nil {
				deck.cards = append(deck.cards, card)
				deck.Remaining++
			}
			card, err = newCard(deck.DeckID, "*", "*")
			if err == nil {
				deck.cards = append(deck.cards, card)
				deck.Remaining++
			}
		}
	}
	if deck.Remaining == len(deck.cards) && deck.Remaining > 0 {
		deck.Success = true
		deck.Shuffled = false
	}
	return deck, nil
}

func ShuffleDeck(deck *Deck) *Deck {
	for i := 1; i < len(deck.cards); i++ {
		// Create a random int up to the number of cards
		r := rand.Intn(i + 1)

		// If the the current card doesn't match the random
		// int we generated then we'll switch them out
		if i != r {
			deck.cards[r], deck.cards[i] = deck.cards[i], deck.cards[r]
		}
	}
	for _, card := range deck.cards {
		card.drawn = false
	}
	deck.Shuffled = true
	return deck
}

func (z *Deck) Draw(amount int) (draw *Draw) {
	draw = &Draw{
		Cards:     make([]*Card, 0),
		Remaining: 0,
		Success:   false,
	}

	if z.Remaining == 0 {
		return
	}
	if amount <= 0 {
		return
	}

	var cards []*Card

	i := 0

	if amount > z.Remaining {
		amount = z.Remaining
	}

	for _, card := range z.cards {
		if !card.drawn {
			if i == amount {
				break
			}
			cards = append(cards, card.draw())
			z.Remaining--
			i++
		}
	}
	draw.Cards = cards
	draw.Success = true
	draw.Remaining = z.Remaining
	return
}

func (z *Deck) String() string {
	var printString []string
	printString = append(printString, fmt.Sprintf("DeckID: %s", z.DeckID))
	printString = append(printString, fmt.Sprintf("Success: %t", z.Success))
	printString = append(printString, fmt.Sprintf("Shuffled: %t", z.Shuffled))
	printString = append(printString, fmt.Sprintf("Remaining: %d", z.Remaining))

	for _, card := range z.cards {
		if !card.drawn {
			printString = append(printString, fmt.Sprintf("%s", card))
		}
	}

	return strings.Join(printString, "\n")
}
