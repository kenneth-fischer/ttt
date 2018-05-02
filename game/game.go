package game

import (
	"fmt"
)

type Game struct {
	board   *Board
	players []Player
}

func NewGame(boardSize int, player1, player2 Player) *Game {
	board := NewBoard(boardSize)
	players := []Player{ player1, player2 }
	game := Game{ board: board, players: players } 
	return &game
}

func (g Game) Board() Board {
	return *g.board
}

func (g Game) IsOver() bool {
	return g.board.IsGameOver()
}

func (g Game) IsTie() bool {
	return g.board.IsTie()
}

func (g Game) Winner() string {
	return g.board.Winner()
}

func (g Game) Status() string {
	return g.board.Status()
}

func (g Game) Players() []Player {
	return g.players
}	

func (g *Game) Play() { 
	for done := g.board.IsGameOver(); !done; done = g.board.IsGameOver() {
		move, err := g.NextMove()
		if err != nil {
			g.board.forfeit()
			fmt.Printf("%d is an invalid move. %s wins by forfeit. %s\n", move.Space, g.board.Winner(), err.Error())
			return
		}
		if move.Reason != "" {
			fmt.Println(move.Reason)
		}
	}
}

func (g *Game) NextMove() (Move, error) { 
	playerIndex := len(g.board.Moves())%2
	player := g.players[playerIndex]
	move := player.NextMove(*g.board)
	err :=  g.board.Mark(move.Space)
	return move, err
}
