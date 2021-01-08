package deck

import (
	"sort"
)

//go:generate stringer -type=Suit
type Suit int

const (
	None Suit = iota - 1
	Clubs
	Diamonds
	Hearts
	Spades
	suitCount int = iota - 1
)

//go:generate stringer -type=Rank
type Rank int

const (
	Joker Rank = iota
	A
	_ //2
	_ //3
	_ //4
	_ //5
	_ //6
	_ //7
	_ //8
	_ //9
	_ //10
	J
	Q
	K
	rankCount int = iota
)

type Card struct {
	Rank Rank
	Suit Suit
}

type DeckOption func(*[]Card)

func SortDeck(less func(i, j int) bool) DeckOption {
	return func(deck *[]Card) {
		sort.Slice(*deck, func(i, j int) bool {
			return less(int((*deck)[i].Rank), int((*deck)[j].Rank))
		})
	}
}

func Shuffle() DeckOption {
	return func(deck *[]Card) {

	}
}

func AddJokers(n int) DeckOption {
	return func(deck *[]Card) {

	}
}

func Filter(r ...Rank) DeckOption {
	return func(deck *[]Card) {

	}
}

func Multiple(n int) DeckOption {
	return func(deck *[]Card) {

	}
}

func New(opts ...DeckOption) []Card {
	//todo add functional opts
	var res []Card

	// make default 52 card deck
	for i := 0; i < suitCount; i++ {
		for j := int(A); j < rankCount; j++ {
			res = append(res, Card{Rank(j), Suit(i)})
		}
	}
	//fmt.Println("deck before", res)

	for _, opt := range opts {
		opt(&res)
	}

	//fmt.Println("deck after", res)

	return res
}

