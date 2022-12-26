package d17

import (
	"fmt"
)

type Piece struct {
	x int
	y int
}

func (p *Piece) String() string {
	return fmt.Sprintf("(%v_%v)", p.x, p.y)
}

func (p *Piece) Key() string {
	return fmt.Sprintf("%v_%v", p.x, p.y)
}
func NewPiece(x int, y int) *Piece {
	p := Piece{}
	p.x = x
	p.y = y
	return &p
}
