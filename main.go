package main

import "fmt"

func main() {
	cards := newDeck()
	//cards.print()
	hand := cards.deal(5)
	fmt.Printf("hand: %d\n", len(hand))
	fmt.Printf("cards: %d\n", len(cards))
}
