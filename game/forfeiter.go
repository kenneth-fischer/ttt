package game

type Forfeiter struct {
	name string
}

func NewForfeiter() Forfeiter {
	p := Forfeiter{ name : "Quitter" }
	return p
}

func (p Forfeiter) Name() string {
	return p.name
}

func (p Forfeiter) NextMove(b Board) Move {
	return Move{ Space: -1, Reason: "I just make illegal moves" }
}

