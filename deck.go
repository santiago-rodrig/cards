package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() (cards deck) {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return
}

func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func (d *deck) deal(n uint) (ret deck) {
	ret = (*d)[:n]
	*d = (*d)[n:]
	return
}

func (d *deck) toString() (result string) {
	for i, card := range *d {
		if i != len(*d)+1 {
			result += card + "\n"
		} else {
			result += card
		}
	}
	return
}

func (d *deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) (deck, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	str := string(bs)
	strs := strings.Split(str, "\n")
	return deck(strs), nil
}
