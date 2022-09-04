package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Creat a new type of 'deck'
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	cards.shuffle()
	return cards
}

func  deal(d deck, handSizes int) (deck, deck) {
	return d[:handSizes], d[handSizes:]

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func (d deck) toString() string {
	return strings.Join([]string(d), ",") 
}

func (d deck) shuffle() {
	// Generate a new seed for rand each time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		random := r.Intn(len(d) - 1)
		d[i], d[random] = d[random], d[i]
	}
}

func (d deck) saveToFile(filename string) error{

	strDeck := d.toString()
	data2Save := []byte(strDeck)
	return os.WriteFile(filename,data2Save, 0666 )
}

func newDeckFromFile(filename string) deck{

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	
	}
	strDeck := strings.Split(string(file),",")
	return deck(strDeck)

}