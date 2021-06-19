package deck

import "github.com/santiago-rodrig/cards/deck/card"

type Deck []*card.Card

var TemplateDeck Deck

func init() {
	TemplateDeck = buildDeck()
}

func buildDeck() (d Deck) {
	for suit := range card.SuitsMap {
		for value := range card.ValuesMap {
			c, _ := card.NewCard(value, suit)
			d = append(d, c)
		}
	}

	return
}

func NewDeck() Deck {
	return TemplateDeck[:]
}
