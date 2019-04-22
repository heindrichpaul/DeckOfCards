package deckofcards

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/twinj/uuid"
)

var suits = [...]string{"S", "D", "C", "H"}

//UnmarshalDeck unmarshals a byte array into a pointer to a Deck for internal use.
func UnmarshalDeck(data []byte) (*Deck, error) {
	var r Deck
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(data, &r)
	return &r, err
}

//Marshal marshals a pointer to a Deck into a byte array for transmission.
func (z *Deck) Marshal() ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(z)
}

//Deck is a type that implements the structure of a Deck.
type Deck struct {
	Remaining int    `json:"remaining"`
	DeckID    string `json:"deckId"`
	Success   bool   `json:"success"`
	Shuffled  bool   `json:"shuffled"`
	cards     Cards
}

/*NewDeck creates an unshuffled amount of decks requested by the parameter (amount).
These decks do not contain jokers.*/
func NewDeck(amount int) *Deck {
	deck, err := newDeck(amount, false)
	if err != nil {
		log.Printf("Failed to create the deck: %s", err.Error())
		deck = &Deck{
			DeckID:    "",
			Success:   false,
			Shuffled:  false,
			Remaining: 0,
			cards:     make(Cards, 0),
		}

	}

	return deck
}

/*NewDeckWithJokers creates an unshuffled amount of decks requested by the parameter (amount).
These decks do contain jokers.*/
func NewDeckWithJokers(amount int) *Deck {
	deck, err := newDeck(amount, true)
	if err != nil {
		log.Printf("Failed to create the deck: %s", err.Error())
		deck = &Deck{
			DeckID:    "",
			Success:   false,
			Shuffled:  false,
			Remaining: 0,
			cards:     make(Cards, 0),
		}
	}

	return deck
}

func newDeck(amount int, jokers bool) (deck *Deck, err error) {
	deck = &Deck{
		DeckID:    uuid.NewV4().String(),
		Success:   false,
		Shuffled:  false,
		Remaining: 0,
		cards:     make(Cards, 0),
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
		if jokers {
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

//ShuffleDeck shuffles the deck that has been passed as a parameter.
func ShuffleDeck(deck *Deck) *Deck {
	deck.cards = shuffle(deck.cards)
	for _, card := range deck.cards {
		card.drawn = false
	}
	deck.Shuffled = true
	return deck
}

//Draw draws the amount of requested cards from the current deck.
func (z *Deck) Draw(amount int) (draw *Draw) {
	draw = &Draw{
		Cards:     make(Cards, 0),
		Remaining: 0,
		Success:   false,
	}

	if z.Remaining == 0 {
		return
	}
	if amount <= 0 {
		return
	}

	var cards Cards

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

//String function serializes the Deck struct into a representable string output.
func (z *Deck) String() string {
	var printString []string
	printString = append(printString, fmt.Sprintf("DeckID: %s", z.DeckID))
	printString = append(printString, fmt.Sprintf("Success: %t", z.Success))
	printString = append(printString, fmt.Sprintf("Shuffled: %t", z.Shuffled))
	printString = append(printString, fmt.Sprintf("Remaining: %d", z.Remaining))

	for _, card := range z.cards {
		if !card.drawn {
			printString = append(printString, card.String())
		}
	}

	return strings.Join(printString, "\n")
}
