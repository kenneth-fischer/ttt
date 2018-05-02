package game

import (
	"testing"
)

func TestForfeit(t *testing.T) {
	p1 := NewRandomPicker()
	p2 := NewForfeiter()

	g := NewGame(3, p1, p2)
	g.Play()
	if !g.board.IsForfeit() {
		t.Fatalf("Expected a forfeit")
	}
	if g.board.Winner() != "x" {
		t.Fatalf("Expected %q to win but %q did", "x", g.board.Winner())
	}
}
