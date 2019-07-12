package main

func main()  {
	cards := deck{"diamonds", newCard()}
	cards = append(cards, "spades")

	cards.print()
}

func newCard() string {
	return "spades"
}