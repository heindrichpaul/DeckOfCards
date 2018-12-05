package deckOfCards

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func UnmarshalDraw(data []byte) (*Draw, error) {
	var r *Draw
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Draw) Marshal() ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(r)
}

type Draw struct {
	Success   bool    `json:"success"`
	Cards     []*Card `json:"cards"`
	Remaining int     `json:"remaining"`
}

func (z *Draw) String() string {
	var printString []string
	/*printString = append(printString, fmt.Sprintf("DeckID: %s", z.DeckID))
	printString = append(printString, fmt.Sprintf("Success: %t", z.Success))
	printString = append(printString, fmt.Sprintf("Shuffled: %t", z.Shuffled))
	printString = append(printString, fmt.Sprintf("Remaining: %d", z.Remaining))
	*/
	for _, card := range z.Cards {
		if !card.drawn {
			printString = append(printString, fmt.Sprintf("%s", card))
		}
	}

	return strings.Join(printString, "\n")

}
