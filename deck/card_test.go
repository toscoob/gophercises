package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNew(t *testing.T){
	t.Run("basic new", func(t *testing.T) {
		want := 13 * 4
		ans := New()
		if len(ans) != want {
			t.Errorf("got %d, want %d", len(ans), want)
		}
	})
}

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Joker
}

func TestFilters(t *testing.T){
	t.Run("filter", func(t *testing.T) {
		want := 52 - 8
		ans := New(Filter(func(c Card) bool {
			return c.Rank == Two || c.Rank == Three
		}))
		if len(ans) != want {
			t.Errorf("got %d, want %d", len(ans), want)
		}
	})

	t.Run("sort", func(t *testing.T) {
		//todo better test
		want := 52
		ans := New(SortDeck(func (deck []Card) func(i, j int) bool {
			return func(i, j int) bool {
				return deck[i].Rank < deck[j].Rank
			}
		}))
		if len(ans) != want {
			t.Errorf("got %d, want %d", len(ans), want)
		}
	})

	t.Run("sort2", func(t *testing.T) {
		//todo better test
		want := 52
		ans := New(SortDeck(LessDefault))
		if len(ans) != want {
			t.Errorf("got %d, want %d", len(ans), want)
		}
	})

	t.Run("multiply", func(t *testing.T) {
		want := 52 * 3
		ans := New(Multiply(3))
		if len(ans) != want {
			t.Errorf("got %d, want %d", len(ans), want)
		}
	})

	t.Run("add jokers", func(t *testing.T) {
		want := 4
		deck := New(AddJokers(want))
		ans := 0
		for _, card := range deck {
			if card.Suit == Joker {
				ans += 1
			}
		}
		if ans != want {
			t.Errorf("got %d, want %d", ans, want)
		}
	})
}

func TestShuffle(t *testing.T){
	want := [] struct {
		idx int
		c Card
	}{
		{ 0, Card{Three, Spade}},
		{ 8, Card{Jack, Heart}},
		{ 15, Card{Queen, Spade}},
		{ 51, Card{Three, Club}},
	}
	ans := New(Shuffle(rand.New(rand.NewSource(0))))

	for _, w := range want {
		if ans[w.idx] != w.c {
			t.Error("got", ans[w.idx], "want", w.c)
		}
	}
}