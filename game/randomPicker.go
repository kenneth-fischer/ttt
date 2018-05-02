package game

import (
	"math/rand"
)

type RandomPicker struct {
	name string
}

func NewRandomPicker() RandomPicker {
	p := RandomPicker{ name : "Randy" }
	return p
}

func (p RandomPicker) Name() string {
	return p.name
}

func (p RandomPicker) NextMove(b Board) Move {
	contents := b.Contents()
	empties := []int{}

	for i := 0; i < len(contents); i++ {
		if b.IsEmpty(i) {
			empties = append(empties, i)
		}
	}
	space := rand.Intn(len(empties))
	return Move{ Space: space, Reason: "Made random pick" }
}

