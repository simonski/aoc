package d17

import (
	"fmt"
)

type Piece struct {
	x     int
	y     int
	solid bool
}

func (p *Piece) Key() string {
	return fmt.Sprintf("%v_%v", p.x, p.y)
}
func NewPiece(x int, y int, solid bool) *Piece {
	p := Piece{}
	p.x = x
	p.y = y
	p.solid = solid
	return &p
}
