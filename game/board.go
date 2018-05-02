package game

import (
	"fmt"
	"regexp"
)

var (
	marks = []string{ "x", "o"}
)

const (
	empty = "-"
)

type Board struct {
	dimension int
	contents  string
	moves     []int
	winner    string
	isForfeit   bool
	rows      []SetOfSpaces
	cols      []SetOfSpaces
	lrdiag    SetOfSpaces
	rldiag    SetOfSpaces
}

func NewBoard(dimension int) *Board {
	b := Board{ dimension : dimension }
	for i := 0; i < dimension; i++ {
		b.rows = append(b.rows, getRow(i, &b))
	}
	for i := 0; i < dimension; i++ {
		b.cols = append(b.cols, getCol(i, &b))
	}
	b.lrdiag = getLRDiag(&b)
	b.rldiag = getRLDiag(&b)
	return &b
}

func (b Board) Contents() string {
	if len(b.contents) == 0 {
		contents := ""
		for row := 0; row < b.dimension; row++ {
			for col := 0; col < b.dimension; col++ {
				contents += empty
			}
		}
		return contents
	}
	return b.contents
}

func (b Board) ContentsOf(index int) string {
	if index < 0 || index >= b.Spaces() {
		return ""
	}
	return  string(b.contents[index])
}

func (b Board) Rows() []SetOfSpaces {
	return b.rows
}

func (b Board) Cols() []SetOfSpaces {
	return b.cols
}

func (b Board) LRDiag() SetOfSpaces {
	return b.lrdiag
}

func (b Board) RLDiag() SetOfSpaces {
	return b.rldiag
}

func (b Board) EmptySpaces() []int {
	results := []int{}

	for i := 0; i < b.Spaces(); i++ {
		if b.IsEmpty(i) {
			results = append(results, i)
		}
	}
	return results
}
 
func (b Board) Sets() []SetOfSpaces {
	sets :=[]SetOfSpaces{}
	sets = append(sets, b.rows...)
	sets = append(sets, b.cols...)
	sets = append(sets, b.lrdiag, b.rldiag)
	return sets
}

// Opponent returns the mark of the player who is not the current player
func (b Board) Opponent() string {
	if b.CurrentPlayer() == marks[0] {
		return marks[1]
	}
	return marks[0]
}

// CurrentPlayer returns the mark of the player who will/would make the next turn
func (b Board) CurrentPlayer() string {
	markIndex := len(b.moves)%len(marks)
	return marks[markIndex]
}

func (b Board) IsEmpty(index int) bool {
	pattern := ""
	for i := 0; i < b.Spaces(); i++ {
		if i == index {
			pattern += empty
		} else {
			pattern += "."
		}
	}
	match, _ := regexp.MatchString(pattern, b.Contents())
	return match
}

func (b Board) IsForfeit() bool {
	return b.isForfeit
}

func (b Board) IsGameOver() bool {
	if b.winner != "" {
		return true
	}

	if len(b.moves) == b.Spaces() {
		return true
	}
	return false
}

func (b Board) IsTie() bool {
	return b.IsGameOver() && b.Winner() == ""
}

func (b Board) LastMove() int {
	if len(b.moves) == 0 {
		return -1
	}
	return b.moves[len(b.moves) - 1]
}

func (b *Board) Mark(index int) error {
	orig := b.Contents()
		
	if index < 0 || index >= len(orig) {
		return fmt.Errorf("%d is out of bounds", index)
	}

	if b.IsGameOver() {
		return fmt.Errorf("Game is over")
	}

	if !b.IsEmpty(index) {
		return fmt.Errorf("%d is not empty", index)
	}
	
	updated := string(orig[:index]) + b.CurrentPlayer() + string(orig[index+1:])
	b.contents  = updated
	b.checkWinner()
	b.moves = append(b.moves, index)
	return nil
}

func (b Board) Moves() []int {
	return b.moves
}

func (b Board) Print() {
	contents := b.Contents()
	for index := 0; index < len(contents); index++ {
		if index > 0 && index % b.dimension == 0 {
			fmt.Println()
		}
		
		mark := string(contents[index])
		fmt.Print(mark)		
	}
	fmt.Println()
	fmt.Println()
}

func (b Board) Spaces() int {
	return b.dimension * b.dimension
}

func (b Board) Status() string {
	if b.IsTie() {
		return "Tie game"
	}

	if b.IsGameOver() {
		return fmt.Sprintf("%s wins\n", b.Winner())
	}

	return "Game in progress"
}

func (b Board) Winner() string {
	return b.winner
}

func (b *Board) Undo() {
	if len(b.moves) == 0 {
		return
	}
	index := b.moves[len(b.moves) - 1]
	
	orig := b.Contents()
	updated := string(orig[:index]) + empty + string(orig[index+1:])
	b.contents  = updated
	b.moves = b.moves[:len(b.moves)-1]
}

func (b *Board) checkWinner() bool {
	if len(b.moves) < 4 {
		return false
	}

	for _, set := range b.Sets() {
		if set.CompletedBy(b.CurrentPlayer()) {
			b.winner = b.CurrentPlayer()
			return true
		}
	}
	return false
}

func (b *Board) forfeit() {
	if b.CurrentPlayer() == marks[0] {
		b.winner = marks[1]
		return
	}
	b.isForfeit = true
	b.winner = marks[0]
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
	return NewSetOfSpaces(b, name, indices)
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
	return NewSetOfSpaces(b, name, indices)
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
	return NewSetOfSpaces(b, "L->R Diag", indices)
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
	return NewSetOfSpaces(b, "R->L Diag", indices)
}
