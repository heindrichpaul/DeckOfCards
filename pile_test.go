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

	pile := NewPile()
	draw := deck.Draw(6)
	if !draw.Success {
		t.Logf("Failed to draw from deck\n")
		t.FailNow()
	}
	pile.AddCardsToPile(draw, draw.Cards)
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

	pile := NewPile()
	draw := deck.Draw(6)
	if !draw.Success {
		t.Logf("Failed to draw from deck\n")
		t.FailNow()
	}
	pile.AddCardsToPile(draw, draw.Cards)
	backupOfCardsInPile := pile.RetrieveCardsInPile()

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

	pile := NewPile()
	draw := deck.Draw(6)
	if !draw.Success {
		t.Logf("Failed to draw from deck\n")
		t.FailNow()
	}
	pile.AddCardsToPile(draw, draw.Cards)
	backupOfCardsInPile := pile.RetrieveCardsInPile()

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

	pile := NewPile()
	draw := deck.Draw(6)
	if !draw.Success {
		t.Logf("Failed to draw from deck\n")
		t.FailNow()
	}
	pile.AddCardsToPile(draw, draw.Cards)
	backupOfCardsInPile := pile.RetrieveCardsInPile()

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

	pile := NewPile()
	draw := deck.Draw(amountOfCardsToDraw)
	if !draw.Success {
		t.Logf("Failed to draw from deck\n")
		t.FailNow()
	}

	cardsToRequestFromPile := make(Cards, 0)
	if draw.Remaining >= amountOfCardsToDraw {
		cardsToRequestFromPile = append(cardsToRequestFromPile, draw.Cards[(amountOfCardsToDraw)-1])
		cardsToRequestFromPile = append(cardsToRequestFromPile, draw.Cards[amountOfCardsToDraw/2])
	}
	pile.AddCardsToPile(draw, draw.Cards)

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

func TestShufflePile(t *testing.T) {
	deck := NewDeckWithJokers(1)
	draw := deck.Draw(deck.Remaining)
	pile := NewPile()
	pile.AddCardsToPile(draw, draw.Cards)
	pile = ShufflePile(pile)
	if strings.EqualFold(pile.cards[53].Value, "JOKER\n") && strings.EqualFold(pile.cards[53].Suit, "NONE") && strings.EqualFold(pile.cards[52].Value, "JOKER") && strings.EqualFold(pile.cards[52].Suit, "NONE") {
		t.Logf("Pile not properly shuffled. Expected last two cards on an shuffled pile to not be JOKERS.\n")
		t.FailNow()
	}
}

func TestGetCardAtID(t *testing.T) {
	deck := NewDeckWithJokers(1)
	draw := deck.Draw(deck.Remaining)
	pile := NewPile()
	pile.AddCardsToPile(draw, draw.Cards)
	AmountOfCards := pile.Remaining
	draw, err := pile.GetCardAtID(53)
	if err != nil {
		t.Logf("Could not retrieve a card from the pile.\n")
		t.FailNow()
	}
	if draw.Remaining != 1 {
		t.Logf("The method drew the wrong amount of cards\n")
		t.FailNow()
	}
	if !strings.EqualFold(draw.Cards[0].Value, "JOKER\n") && !strings.EqualFold(draw.Cards[0].Suit, "NONE") {
		t.Logf("Pile not properly shuffled. Expected card to be a JOKERS.\n")
		t.FailNow()
	}
	if AmountOfCards == pile.Remaining {
		t.Logf("Did not remove the card from the pile which was returned.\n")
		t.FailNow()
	}
	_, err = pile.GetCardAtID(54)
	if err == nil {
		t.Logf("Did retrieve a card from the pile for a faulty id.\n")
		t.FailNow()
	}
}

func TestPileString(t *testing.T) {
	deck := NewDeckWithJokers(1)
	draw := deck.Draw(deck.Remaining)
	pile := NewPile()
	pile.AddCardsToPile(draw, draw.Cards)

	actualString := pile.String()
	expectedString := fmt.Sprintf("PileID: %s\n%s", pile.PileID, pile.cards.String())
	if !strings.EqualFold(actualString, expectedString) {
		t.Logf("expected:[%s] but received:[%s]\n", expectedString, actualString)
		t.FailNow()
	}
}
