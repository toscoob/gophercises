package blackjack

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestNew(t *testing.T){
	t.Run("basic new", func(t *testing.T) {
		numPlayers := 1
		numCards := 2
		game := New(uint8(numPlayers))

		game.Init()

		if len(game.players) != 1 {
			t.Errorf("got %d, want %d", len(game.players), 1)
		}

		deckLen := 52 * 2
		wantDeckLen := deckLen - numCards * numPlayers - numCards
		if len(game.deck) != wantDeckLen {
			t.Errorf("deck len got %d, want %d", len(game.deck), wantDeckLen)
		}

		for j, p := range game.players {
			if len(p.hand) != numCards {
				t.Errorf("player %d hand len got %d, want %d", j, len(p.hand), numCards)
			}
		}

		if len(game.dealer.hand) != numCards {
			t.Errorf("dealer hand len got %d, want %d", len(game.dealer.hand), numCards)
		}
	})
}

