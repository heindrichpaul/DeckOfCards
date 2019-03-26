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
		t.Logf("Failed to create deck\n")
		t.FailNow()
	}

	fmt.Printf("%s\n", (fmt.Sprintf("%s", deck)))
	pile := NewPile()
	draw := deck.Draw(6)
	if !draw.Success {
		t.Logf("Failed to draw from deck\n")
		t.FailNow()
	}
	fmt.Printf("%s\n", (fmt.Sprintf("%s", pile)))
	pile.AddCardsToPile(draw, draw.Cards)
	fmt.Printf("%s\n", (fmt.Sprintf("%s", pile)))

	found := false
	for _, pileCard := range pile.RetrieveCardsInPile() {
		for _, drawCard := range draw.Cards {
			found = pileCard.Equals(drawCard)
			if found {
				break
			}
		}
	}
	if !found {
		t.Logf("Not all cards in the draw is in the pile.\n")
		t.FailNow()
	}
}

func TestPile_PickAmountOfCardsFromBottomOfPile(t *testing.T) {
	deck := NewDeck(1)
	if !deck.Success {
		t.Logf("Failed to create deck\n")
		t.FailNow()
	}
	amountOfCards := 4

	fmt.Printf("%s\n", (fmt.Sprintf("%s", deck)))
	pile := NewPile()
	draw := deck.Draw(6)
	if !draw.Success {
		t.Logf("Failed to draw from deck\n")
		t.FailNow()
	}
	pile.AddCardsToPile(draw, draw.Cards)
	backupOfCardsInPile := pile.RetrieveCardsInPile()

	fmt.Printf("PickAmountOfCardsFromBottomOfPile\n%s\n", (fmt.Sprintf("%s", pile)))

	cardsFromPile := pile.PickAmountOfCardsFromBottomOfPile(amountOfCards)
	if cardsFromPile.Remaining != amountOfCards {
		t.Logf("Failed to draw (%d) from the deck\n", amountOfCards)
		t.FailNow()
	}
	amountOfCardsInPile := len(backupOfCardsInPile)
	found := false
	for _, pileCard := range backupOfCardsInPile[amountOfCardsInPile-cardsFromPile.Remaining:] {
		for _, drawCard := range cardsFromPile.Cards {
			found = pileCard.Equals(drawCard)
			if found {
				break
			}
		}
	}
	if !found {
		t.Logf("Not all cards in the draw is in the pile.\n")
		t.FailNow()
	}

	pile.AddCardsToPile(cardsFromPile, cardsFromPile.Cards)
	amountOfCards = pile.Remaining + 1
	backupOfCardsInPile = pile.RetrieveCardsInPile()
	amountOfCardsInPile = pile.Remaining
	cardsFromPile = pile.PickAmountOfCardsFromBottomOfPile(amountOfCards)
	if cardsFromPile.Remaining != amountOfCardsInPile {
		t.Logf("Failed to draw (%d) from the deck\n", amountOfCards)
		t.FailNow()
	}

	found = false
	for _, pileCard := range backupOfCardsInPile[amountOfCardsInPile-cardsFromPile.Remaining:] {
		for _, drawCard := range cardsFromPile.Cards {
			found = pileCard.Equals(drawCard)
			if found {
				break
			}
		}
	}
	if !found {
		t.Logf("Not all cards in the draw is in the pile.\n")
		t.FailNow()
	}

	cardsFromPile = pile.PickAmountOfCardsFromBottomOfPile(0)
	if cardsFromPile.Remaining != 0 {
		t.Logf("Failed to draw (%d) from the deck\n", amountOfCards)
		t.FailNow()
	}
}

func TestPile_PickAmountOfCardsFromTopOfPile(t *testing.T) {
	deck := NewDeck(1)
	if !deck.Success {
		t.Logf("Failed to create deck\n")
		t.FailNow()
	}
	amountOfCards := 4

	fmt.Printf("%s\n", (fmt.Sprintf("%s", deck)))
	pile := NewPile()
	draw := deck.Draw(6)
	if !draw.Success {
		t.Logf("Failed to draw from deck\n")
		t.FailNow()
	}
	pile.AddCardsToPile(draw, draw.Cards)
	backupOfCardsInPile := pile.RetrieveCardsInPile()

	fmt.Printf("PickAmountOfCardsFromTopOfPile\n%s\n", (fmt.Sprintf("%s", pile)))

	cardsFromPile := pile.PickAmountOfCardsFromTopOfPile(amountOfCards)
	if cardsFromPile.Remaining != amountOfCards {
		t.Logf("Failed to draw (%d) from the deck\n", amountOfCards)
		t.FailNow()
	}
	amountOfCardsInPile := len(backupOfCardsInPile)
	found := false
	for _, pileCard := range backupOfCardsInPile[:amountOfCardsInPile-cardsFromPile.Remaining] {
		for _, drawCard := range cardsFromPile.Cards {
			found = pileCard.Equals(drawCard)
			if found {
				break
			}
		}
	}
	if !found {
		t.Logf("Not all cards in the draw is in the pile.\n")
		t.FailNow()
	}

	pile.AddCardsToPile(cardsFromPile, cardsFromPile.Cards)
	amountOfCards = pile.Remaining + 1
	backupOfCardsInPile = pile.RetrieveCardsInPile()
	amountOfCardsInPile = pile.Remaining
	cardsFromPile = pile.PickAmountOfCardsFromTopOfPile(amountOfCards)
	if cardsFromPile.Remaining != amountOfCardsInPile {
		t.Logf("Failed to draw (%d) from the deck\n", amountOfCards)
		t.FailNow()
	}

	found = false
	for _, pileCard := range backupOfCardsInPile[:cardsFromPile.Remaining] {
		for _, drawCard := range cardsFromPile.Cards {
			found = pileCard.Equals(drawCard)
			if found {
				break
			}
		}
	}
	if !found {
		t.Logf("Not all cards in the draw is in the pile.\n")
		t.FailNow()
	}

	cardsFromPile = pile.PickAmountOfCardsFromTopOfPile(0)
	if cardsFromPile.Remaining != 0 {
		t.Logf("Failed to draw (%d) from the deck\n", amountOfCards)
		t.FailNow()
	}
}

func TestPile_PickAllCardsFromPile(t *testing.T) {
	deck := NewDeck(1)
	if !deck.Success {
		t.Logf("Failed to create deck\n")
		t.FailNow()
	}

	fmt.Printf("%s\n", (fmt.Sprintf("%s", deck)))
	pile := NewPile()
	draw := deck.Draw(6)
	if !draw.Success {
		t.Logf("Failed to draw from deck\n")
		t.FailNow()
	}
	pile.AddCardsToPile(draw, draw.Cards)
	backupOfCardsInPile := pile.RetrieveCardsInPile()

	fmt.Printf("PickAllCardsFromPile\n%s\n", (fmt.Sprintf("%s", pile)))

	cardsFromPile := pile.PickAllCardsFromPile()
	amountOfCardsInPile := len(backupOfCardsInPile)
	if cardsFromPile.Remaining != amountOfCardsInPile {
		t.Logf("Failed to draw all from the pile\n")
		t.FailNow()
	}

	found := false
	for _, pileCard := range backupOfCardsInPile {
		for _, drawCard := range cardsFromPile.Cards {
			found = pileCard.Equals(drawCard)
			if found {
				break
			}
		}
	}
	if !found {
		t.Logf("Not all cards in the draw is in the pile.\n")
		t.FailNow()
	}
}

func TestPile_GetCardsFromPile(t *testing.T) {
	amountOfCardsToDraw := 6
	deck := NewDeck(1)
	if !deck.Success {
		t.Logf("Failed to create deck\n")
		t.FailNow()
	}

	fmt.Printf("%s\n", (fmt.Sprintf("%s", deck)))
	pile := NewPile()
	draw := deck.Draw(amountOfCardsToDraw)
	if !draw.Success {
		t.Logf("Failed to draw from deck\n")
		t.FailNow()
	}

	cardsToRequestFromPile := make([]*Card, 0)
	if draw.Remaining >= amountOfCardsToDraw {
		cardsToRequestFromPile = append(cardsToRequestFromPile, draw.Cards[(amountOfCardsToDraw/1)-1])
		cardsToRequestFromPile = append(cardsToRequestFromPile, draw.Cards[amountOfCardsToDraw/2])
	}
	pile.AddCardsToPile(draw, draw.Cards)

	fmt.Printf("GetCardsFromPile\n%s\n", (fmt.Sprintf("%s", pile)))

	cardsFromPile := pile.GetCardsFromPile(cardsToRequestFromPile)
	if cardsFromPile.Remaining != 2 {
		t.Logf("Failed to draw all requested cards from the pile\n")
		t.FailNow()
	}

	found := false
	for _, pileCard := range cardsToRequestFromPile {
		for _, drawCard := range cardsFromPile.Cards {
			found = pileCard.Equals(drawCard)
			if found {
				break
			}
		}
	}
	if !found {
		t.Logf("Not all cards requested cards were in the draw.\n")
		t.FailNow()
	}
}
