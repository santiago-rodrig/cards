package main

import "fmt"

type Suit uint8

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
)

var SuitsMap = map[Suit]string{
	Spades:   "Spades",
	Diamonds: "Diamonds",
	Clubs:    "Clubs",
	Hearts:   "Hearts",
}

func (s Suit) String() string {
	return SuitsMap[s]
}

type Value uint8

const (
	Ace Value = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var ValuesMap = map[Value]string{
	Ace:   "Ace",
	Two:   "Two",
	Three: "Three",
	Four:  "Four",
	Five:  "Five",
	Six:   "Six",
	Seven: "Seven",
	Eight: "Eight",
	Nine:  "Nine",
	Ten:   "Ten",
	Jack:  "Jack",
	Queen: "Queen",
	King:  "King",
}

func (v Value) String() string {
	return ValuesMap[v]
}

type Card struct {
	Suit  Suit
	Value Value
}

func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.Value, c.Suit)
}

func (c *Card) SameAs(card *Card) bool {
	suitsMatch := c.Suit == card.Suit
	valuesMatch := c.Value == card.Value
	return suitsMatch && valuesMatch
}

func NewCard(value Value, suit Suit) (*Card, error) {
	if uint8(suit) > 3 {
		return nil, NewErrOutOfBoundsSuit(suit)
	} else if uint8(value) > 12 {
		return nil, NewErrOutOfBoundsValue(value)
	}

	return &Card{
		Suit:  suit,
		Value: value,
	}, nil
}
