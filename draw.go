package deckofcards

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

//UnmarshalDraw unmarshals a byte array into a pointer to a Draw for internal use.
func UnmarshalDraw(data []byte) (*Draw, error) {
	var r *Draw
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(data, &r)
	return r, err
}

//Marshal marshals a pointer to a Draw into a byte array for transmission.
func (z *Draw) Marshal() ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(z)
}

//Draw is a type that implements the structure of a Draw.
type Draw struct {
	Success   bool  `json:"success"`
	Cards     Cards `json:"cards"`
	Remaining int   `json:"remaining"`
}

//String function serializes the Draw struct into a representable string output.
func (z *Draw) String() string {
	var printString []string
	printString = append(printString, fmt.Sprintf("Success: %t", z.Success))
	printString = append(printString, fmt.Sprintf("Remaining: %d", z.Remaining))
	printString = append(printString, z.Cards.String())

	return strings.Join(printString, "\n")

}

//AreAllCardsInThisDraw returns true if all the cards in the cards slice are in the draw.
func (z *Draw) AreAllCardsInThisDraw(cards Cards) bool {
	for _, cardFromCards := range cards {
		for d, cardFromDraw := range z.Cards {
			if cardFromCards.Equals(cardFromDraw) {
				break
			}
			if d == len(z.Cards)-1 && !cardFromCards.Equals(cardFromDraw) {
				return false
			}
		}
	}
	return true
}
