package game

import (
	"fmt"
)

func getSets(b *Board) []SetOfSpaces {
        sets := []SetOfSpaces{}
	for i := 0; i < b.dimension; i++ {
                sets = append(sets, getRow(i, b))
	}
	for i := 0; i < b.dimension; i++ {
		sets = append(sets, getCol(i, b))
	}
	sets = append(sets, getLRDiag(b))
	sets = append(sets, getRLDiag(b))
	return sets
}
	
func getRow(index int, b *Board) SetOfSpaces {
	indices := []int{}
	for row := 0; row < b.dimension; row++ {
		for col := 0; col < b.dimension; col++ {
			if row == index {
				indices = append(indices, (row * b.dimension) + col)
			}
		}
	}
	name := fmt.Sprintf("Row %d", index)
	result :=  NewSetOfSpaces(name)
	result.Add(indices...)
	return result
}

func getCol(index int, b *Board) SetOfSpaces {
	indices := []int{}
	for row := 0; row < b.dimension; row++ {
		for col := 0; col < b.dimension; col++ {
			if col == index {
				indices = append(indices, (row * b.dimension) + col)
			}
		}
	}
	name := fmt.Sprintf("Col %d", index)
	result := NewSetOfSpaces(name)
	result.Add(indices...)
	return result
}

func getLRDiag(b *Board) SetOfSpaces {
	indices := []int{}
	for row := 0; row < b.dimension; row++ {
		for col := 0; col < b.dimension; col++ {
			if row == col {
				indices = append(indices, (row * b.dimension) + col)
			}
		}
	}
	result := NewSetOfSpaces("L->R Diag")
	result.Add(indices...)
	return result
}

func getRLDiag(b *Board) SetOfSpaces {
	indices := []int{}
	for row := 0; row < b.dimension; row++ {
		for col := 0; col < b.dimension; col++ {
			if row+col == b.dimension-1 {
				indices = append(indices, (row * b.dimension) + col)
			}
		}
	}
	result :=  NewSetOfSpaces("R->L Diag")
	result.Add(indices...)
	return result
}
