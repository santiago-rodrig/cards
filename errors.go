package cards

import "fmt"

type ErrOutOfBoundsSuit Suit

func (err ErrOutOfBoundsSuit) Error() string {
	return fmt.Sprintf("%d is a suit that is out of bounds", err)
}

func NewErrOutOfBoundsSuit(suit Suit) ErrOutOfBoundsSuit {
	return ErrOutOfBoundsSuit(suit)
}

type ErrOutOfBoundsValue Value

func (err ErrOutOfBoundsValue) Error() string {
	return fmt.Sprintf("%d is a value that is out of bounds", err)
}

func NewErrOutOfBoundsValue(value Value) ErrOutOfBoundsValue {
	return ErrOutOfBoundsValue(value)
}

type ErrNoCardsInDeck struct {
	Deck Deck
}

func (err ErrNoCardsInDeck) Error() string {
	return "the deck has no cards to take from"
}

func NewErrNoCardsInDeck(deck Deck) ErrNoCardsInDeck {
	return ErrNoCardsInDeck{
		Deck: deck,
	}
}
