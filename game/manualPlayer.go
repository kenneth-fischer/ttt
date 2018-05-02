package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ManualPlayer struct {
	name string
}

func NewManualPlayer() ManualPlayer {
	p := ManualPlayer{ name : "Manny" }
	return p
}

func (p ManualPlayer) Name() string {
	return p.name
}

func (p ManualPlayer) NextMove(b Board) Move {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		b.Print()
		fmt.Printf("Put %s in > ", b.CurrentPlayer())
		
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println(scanner.Err().Error())
			continue
		}
		
		space, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Try again. %s not a valid space number\n", scanner.Text())
			continue
		}

		if err = b.Mark(space); err != nil {
			fmt.Printf("Try again. %s\n", err.Error())
			continue
		}
		return Move{ Space: space, Reason: "" }

	}

	return Move{ Space: -1, Reason: "This move never should have been chosen" }
}

