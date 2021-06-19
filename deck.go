package cards

import (
	"math/rand"
	"time"
)

type Deck []*Card

var TemplateDeck Deck

func init() {
	TemplateDeck = buildDeck()
	rand.Seed(time.Now().UnixNano())
}

func buildDeck() (d Deck) {
	for suit := range SuitsMap {
		for value := range ValuesMap {
			c, _ := NewCard(value, suit)
			d = append(d, c)
		}
	}

	return
}

func NewDeck() Deck {
	return TemplateDeck[:]
}

func (d *Deck) PickRandomCard() (*Card, error) {
	if len(*d) < 1 {
		return nil, NewErrNoCardsInDeck(*d)
	}
	position := rand.Intn(len(*d))
	card := (*d)[position]
	*d = append((*d)[:position], (*d)[position+1:]...)
	return card, nil
}
