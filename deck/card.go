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

type DeckOption func(*[]Card)

//todo looks kinda ugly
func SortDeck(less func(i, j int) bool) DeckOption {
	return func(deck *[]Card) {
		sort.Slice(*deck, func(i, j int) bool {
			return less(int((*deck)[i].Rank), int((*deck)[j].Rank))
		})
	}
}

func Shuffle() DeckOption {
	return func(deck *[]Card) {
		n := len(*deck)
		for i := 0; i < n; i++ {
			// choose index uniformly in [i, N-1]
			r := i + rand.Intn(n-i)
			(*deck)[r], (*deck)[i] = (*deck)[i], (*deck)[r]
		}
	}
}

func AddJokers(n int) DeckOption {
	return func(deck *[]Card) {
		for i := 0; i < n; i++ {
			*deck = append(*deck, Card{Suit: Joker})
		}
	}
}

func Filter(ranks ...Rank) DeckOption {
	return func(deck *[]Card) {
		var filteredDeck []Card

		//just create another deck with filtered cards
		// more optimal approach possible, but ok for now
		for _, card := range *deck {
			doFilter := false

			for _, r := range ranks {
				if card.Rank == r {
					doFilter = true
					break
				}
			}

			if !doFilter {
				filteredDeck = append(filteredDeck, card)
			}
		}

		*deck = filteredDeck
	}
}

func AddCopies(n int) DeckOption {
	return func(deck *[]Card) {
		l := len(*deck)
		for i := 0; i < n; i++ {
			*deck = append(*deck, (*deck)[:l]...)
		}
	}
}

func New(opts ...DeckOption) []Card {
	//todo add functional opts
	var res []Card

	// make default 52 card deck
	for _, s := range suits {
		for r := minRank; r <= maxRank; r++ {
			res = append(res, Card{r, s})
		}
	}
	//fmt.Println("deck before", res)

	for _, opt := range opts {
		opt(&res)
	}

	//fmt.Println("deck after", res)

	return res
}

