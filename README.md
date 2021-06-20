# cards

Go program that **prints random cards to STDOUT**.

## Installation

```bash
go install github.com/santiago-rodrig/cards/cmd
```

**Alternatively**, you can also download the binary
that comes with
[the release]()
and place it
somewhere in your path.

To just get **the library** to use it in your Go
programs you do as follows.

```bash
go get github.com/santiago-rodrig/cards
```

And now you can use it in your code like this.

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/santiago-rodrig/cards"
)

func main() {
	deck := cards.NewDeck()
	card, err := deck.PickRandomCard()

	if err != nil {
		fmt.Println(err) // no more cards
		os.Exit(1)
    }
    
    fmt.Println(card) // Ace of Spades or some other card
}
```

## License

You can read the licensing terms behind this code
by referring to the
[license file in this repo](./LICENSE).