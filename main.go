package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remaningDeck := deal(cards, 3)

	hand.print()
	fmt.Println("------------")
	remaningDeck.print()
}
