package d17

import (
	"fmt"
	"strings"
)

type Rock struct {
	pieces map[string]*Piece
	width  int
	height int
	input  string
	x      int
	y      int
	Name   string
	Number int
}

func NewRock(name, input string) *Rock {
	rock := &Rock{}
	rock.Name = name
	rock.input = input
	rock.pieces = make(map[string]*Piece)
	rows := strings.Split(input, ",")
	for y, row := range rows {
		for x := 0; x < len(row); x++ {
			value := row[x : x+1]
			if value != "." {
				p := NewPiece(x, y, rock)
				rock.AddPiece(p)
			}
		}
	}
	rock.height = len(rows)
	rock.width = len(rows[0])
	return rock
}

func (r *Rock) Clone() *Rock {
	return NewRock(r.Name, r.input)
}

func (r *Rock) Equals(rock *Rock) bool {
	return r.input == rock.input
}

func (r *Rock) Occupies(x int, y int) bool {
	r_minx := r.x
	r_maxx := r.x + r.width
	r_miny := r.y - r.height
	r_maxy := r.y
	return x >= r_minx && x <= r_maxx && y >= r_miny && y <= r_maxy
}

func (r *Rock) GetPieceAbsoluteXY(x int, y int) *Piece {
	key := fmt.Sprintf("%v_%v", x, y)
	value := r.pieces[key]
	// if value == nil {
	// 	fmt.Printf("rock.GetPieceAbsolute(%v,%v) is nil\n", x, y)
	// }
	return value
}

func (r *Rock) GetPieceChamberXY(x int, y int, DEBUG bool) *Piece {
	xx := x - r.x
	// yy := goutils.Abs(y - r.y) // r.height
	yy := r.y - y // - r.y // r.height

	if DEBUG {
		fmt.Printf("rock.getPieceChamberXY(%v,%v) (Rock=%v,%v), (x=%v, y=%v)\n", x, y, r.x, r.y, xx, yy)
	}

	if xx < 0 || xx > r.width-1 {
		return nil
	}
	if yy < 0 || yy > r.height-1 {
		return nil
	}
	key := fmt.Sprintf("%v_%v", xx, yy)
	value := r.pieces[key]
	// fmt.Printf("rock.getPieceChamberXY(%v,%v) (Rock=%v,%v), (x=%v, y=%v) = %v\n", x, y, r.x, r.y, xx, yy, value)
	return value
}

func (r *Rock) Debug() string {
	line := ""
	for row := r.height - 1; row >= 0; row-- {
		for col := 0; col < r.width; col++ {
			piece := r.GetPieceAbsoluteXY(col, row)
			if piece != nil {
				line = fmt.Sprintf("%v%v", line, "#")
			} else {
				line = fmt.Sprintf("%v%v", line, ".")
			}
		}
		if row > 0 {
			line = fmt.Sprintf("%v\n", line)
		}
	}
	return line
}

func (r *Rock) AddPiece(p *Piece) {
	r.pieces[p.Key()] = p
}

func (r *Rock) GetLeftmostPiecesX(c *Chamber) []*Piece {
	result := make([]*Piece, 0)
	if r.Equals(c.ROCK_PLUS) {
		result = append(result, r.GetPieceAbsoluteXY(1, 0))
		result = append(result, r.GetPieceAbsoluteXY(0, 1))
		result = append(result, r.GetPieceAbsoluteXY(1, 2))
	} else if r.Equals(c.ROCK_L) {
		result = append(result, r.GetPieceAbsoluteXY(2, 0))
		result = append(result, r.GetPieceAbsoluteXY(2, 1))
		result = append(result, r.GetPieceAbsoluteXY(0, 2))
	} else if r.Equals(c.ROCK_HORIZONTAL) {
		result = append(result, r.GetPieceAbsoluteXY(0, 0))
	} else if r.Equals(c.ROCK_VERTICAL) {
		result = append(result, r.GetPieceAbsoluteXY(0, 0))
		result = append(result, r.GetPieceAbsoluteXY(0, 1))
		result = append(result, r.GetPieceAbsoluteXY(0, 2))
		result = append(result, r.GetPieceAbsoluteXY(0, 3))
	} else if r.Equals(c.ROCK_SQUARE) {
		result = append(result, r.GetPieceAbsoluteXY(0, 0))
		result = append(result, r.GetPieceAbsoluteXY(0, 1))
	} else {
		panic("no leftmost rock piece")
	}
	return result
}

func (r *Rock) GetRightmostPiecesX(c *Chamber) []*Piece {
	result := make([]*Piece, 0)
	if r.Equals(c.ROCK_PLUS) {
		result = append(result, r.GetPieceAbsoluteXY(1, 0))
		result = append(result, r.GetPieceAbsoluteXY(2, 1))
		result = append(result, r.GetPieceAbsoluteXY(1, 2))
	} else if r.Equals(c.ROCK_L) {
		result = append(result, r.GetPieceAbsoluteXY(2, 0))
		result = append(result, r.GetPieceAbsoluteXY(2, 1))
		result = append(result, r.GetPieceAbsoluteXY(2, 2))
	} else if r.Equals(c.ROCK_HORIZONTAL) {
		result = append(result, r.GetPieceAbsoluteXY(3, 0))
	} else if r.Equals(c.ROCK_VERTICAL) {
		result = append(result, r.GetPieceAbsoluteXY(0, 0))
		result = append(result, r.GetPieceAbsoluteXY(0, 1))
		result = append(result, r.GetPieceAbsoluteXY(0, 2))
		result = append(result, r.GetPieceAbsoluteXY(0, 3))
	} else if r.Equals(c.ROCK_SQUARE) {
		result = append(result, r.GetPieceAbsoluteXY(1, 0))
		result = append(result, r.GetPieceAbsoluteXY(1, 1))

	} else {
		panic("no rightmost rock piece")
	}

	// fmt.Printf(":this rock has %v rightmost pieces\n", len(result))
	return result
}

func (r *Rock) GetBottomPiecesX(c *Chamber) []*Piece {
	result := make([]*Piece, 0)
	if r.Equals(c.ROCK_PLUS) {
		result = append(result, r.GetPieceAbsoluteXY(0, 1))
		result = append(result, r.GetPieceAbsoluteXY(1, 2))
		result = append(result, r.GetPieceAbsoluteXY(2, 1))
	} else if r.Equals(c.ROCK_L) {
		result = append(result, r.GetPieceAbsoluteXY(0, 2))
		result = append(result, r.GetPieceAbsoluteXY(1, 2))
		result = append(result, r.GetPieceAbsoluteXY(2, 2))
	} else if r.Equals(c.ROCK_HORIZONTAL) {
		result = append(result, r.GetPieceAbsoluteXY(0, 0))
		result = append(result, r.GetPieceAbsoluteXY(1, 0))
		result = append(result, r.GetPieceAbsoluteXY(2, 0))
		result = append(result, r.GetPieceAbsoluteXY(3, 0))
	} else if r.Equals(c.ROCK_VERTICAL) {
		result = append(result, r.GetPieceAbsoluteXY(0, 3))
	} else if r.Equals(c.ROCK_SQUARE) {
		result = append(result, r.GetPieceAbsoluteXY(0, 1))
		result = append(result, r.GetPieceAbsoluteXY(1, 1))

	} else {
		panic("no bottom rock piece")
	}
	return result
}
