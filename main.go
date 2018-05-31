package main

import (
	"flag"
	"fmt"
	"github.com/kenneth-fischer/ttt/game"
	"github.com/kenneth-fischer/ttt/players"
)

func main() {
	var boardSize int
	var threeD    bool

	flag.IntVar(&boardSize, "board", 3, "Choose size of board (e.g., \"3\" implies 3x3 board")
	flag.BoolVar(&threeD, "3D", false, "Play on 3D board when true, 2D otherwise")
	flag.Parse()
	description := fmt.Sprintf("%d x %d", boardSize, boardSize)
	if threeD {
		description = fmt.Sprintf("%s x %d", description, boardSize)
	}
	fmt.Printf("Starting game on %s board\n", description)
	m := game.NewMatch(boardSize, game.NewManualPlayer(), players.NewWinBlockPos())
	m.Play()
	fmt.Println(m.Status())
}
