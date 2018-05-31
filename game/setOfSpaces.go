package game

import (
	//"fmt"
	//"regexp"
)

// SetOfSpaces is used to analyze a set of spaces on the tic tac toe board
type SetOfSpaces struct {
	name        string
	indices     []int
}

// NewSetOfSpaces creates a new set of spaces.
func NewSetOfSpaces(name string, indices []int) SetOfSpaces {
	return SetOfSpaces{ name : name, indices : indices }
}

// Name returns a description name for the set of spaces
func (s SetOfSpaces) Name() string {
	return s.name
}

// EmptySpaces returns the index of each unoccupied space in the given set of spaces
func (s SetOfSpaces) EmptySpaces(b Board) []int {
	empties := []int{}

	for _, index := range s.indices {
		if b.IsEmpty(index) {
			empties = append(empties, index)
		}
	}
	return empties
}

func (s SetOfSpaces) CanCompleteWith(occurrences int, mark string, b Board) bool {
	if mark == marks[0] {
		return s.Count(marks[1], b) == 0 && s.Count(marks[0], b) >= b.dimension - occurrences
	}
	if mark == marks[1] {
		return s.Count(marks[0], b) == 0 && s.Count(marks[1], b) >= b.dimension - occurrences
	}
	return false
}

// CompletedBy indicates whether or not each of the spaces in the given set of spaces contains
// the given mark. This indicates that the associated player has tic tac toe.
func (s SetOfSpaces) CompletedBy(mark string, b Board) bool {
	return s.Count(mark, b) == b.dimension
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

func (s SetOfSpaces) Count(mark string, b Board) int {
	count := 0
	for _, index := range s.indices {
		current := string(b.Contents()[index])
		if current == mark {
			count++
		}
	}
	return count
}

func (s SetOfSpaces) IsBlocked(b Board) bool {
	return s.Count(marks[0], b) > 0 && s.Count(marks[1], b) > 0 
}
