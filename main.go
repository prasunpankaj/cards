package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCard := deal(cards, 5) 
	hand.print()
	remainingCard.print()
	fmt.Println("*****")
}