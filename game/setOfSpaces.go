package game

import (
	//"fmt"
	//"regexp"
)

// SetOfSpaces is used to analysis a set of spaces on the tic tac toe board
type SetOfSpaces struct {
	name        string
	board       *Board
	indices     []int
}

// NewSetOfSpaces creates a new set of spaces for the given board.
func NewSetOfSpaces(b *Board, name string, indices []int) SetOfSpaces {
	return SetOfSpaces{ board : b, name : name, indices : indices }
}

// Name returns a description name for the set of spaces
func (s SetOfSpaces) Name() string {
	return s.name
}

// EmptySpaces returns the index of each unoccupied space in the given set of spaces
func (s SetOfSpaces) EmptySpaces() []int {
	empties := []int{}

	for _, index := range s.indices {
		if s.board.IsEmpty(index) {
			empties = append(empties, index)
		}
	}
	return empties
}

func (s SetOfSpaces) CanCompleteWith(occurrences int, mark string) bool {
	if mark == marks[0] {
		return s.Count(marks[1]) == 0 && s.Count(marks[0]) >= s.board.dimension - occurrences
	}
	if mark == marks[1] {
		return s.Count(marks[0]) == 0 && s.Count(marks[1]) >= s.board.dimension - occurrences
	}
	return false
}

// CompletedBy indicates whether or not each of the spaces in the given set of spaces contains
// the given mark. This indicates that the associated player has tic tac toe.
func (s SetOfSpaces) CompletedBy(mark string) bool {
	return s.Count(mark) == s.board.dimension
}


// Contains indicates whether or not the given set of spaces contains a specific space 
// the designated number of occurrences of the designated mark.
func (s SetOfSpaces) Contains(space int) bool {
	for _, index := range s.indices {
		if index == space {
			return true
		}
	}
	return false
}

func (s SetOfSpaces) Count(mark string) int {
	count := 0
	for _, index := range s.indices {
		current := string(s.board.Contents()[index])
		if current == mark {
			count++
		}
	}
	return count
}

func (s SetOfSpaces) IsBlocked() bool {
	return s.Count(marks[0]) > 0 && s.Count(marks[1]) > 0 
}
