package cards

import "testing"

func TestNewCard(t *testing.T) {
	t.Run("proper value and suit for a cards", func(t *testing.T) {
		t.Run("creates a cards successfully", func(t *testing.T) {
			suit := Clubs
			value := Ace
			got, _ := NewCard(value, suit)
			assertCardHasTheRightValueAndSuit(t, got, suit, value)
		})
	})

	t.Run("out of bounds suit", func(t *testing.T) {
		t.Run("it returns nil and an error", func(t *testing.T) {
			suit := Suit(22)
			value := Ace
			_, err := NewCard(value, suit)
			assertError(t, err, NewErrOutOfBoundsSuit(suit))
		})
	})

	t.Run("out of bounds value", func(t *testing.T) {
		t.Run("it returns nil and an error", func(t *testing.T) {
			suit := Hearts
			value := Value(34)
			_, err := NewCard(value, suit)
			assertError(t, err, NewErrOutOfBoundsValue(value))
		})
	})
}

func TestCard(t *testing.T) {
	t.Run("SameAs", func(t *testing.T) {
		t.Run("same cards", func(t *testing.T) {
			cardA, _ := NewCard(Ace, Clubs)
			cardB, _ := NewCard(Ace, Clubs)

			t.Run("returns true", func(t *testing.T) {
				same := cardA.SameAs(cardB)

				if !same {
					t.Errorf("%q should be the same as %q", cardA, cardB)
				}
			})
		})

		t.Run("different cards", func(t *testing.T) {
			cardA, _ := NewCard(Ace, Clubs)
			cardB, _ := NewCard(Four, Hearts)

			t.Run("returns false", func(t *testing.T) {
				same := cardA.SameAs(cardB)

				if same {
					t.Errorf("%q should not be the same as %q", cardA, cardB)
				}
			})
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run(`returns "ValueHere of SuitHere"`, func(t *testing.T) {
			card, _ := NewCard(King, Spades)
			got := card.String()
			want := "King of Spades"

			if got != want {
				t.Errorf("%q is not equal to %q", got, want)
			}
		})
	})
}
