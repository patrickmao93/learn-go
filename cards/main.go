package main

func main() {
	deck := newDeck()
	deck.saveToFile("./deck")
	deck2 := newDeckFromFile("./deck")
	deck2 = deck2.shuffle()
	deck2.print()
}
