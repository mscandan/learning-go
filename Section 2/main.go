package main

func main() {
	// // var card string = "Ace of Spades" // create new variable in string type and name it 'card' and assing "Ace of Spades" to it. Only value of string can be assingned to this variable
	// card := "Ace of Spades" // this is the most common way to declare a variable. Variable's type will be set which value you put after the ':='
	// // card = "Five of Diamonds" // this line is re-assings a value to a declared variable. When you re-assing a values to variables you can't use the ':='. You should use '='
	// card2 := newCard() // create a variable named 'card2' and assign newCard function's return value to it.
	// fmt.Println(card)
	// fmt.Println(card2)

	// // declare a slice
	// cards := []string{"Ace of Spades", newCard()}
	// fmt.Println(cards)
	// cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	// // for loop
	// for i, singleCard := range cards {
	// 	fmt.Println(i, singleCard)
	// }

	// newDeck := deck{"Ace of Spades", newCard()}

	// for i, cardInDeck := range newDeck {
	// 	fmt.Println(i, cardInDeck)
	// }

	// newDeck.print()

	newDeck := newDeck()
	// newDeck.print()
	// hand, remainingCards := deal(newDeck, 5)
	newDeck.saveToFile("cards")
	// newDeckFromDisk := newDeckFromFile("cards")
	// fmt.Println(newDeckFromDisk)
	newDeck.shuffle()
	newDeck.print()
}

// creating a function and setting it's type of return value
func newCard() string {
	return "Five of Diamonds"
}
