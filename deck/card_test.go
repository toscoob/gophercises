package deck

import (
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
