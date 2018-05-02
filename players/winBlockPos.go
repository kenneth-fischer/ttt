package players

import (
	//"fmt"
	"github.com/kenneth-fischer/ttt/game"
	//"math/rand"
)

type WinBlockPos struct {
	name     string
}

func NewWinBlockPos() WinBlockPos {
	p := WinBlockPos{ name : "Winthrop B. Pos" }
	return p
}

func (p WinBlockPos) Name() string {
	return p.name
}

func (p WinBlockPos) NextMove(b game.Board) game.Move {
	if space := p.getWinningMove(b); space > -1 {
		return game.Move{ Space: space, Reason: "I'm going to win" }
	}
	if space := p.getBlockingMove(b); space > -1 {
		return game.Move{ Space: space, Reason: "I am blocking you" }
	}
	space :=  p.pickBestSpace(b)
	return game.Move{ Space: space, Reason: "This space has a lot going on" }
}

func (p WinBlockPos) getWinningMove(b game.Board) int {
	for _, set := range b.Sets() {
		if set.CanCompleteWith(1, b.CurrentPlayer()) {
			winner := set.EmptySpaces()[0]
			return winner
		}
	}
	return -1
}

func (p WinBlockPos) getBlockingMove(b game.Board) int {
	for _, set := range b.Sets() {
		if set.CanCompleteWith(1, b.Opponent()) {
			blocker := set.EmptySpaces()[0]
			return blocker
		}
	}
	return -1
}

func (p WinBlockPos) pickBestSpace(b game.Board) int {
	moves := b.EmptySpaces()
	bestMove := -1
	bestScore := -1

	for _, move := range moves {
		score := p.scoreMove(move, b)
		if score > bestScore {
			bestScore = score
			bestMove = move
		} 
	}
	return bestMove
}

func (p WinBlockPos) scoreMove(move int, b game.Board) int {
	score := 0
	for _, set := range b.Sets() {
		if set.Contains(move) && !set.IsBlocked() {
			score += set.Count("x") + set.Count("o")
		}
	}
	return score
}
