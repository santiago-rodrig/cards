package main

import "fmt"

func main() {
	cards := newDeck()
	Must(cards.saveToFile("cards.txt"))
	cards, err := newDeckFromFile("cards.txt")
	Must(err)
	fmt.Println(cards)
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
