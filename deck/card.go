package deck

import (
	"fmt"
	"math/rand"
	"sort"
)

//go:generate stringer -type=Suit
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Heart
	Club
	Joker
	//suitCount int = iota - 1
)

var suits = [...]Suit{ Spade, Diamond, Heart, Club}

//go:generate stringer -type=Rank
type Rank uint8

const (
	_ Rank = iota //to make it 0-based
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	//rankCount int = iota
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Rank
	Suit
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}


type DeckOption func([]Card) []Card

func (c Card) absRank() int {
	return int(c.Suit) * int(maxRank) + int(c.Rank)
}

func LessDefault(deck []Card) func(i, j int) bool {
	return func(i, j int) bool{
		return deck[i].absRank() < deck[j].absRank()
	}
}

func SortDeck(less func (deck []Card) func(i, j int) bool) DeckOption {
	return func(deck []Card) []Card {
		sort.Slice(deck, less(deck))
		return deck
	}
}

func Shuffle(r *rand.Rand) DeckOption {
	return func(deck []Card) []Card {
		n := len(deck)

		randNum := func (n int) int {
			if r != nil {
				return r.Intn(n)
			}

			return rand.Intn(n)
		}

		for i := 0; i < n; i++ {
			// choose index uniformly in [i, N-1]
			r := i + randNum(n-i)
			(deck)[r], (deck)[i] = (deck)[i], (deck)[r]
		}
		return deck
	}
}

func AddJokers(n int) DeckOption {
	return func(deck []Card) []Card {
		for i := 0; i < n; i++ {
			deck = append(deck, Card{Suit: Joker})
		}
		return deck
	}
}

func Filter(f func (c Card) bool) DeckOption {
	return func(deck []Card) []Card {
		var filteredDeck []Card

		//just create another deck with filtered cards
		// more optimal approach possible, but ok for now
		for _, card := range deck {
			if !f(card) {
				filteredDeck = append(filteredDeck, card)
			}
		}

		return filteredDeck
	}
}

func Multiply(n int) DeckOption {
	return func(deck []Card) []Card {
		if n >= 2 {
			l := len(deck)
			for i := 0; i < n-1; i++ {
				deck = append(deck, deck[:l]...)
			}
		}

		return deck
	}
}

func New(opts ...DeckOption) []Card {
	var res []Card

	// make default 52 card deck
	for _, s := range suits {
		for r := minRank; r <= maxRank; r++ {
			res = append(res, Card{r, s})
		}
	}
	//fmt.Println("deck before", res)

	for _, opt := range opts {
		res = opt(res)
	}

	//fmt.Println("deck after", res)

	return res
}

