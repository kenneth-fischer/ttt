package game

import (
	"fmt"
)

type Match struct {
	boardSize int
	players   []Player
	games     []*Game
}

func NewMatch(boardSize int, player1, player2 Player) Match {
	m := Match{}
	m.boardSize = boardSize
	m.players = []Player { player1, player2 }
	game1 := NewGame(boardSize, player1, player2)
	game2 := NewGame(boardSize, player2, player1)
	m.games = []*Game{ game1, game2 }
	return m
}

func (m Match) Games() []Game {
	games := []Game{}
	for _, game := range m.games {
		games = append(games, *game)
	}
	return games
}

func (m Match) Players() []Player {
	return m.players
}	

func (m *Match) Play() {
	fmt.Println("Start game 1")
	fmt.Println()
	m.games[0].Play()
	m.games[0].Board().Print()
	fmt.Println(m.games[0].Board().Status())
	fmt.Println()
	fmt.Println("Start game 2")
	fmt.Println()
	m.games[1].Play()
	m.games[1].Board().Print()
	fmt.Println(m.games[1].Board().Status())
	fmt.Println()
}

func (m Match) Status() string {
	if !m.games[0].IsOver() {
		return "Game 1 in progress"
	}
	if !m.games[1].IsOver() {
		return fmt.Sprintf("Game 2 in progress. Game 1: %s", m.games[0].Status()) 
	}
	wins := []int{ 0, 0 }

	game := m.games[0]
	if !game.IsTie() {
		if game.Winner() == marks[0] {
			wins[0] += 1
		} else {
			wins[1] += 1
		}
	}

	game = m.games[1]
	if !game.IsTie() {
		if game.Winner() == marks[1] {
			wins[0] += 1
		} else {
			wins[1] += 1
		}
	}

	if wins[0] == wins[1] {
		return "Tie match"
	}
	if wins[0] > wins[1] {
		return fmt.Sprintf("Player 1 wins match %d to %d", wins[0], wins[1])
	}
	return fmt.Sprintf("Player 2 wins match %d to %d", wins[1], wins[0])
}
		
