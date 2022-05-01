package main


func main() {
	//cards := newDeck()
	// hand, remainingCard := deal(cards, 5) 
	// hand.print()
	// remainingCard.print()
	// fmt.Println("*****")
	// greeting := "Hi There!"
	// fmt.Println([]byte(greeting))

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards.txts")

	//cards := newDeckFromFile("my_cards.txts")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}