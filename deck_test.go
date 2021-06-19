package cards

import "testing"

func TestNewDeck(t *testing.T) {
	t.Run("creates a deck of 52 cards", func(t *testing.T) {
		got := NewDeck()
		want := TemplateDeck
		assertEqualDecks(t, got, want)
	})
}

func TestDeck(t *testing.T) {
	deck := NewDeck()

	t.Run("PickRandomCard", func(t *testing.T) {
		t.Run("decreases the length of the deck by one", func(t *testing.T) {
			before := len(deck)
			deck.PickRandomCard()
			after := len(deck)

			if after != before-1 {
				t.Errorf("expected the length of the deck to be decreased by one, %d != %d", before, after)
			}
		})

		t.Run("takes out the returned card from the deck", func(t *testing.T) {
			card, _ := deck.PickRandomCard()
			assertCardNotInDeck(t, card, deck)
		})

		t.Run("returns an error if there are no more cards in the deck", func(t *testing.T) {
			deck = Deck{}
			_, err := deck.PickRandomCard()
			assertErrNoCardsInDeck(t, err.(ErrNoCardsInDeck), NewErrNoCardsInDeck(deck))
		})
	})
}
