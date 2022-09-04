package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){

	d := newDeck()

	if  len(d) != 51 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
		
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")
	
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")


	if len(loadedDeck) != 51{
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")


}