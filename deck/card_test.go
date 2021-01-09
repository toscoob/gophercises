package deck

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T){
	t.Run("basic new", func(t *testing.T) {
		want := 52
		ans := New()
		if len(ans) != want {
			t.Errorf("got %d, want %d", len(ans), want)
		}
		ans2 := New(SortDeck(func(i, j int) bool {
			return j > i
		}))
		if len(ans2) != want {
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
		ans := New(Filter(Two, Three))
		if len(ans) != want {
			t.Errorf("got %d, want %d", len(ans), want)
		}
	})

	t.Run("sort", func(t *testing.T) {
		//todo better test
		want := 52
		ans := New(SortDeck(func(i, j int) bool {
			return j > i
		}))
		if len(ans) != want {
			t.Errorf("got %d, want %d", len(ans), want)
		}
	})

	t.Run("shuffle", func(t *testing.T) {
		//todo find out how to test shuffling
		want := 52
		ans := New(Shuffle())
		if len(ans) != want {
			t.Errorf("got %d, want %d", len(ans), want)
		}
	})

	t.Run("multiply", func(t *testing.T) {
		want := 52 * 3
		ans := New(AddCopies(2))
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
