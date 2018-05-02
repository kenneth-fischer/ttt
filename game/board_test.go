package game

import (
	"fmt"
	"testing"
)

func TestUndo(t *testing.T) {
	b := NewBoard(3)

	assertIsEmpty(t, *b, 4)
	b.Mark(4)
	assertNotEmpty(t, *b, 4)

	assertIsEmpty(t, *b, 8)
	b.Mark(8)
	assertNotEmpty(t, *b, 8)

	b.Undo()
	assertIsEmpty(t, *b, 8)

	b.Undo()
	assertIsEmpty(t, *b, 4)
}

func TestTieGame(t *testing.T) {
	b := NewBoard(3)

	b.Mark(4)
	b.Mark(0)
	b.Mark(2)
	b.Mark(6)
	b.Mark(3)
	b.Mark(5)
	b.Mark(7)
	b.Mark(1)
	b.Mark(8)
	if b.Winner() != "" {
		t.Fatalf("%q was not expected to win", b.Winner())
	}
	if !b.IsTie() {
		t.Fatalf("Expected a tie game")
	}
}

func assertEqual(t *testing.T, expected, actual interface{}, msgParts ...string) {
	if expected != actual {
		msg := fmt.Sprintf("Expected %v. Got %v", expected, actual)
		if len(msgParts) > 0 {
			msg = fmt.Sprintf("%s. %s", msg, msgParts[0])
		}
		t.Fatalf(msg)
	}
}

func assertIsEmpty(t *testing.T, b Board, space int) {
	if !b.IsEmpty(space) {
		t.Fatalf("Space %d should be empty", space)
	}
}

func assertNotEmpty(t *testing.T, b Board, space int) {
	if b.IsEmpty(space) {
		t.Fatalf("Space %d should be empty", space)
	}
}
