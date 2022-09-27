// Create and manipulate a deck
package main

import "fmt"

func main() {
	deck := newDeck() //Create a new deck of cards
	fmt.Println("Initial deck:")
	deck.printDeck()
	hand, remainingDeck := deal(deck, 5)
	fmt.Println("Dealt hand")
	hand.printDeck()
	fmt.Println("Remaining deck")
	remainingDeck.printDeck()
}
