package deckOfCards

import (
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
