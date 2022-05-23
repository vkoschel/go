package main

func main() {
	cards := newDeckFromFile("victor")
	cards.print()
	cards.shuffle()
	cards.print()
}
