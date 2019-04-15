package deckofcards

import (
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

	for _, card := range z.Cards {
		printString = append(printString, card.String())
	}

	return strings.Join(printString, "\n")

}
