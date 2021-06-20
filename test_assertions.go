package main

import "testing"

func assertError(t *testing.T, errGot, errWant error) {
	t.Helper()
	if errGot != errWant {
		t.Errorf("expected %q to equal %q", errGot, errWant)
	}
}

func assertCardHasTheRightValueAndSuit(t *testing.T, card *Card, suit Suit, value Value) {
	t.Helper()
	if card.Suit != suit {
		t.Errorf("%q is not the same as %q", card.Suit, suit)
	} else if card.Value != value {
		t.Errorf("%q is not the same as %q", card.Value, value)
	}
}

func assertCardNotInDeck(t *testing.T, card *Card, deck Deck) {
	t.Helper()
	for _, deckCard := range deck {
		if deckCard.SameAs(card) {
			t.Fatalf("expected card %q not be found in the deck", card)
		}
	}
}

func assertEqualDecks(t *testing.T, deckA, deckB Deck) {
	t.Helper()
	if len(deckA) != len(deckB) {
		t.Errorf("decks do not have the same length, %d != %d", len(deckA), len(deckB))
	}
	for i, v := range deckA {
		if !v.SameAs(deckB[i]) {
			t.Fatalf("expected decks to be equal, but %q is note the same as %q", v, deckB[i])
		}
	}
}

func assertErrNoCardsInDeck(t *testing.T, errGot, errWant ErrNoCardsInDeck) {
	t.Helper()
	deckA := errGot.Deck
	deckB := errWant.Deck
	assertEqualDecks(t, deckA, deckB)
}
