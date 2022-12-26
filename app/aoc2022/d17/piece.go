package d17

import (
	"fmt"
)

type Piece struct {
	x    int
	y    int
	Rock *Rock
}

func (p *Piece) String() string {
	return fmt.Sprintf("(%v_%v)", p.x, p.y)
}

func (p *Piece) Key() string {
	return fmt.Sprintf("%v_%v", p.x, p.y)
}
func NewPiece(x int, y int, rock *Rock) *Piece {
	p := Piece{}
	p.x = x
	p.y = y
	p.Rock = rock
	return &p
}
