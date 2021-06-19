package deck

import "testing"

func TestNewDeck(t *testing.T) {
	t.Run("creates a deck of 52 cards", func(t *testing.T) {
		got := NewDeck()
		want := TemplateDeck
		assertEqualDecks(t, got, want)
	})
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
