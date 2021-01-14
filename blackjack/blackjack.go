package blackjack

import (
	"fmt"
	"github.com/gophercises/deck"
)

//player controller - ai/human
//display cards and calc score

//hit - deal new card
//stand - next player

type Player struct {
	hand []deck.Card
}

type Game struct {
	deck []deck.Card
	players []Player
	dealer Player
}

func (p Player) Score() int {
	//sum up scores based on rank
	//separately add Ace value
	return 0
}

func (p Player) String() string {
	s := fmt.Sprintf("Player %s: ", "empty") //todo add player name
	for _, c := range p.hand {
		s = s + fmt.Sprint(c) + " "
	}

	return s
}

func (g * Game) Play() {
	//iterate players and calc result
}

func (g *Game) Init() {
	//deal 2 cards for each player
	initialCards := 2
	numPlayers := len(g.players)

	for i := 0; i<initialCards; i++ {
		for j, p := range g.players {
			cardIdx := initialCards * i + j
			p.hand = append(p.hand, g.deck[cardIdx])
		}
		g.dealer.hand = append(g.dealer.hand, g.deck[initialCards * i + numPlayers])
	}

	//"remove" dealed cards from deck. Instead could add index for current card or something
	g.deck = g.deck[initialCards * (numPlayers + 1):]
}

type GameOption func(*Game)

func New(players int, opts ...GameOption) Game {
	var g Game

	g.deck = deck.New(deck.Multiply(2), deck.Shuffle(nil))

	g.players = make([]Player, players, players)

	for _, opt := range opts {
		opt(&g)
	}

	return g
}