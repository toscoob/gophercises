package blackjack

import (
	"bufio"
	"fmt"
	"github.com/gophercises/deck"
	"os"
)

// player controller - ai/human

const ochko = 21

type Player struct {
	hand []deck.Card
	name string
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

	// ace can be 1 or 11
	if numAces > 0 && score + 10 <= ochko {
		score += 10
	}

	return score
}

func (p Player) String() string {
	s := fmt.Sprintf("%s:\n", p.name)
	for _, c := range p.hand {
		s = s + "\t" + fmt.Sprint(c) + "\n"
	}

	return s
}

func (d Dealer) String() string {
	s := fmt.Sprintf("%s:\n", d.Player.name)
	// only show first dealer card
	if len(d.hand) > 0 {
		s = s + "\t" + fmt.Sprint(d.hand[1]) + "\n"
	}

	return s
}

func (g Game) String() string {
	var s string
	for _, p := range g.players {
		s += fmt.Sprint(p)
	}
	// only show one card for dealer
	s += fmt.Sprint(g.dealer)

	return s
}

func dealCard(h []deck.Card, d []deck.Card) (hand []deck.Card, deck []deck.Card) {
	hand = h
	deck = d

	if deck == nil || len(deck) == 0 {
		// todo handle it in some way
		return
	}

	hand = append(hand, deck[0])
	deck = deck[1:]

	return
}

func (g *Game) playerTurn(p *Player) {
	for {
		score := p.Score()
		fmt.Println(p, score)
		if score >= ochko {
			break
		}
		fmt.Println("press 'h' to hit or 's' to stand")

		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
			break
		}

		switch char {
		case 'H':
		case 'h':
			fmt.Println("Hit chosen")
			p.hand, g.deck = dealCard(p.hand, g.deck)
			break
		case 'S':
		case 's':
			fmt.Println("Stand chosen")
			return
		}
	}
}

func (g *Game) Play() {
	hScore := 0
	hPlayer := 0

	for i, p := range g.players {
		fmt.Printf("%s turn\n", p.name)
		g.playerTurn(&p)

		score := p.Score()
		if score > hScore && score <= ochko {
			hScore = score
			hPlayer = i
		}
	}

	dealerScore := g.dealer.Score()
	// for dealer - just display cards
	fmt.Println(g.dealer.Player, dealerScore)

	// announce winner
	if hScore > dealerScore {
		fmt.Printf("Player %d wins with score %d\n", hPlayer, hScore)
	} else {
		fmt.Printf("Dealer wins with score %d\n", dealerScore)
	}
}

func (g *Game) Init() {
	g.deck = deck.New(deck.Multiply(2), deck.Shuffle(nil))

	initialCards := 2

	// reset player hands
	for i, _ := range g.players {
		g.players[i].hand = make([]deck.Card, 0, initialCards)
	}
	g.dealer.hand = make([]deck.Card, 0, initialCards)

	// deal 2 cards for each player
	for i := 0; i<initialCards; i++ {
		for j, _ := range g.players {
			g.players[j].hand, g.deck = dealCard(g.players[j].hand, g.deck)
		}
		g.dealer.hand, g.deck = dealCard(g.dealer.hand, g.deck)
	}

	fmt.Println("game initialized")
	fmt.Println(g)
	fmt.Println("---------")
}

type GameOption func(*Game)

func New(players uint8, opts ...GameOption) Game {
	var g Game

	g.players = make([]Player, players, players)
	for i, _ := range g.players {
		g.players[i].name = fmt.Sprintf("player%d", i)
	}
	g.dealer.Player.name = "dealer"

	for _, opt := range opts {
		opt(&g)
	}

	return g
}