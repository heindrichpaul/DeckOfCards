package deckofcards

import (
	"testing"
)

func TestMarshallingOfDraw(t *testing.T) {
	deck := NewDeck(1)
	if !deck.Success {
		t.Logf("Deck not successfully created.\n")
		t.FailNow()
	}
	draw := deck.Draw(2)
	if !draw.Success {
		t.Logf("Draw was not successful\n")
		t.FailNow()
	}

	marshallDraw, err := draw.Marshal()
	if err != nil {
		t.Logf("There was an error during marshaling the draw: %s\n", err.Error())
		t.FailNow()
	}

	newDraw, err := UnmarshalDraw(marshallDraw)
	if err != nil {
		t.Logf("There was an error during unmarshaling the draw: %s\n", err.Error())
		t.FailNow()
	}

	if newDraw.Success != draw.Success {
		t.Logf("The success property of the draws do not match.\n")
		t.FailNow()
	}

	if newDraw.Remaining != draw.Remaining {
		t.Logf("The Remaining property of the draws do not match.\n")
		t.FailNow()

	}

}
