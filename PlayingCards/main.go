// Create and manipulate a deck
package main

import "fmt"

func main() {
	deck := newDeck() //Create a new deck of cards
	deck.printDeck()
	//hand, remainingDeck := deal(deck, 5)
	//hand.printDeck()
	//remainingDeck.printDeck()
	fmt.Println(deck.toString())
}
