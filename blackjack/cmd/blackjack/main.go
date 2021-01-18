package main

import (
	"fmt"
	"github.com/gophercises/blackjack"
)

func main() {
	g := blackjack.New(1)
	g.Init()
	fmt.Println(g)
	//init table inst
	//deal cards
	//for each player, do the moves, check victory
	//dealer turn
	//ask if quit/continue
}