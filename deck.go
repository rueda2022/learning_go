package main

import "fmt"

//Create a new type of 'deck' wich is a slice of strings

type deck []string

func new_deck() deck {

	cards := deck{}

	pintas := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	numbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, pinta := range pintas {
		for _, number := range numbers {
			cards = append(cards, number+" of "+pinta)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)

	}
}
