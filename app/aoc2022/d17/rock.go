package d17

import (
	"fmt"
	"strings"

	"github.com/simonski/goutils"
)

type Rock struct {
	pieces map[string]*Piece
	width  int
	height int
	input  string
	x      int
	y      int
}

func NewRock(input string) *Rock {
	rock := &Rock{}
	rock.input = input
	rock.pieces = make(map[string]*Piece)
	rows := strings.Split(input, ",")
	for y, row := range rows {
		for x := 0; x < len(row); x++ {
			value := row[x : x+1]
			p := NewPiece(x, y, value == "#")
			rock.AddPiece(p)
		}
	}
	rock.height = len(rows)
	rock.width = len(rows[0])
	return rock
}

func (r *Rock) Clone() *Rock {
	return NewRock(r.input)
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

	// return (x >= r.x && x <= (r.x+r.width)) && (y >= r.y && y <= (r.y+r.height))
}

func (r *Rock) GetPiece(x int, y int) *Piece {
	xx := x - r.x
	yy := goutils.Abs(y - r.y) // r.height
	key := fmt.Sprintf("%v_%v", xx, yy)
	value := r.pieces[key]
	// fmt.Printf("rock.getPiece(%v,%v) (x=%v, y=%v) = %v\n", x, y, xx, yy, value)
	return value
}

func (r *Rock) Debug() string {
	result := ""
	for index, line := range strings.Split(r.input, ",") {
		if index > 0 {
			result = fmt.Sprintf("%v\n%v", result, line)
		} else {
			result = line
		}
	}
	return result
}

func (r *Rock) AddPiece(p *Piece) {
	r.pieces[p.Key()] = p
}

func (r *Rock) GetLeftmostPieces() []*Piece {
	result := make([]*Piece, 0)
	for _, piece := range r.pieces {
		if piece.x == 0 {
			result = append(result, piece)
		}
	}
	return result
}

func (r *Rock) GetRightmostPieces() []*Piece {
	result := make([]*Piece, 0)
	for _, piece := range r.pieces {
		is_right := piece.x+1 == r.width
		// fmt.Printf("piece(%v,%v), rock width %v right?=%v\n", piece.x, piece.y, r.width, is_right)
		if is_right {
			result = append(result, piece)
		}
	}
	// fmt.Printf(":this rock has %v rightmost pieces\n", len(result))
	return result
}

func (r *Rock) GetBottomPieces() []*Piece {
	result := make([]*Piece, 0)
	for _, piece := range r.pieces {
		outcome := piece.y+1 == r.height
		// fmt.Printf("bottom(%v,%v) on rock height %v, result=%v\n", piece.x, piece.y, r.height, outcome)
		if outcome {
			result = append(result, piece)
		}
	}
	return result
}
