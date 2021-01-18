package blackjack

import (
	"fmt"
	"github.com/gophercises/deck"
)

//player controller - ai/human


//hit - deal new card
//stand - next player
const ochko = 21

type Player struct {
	hand []deck.Card
}

type Dealer struct {
	Player
}

type Game struct {
	deck []deck.Card
	players []Player
	dealer Dealer
}

func (p Player) Score() int {
	score := 0
	numAces := 0

	for _, card := range p.hand {
		switch {
		case card.Rank == deck.Ace:
			score += 1
			numAces += 1
		case card.Rank >= deck.Two && card.Rank <= deck.Ten:
			score += int(card.Rank)
		case card.Rank >= deck.Jack:
			score += 10
		}
	}

	//ace can be 1 or 11
	if numAces > 0 && score + 10 <= ochko {
		score += 10
	}

	return score
}

func (p Player) String() string {
	s := fmt.Sprintf("Player %s:\n", "Noname") //todo add player name
	for _, c := range p.hand {
		s = s + "\t" + fmt.Sprint(c) + "\n"
	}

	return s
}

func (d Dealer) String() string {
	s := fmt.Sprintf("Dealer:\n")
	//only show first dealer card
	if len(d.hand) > 0 {
		s = s + "\t" + fmt.Sprint(d.hand[1]) + "\n"
	}

	return s
}

func (g Game) String() string {
	var s string
	for _, p := range g.players {
		s += fmt.Sprintln(p)
	}
	//todo only show one card for dealer
	s += fmt.Sprintln(g.dealer)

	return s
}

func (g *Game) Play() {
	//iterate players and calc result
}

func (g *Game) Init() {
	g.deck = deck.New(deck.Multiply(2), deck.Shuffle(nil))

	initialCards := 2

	//reset player hands
	for i, _ := range g.players {
		g.players[i].hand = make([]deck.Card, initialCards, initialCards)
	}
	g.dealer.hand = make([]deck.Card, initialCards, initialCards)

	for _, p := range g.players {
		fmt.Println(len(p.hand))
	}
	//deal 2 cards for each player
	numPlayers := len(g.players)

	for i := 0; i<initialCards; i++ {
		for j, _ := range g.players {
			cardIdx := initialCards * i + j
			g.players[j].hand[i] = g.deck[cardIdx]
		}
		g.dealer.hand[i] = g.deck[initialCards * i + numPlayers]
	}

	//"remove" dealed cards from deck. Instead could add index for current card or something
	g.deck = g.deck[initialCards * (numPlayers + 1):]
}

type GameOption func(*Game)

func New(players int, opts ...GameOption) Game {
	var g Game

	g.players = make([]Player, players, players)

	for _, opt := range opts {
		opt(&g)
	}

	return g
}