package main

func main() {
	cards := newDeck()
	Must(cards.saveToFile("cards.txt"))
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
