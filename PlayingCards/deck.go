/*
Objectives:

	Create and return a list of playing cards
*/
package main

import "fmt"

type deck []string //A deck is a slice of strings

func newDeck() deck {
	cards := deck{} //Create an empty deck of cards
	//Declare card suits
	cardSuits := []string{"Diamonds", "Hearts", "Spades", "Clubs"}
	//Declare card values
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//Create a receiver function of type deck
func (d deck) printDeck() {
	//Reference the instance variable d of the created deck
	for _, card := range d {
		fmt.Println(card)
	}
}
