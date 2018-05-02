package game

type Player interface {
	NextMove(b Board) Move
	Name() string
}

type Move struct {
	Space  int
	Reason string
}
