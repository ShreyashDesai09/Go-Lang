package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := newcard()
	fmt.Println("Card = ", card)

	cards := deck{"Ace of heart", card}
	cards = append(cards, "Six of Spade")

	fmt.Println(cards)

	cards.print()
}

func newcard() string {
	return "Five of Dimonds"
}
