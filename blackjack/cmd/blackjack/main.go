package main

import (
	"github.com/gophercises/blackjack"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	g := blackjack.New(1)
	g.Init()
	g.Play()
}