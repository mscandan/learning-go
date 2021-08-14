package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // create new variable in string type and name it 'card' and assing "Ace of Spades" to it. Only value of string can be assingned to this variable
	card := "Ace of Spades" // this is the most common way to declare a variable. Variable's type will be set which value you put after the ':='
	// card = "Five of Diamonds" // this line is re-assings a value to a declared variable. When you re-assing a values to variables you can't use the ':='. You should use '='
	card2 := newCard() // create a variable named 'card2' and assign newCard function's return value to it.
	fmt.Println(card)
	fmt.Println(card2)
}

// creating a function and setting it's type of return value
func newCard() string {
	return "Five of Diamonds"
}
