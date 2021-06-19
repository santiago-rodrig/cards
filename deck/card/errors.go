package card

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
