package main

import (
	"flag"
	"fmt"
	"github.com/kenneth-fischer/ttt/game"
	"github.com/kenneth-fischer/ttt/players"
)

func main() {
	var boardSize int
	flag.IntVar(&boardSize, "board", 3, "Choose size of board (e.g., \"3\" implies 3x3 board")
	flag.Parse()
	fmt.Printf("Starting game on %d x %d board\n", boardSize, boardSize)
	m := game.NewMatch(boardSize, game.NewManualPlayer(), players.NewWinBlockPos())
	m.Play()
	fmt.Println(m.Status())
}
