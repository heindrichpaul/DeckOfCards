package deckofcards

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewPile(t *testing.T) {
	pile := NewPile()
	if pile == nil {
		t.Logf("The pile was not successfully created.\n")
		t.FailNow()
	}
	if strings.EqualFold(pile.PileID, "") {
		t.Logf("The PileID was empty.\n")
		t.FailNow()
	}
}

func TestAddCardsToPile(t *testing.T) {
	deck := NewDeck(1)
	if !deck.Success {
	}

	fmt.Printf("%s\n", (fmt.Sprintf("%s", deck)))
	pile := NewPile()
	draw := deck.Draw(6)
	if !draw.Success {
	}
	fmt.Printf("%s\n", (fmt.Sprintf("%s", pile)))
	pile.AddCardsToPile(draw, draw.Cards)
	fmt.Printf("%s\n", (fmt.Sprintf("%s", pile)))
}
