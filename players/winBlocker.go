package players

import (
	//"fmt"
	"github.com/kenneth-fischer/ttt/game"
	"math/rand"
)

type WinBlocker struct {
	name     string
}

func NewWinBlocker() WinBlocker {
	p := WinBlocker{ name : "Winnie" }
	return p
}

func (p WinBlocker) Name() string {
	return p.name
}

func (p WinBlocker) NextMove(b game.Board) game.Move {
	if space := p.getWinningMove(b); space > -1 {
		return game.Move{ Space: space, Reason: "I'm going to win" }
	}
	if space := p.getBlockingMove(b); space > -1 {
		return game.Move{ Space: space, Reason: "I am blocking you" }
	}
	space :=  p.pickRandomMove(b)
	return game.Move{ Space: space, Reason: "I randomly picked this move" }
}

func (p WinBlocker) getWinningMove(b game.Board) int {
	for _, set := range b.Sets() {
		if set.CanCompleteWith(1, b.CurrentPlayer()) {
			winner := set.EmptySpaces()[0]
			//fmt.Printf("Pick winner: %d\n", winner)
			return winner
		}
	}
	return -1
}

func (p WinBlocker) getBlockingMove(b game.Board) int {
	for _, set := range b.Sets() {
		if set.CanCompleteWith(1, b.Opponent()) {
			blocker := set.EmptySpaces()[0]
			//fmt.Printf("Pick blocker: %d\n", blocker)
			return blocker
		}
	}
	return -1
}

func (p WinBlocker) pickRandomMove(b game.Board) int {
	empties := b.EmptySpaces()
	index := rand.Intn(len(empties))
	move := empties[index]
	//fmt.Printf("Random Pick: %d\n", move)
	return move
}
