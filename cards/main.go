package main

import "fmt"


func main() {

	// var card string = "Ace of Spades"
	cards := [] string{newCard(), newCard()}
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	// fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}