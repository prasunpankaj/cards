package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades","Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+ value) 
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i,card)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil{
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the Error and Entirely quit the program

		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")// Ace of Spades, Two of Spades, Three of Spades,

	return deck(s)
}