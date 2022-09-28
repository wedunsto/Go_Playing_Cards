// Create and manipulate a deck
package main

func main() {
	deck := newDeck() //Create a new deck of cards
	hand, remainingDeck := deal(deck, 5)
	hand.saveToFile("hand.txt")
	remainingDeck.saveToFile("remaining.txt")
	deck = loadFromFile("hand.txt")
	deck.shuffle()
	deck.printDeck()
}
