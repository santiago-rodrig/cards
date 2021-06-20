package main

import (
	"flag"
	"fmt"
	"github.com/santiago-rodrig/cards"
	"os"
	"strconv"
)

const notEnoughArgsMsg = "not enough arguments"
const badArgMsg = "invalid argument, must be an integer greater than zero"

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		printError(notEnoughArgsMsg)
		printUsage()
		os.Exit(1)
	}

	arg, err := strconv.Atoi(flag.Arg(0))

	if err != nil {
		printError(badArgMsg)
		printUsage()
		os.Exit(1)
	}

	deck := cards.NewDeck()

	for i := 0; i < arg; i++ {
		card, err := deck.PickRandomCard()
		if err != nil {
			break
		}
		fmt.Println(card)
	}
}

func printError(message string) {
	fmt.Println("ERROR:", message)
}

func printUsage() {
	fmt.Println("cards numberOfCards")
}
