package d17

import (
	"fmt"

	"github.com/simonski/goutils"
)

type Chamber struct {
	rockTick        int
	Input           string
	ROCK_HORIZONTAL *Rock
	ROCK_VERTICAL   *Rock
	ROCK_PLUS       *Rock
	ROCK_L          *Rock
	ROCK_SQUARE     *Rock
	// RocksXY         map[string]*Rock
	Rocks       []*Rock
	CurrentRock *Rock
	Width       int
	Height      int
	Pieces      map[string]*Piece
}

func NewChamber(input string) *Chamber {
	c := Chamber{Input: input}
	c.ROCK_HORIZONTAL = NewRock("H", "####")
	c.ROCK_PLUS = NewRock("P", ".#.,###,.#.")
	c.ROCK_L = NewRock("L", "..#,..#,###")
	c.ROCK_VERTICAL = NewRock("V", "#,#,#,#")
	c.ROCK_SQUARE = NewRock("SQ", "##,##")
	c.Width = 7
	c.Pieces = make(map[string]*Piece)
	return &c
}

func (c *Chamber) NewRock() *Rock {
	c.rockTick++
	if c.rockTick == 6 {
		c.rockTick = 1
	}
	tick := c.rockTick

	if tick == 1 {
		return c.ROCK_HORIZONTAL.Clone()
	} else if tick == 2 {
		return c.ROCK_PLUS.Clone()
	} else if tick == 3 {
		return c.ROCK_L.Clone()
	} else if tick == 4 {
		return c.ROCK_VERTICAL.Clone()
	} else if tick == 5 {
		return c.ROCK_SQUARE.Clone()
	} else {
		return nil
	}
}

// func (c *Chamber) Height() int {

// 	if len(c.Rocks) == 0 {
// 		return 0
// 	}
// 	y := 0
// 	for _, rock := range c.Rocks {
// 		if rock.y > y {
// 			y = rock.y
// 		}
// 	}
// 	return y + 1
// }

func (c *Chamber) AddRock(rock *Rock) {
	// reset the height
	maxY := c.Height

	// fmt.Println(rock.Debug())
	// fmt.Printf("AddRock - initial height was %v.\n", maxY)
	// fmt.Println(rock.Debug())

	c.CurrentRock = rock
	c.Rocks = append(c.Rocks, rock)
	rock.Number = len(c.Rocks)

	// make this rock 2 units space left three units above highest rock (or floor)
	rock.x = 2
	// rock.y = maxY + 3 + (rock.height) - 1
	rock.y = maxY + 3 + (rock.height - 1)
	c.Height = rock.y + 1

	c.AddRockToMap(rock)

	// rock.y =
	// fmt.Printf("AddRock - rock x,y is (%v,%v), width=%v, height=%v\n", rock.x, rock.y, rock.width, rock.height)
	// fmt.Printf("AddRock - c height is now %v, rock x=%v, rock y=%v, rock width=%v, rock height=%v\n", c.Height, rock.x, rock.y, rock.width, rock.height)

}

func (c *Chamber) CanRockMoveLeft(rock *Rock) bool {
	if rock.x == 0 {
		return false // just can't go left anymore
	}
	// otherwise, is it obscured by anything if it goes left?
	for _, piece := range rock.pieces {
		x := rock.x + piece.x - 1
		y := rock.y - piece.y
		// answer := c.IsOccupiedByRockOrEmpty(x, y, rock)
		// if rock.Number == 8 || rock.Number == 9 {
		// 	r := c.GetRock(x, y)
		// 	reason := ""
		// 	if r == nil {
		// 		reason = "rock is nil"
		// 	} else {
		// 		reason = "rock is not nil"
		// 		if r == rock {
		// 			reason += ", rock is same"
		// 		} else {
		// 			reason += fmt.Sprintf(", rock is not same (%v != %v)", r.Number, rock.Number)
		// 		}
		// 	}

		// 	fmt.Printf("CanRockMoveLeft(rock=%v), (rock x,y=%v,%v), (piece x,y=%v,%v), (chamber x,y=%v,%v), answer %v, reason=%v\n", rock.Number, rock.x, rock.y, piece.x, piece.y, x, y, answer, reason)
		// }

		if !c.IsOccupiedByRockOrEmpty(x, y, rock) {
			// if c.IsOccupied(x, y) || rock.GetPieceChamberXY(x, y, false) != nil {
			// obscured
			return false
		}
	}
	return true
}

func (c *Chamber) CanRockMoveRight(rock *Rock) bool {
	if rock.x+rock.width >= c.Width {
		return false
	}
	// otherwise, is it obscured by anything if it goes left?
	for _, piece := range rock.pieces { // GetRightmostPieces(c) {
		x := rock.x + piece.x + 1
		if x > c.Width {
			return false
		}
		y := rock.y - piece.y
		if !c.IsOccupiedByRockOrEmpty(x, y, rock) {
			return false

			// 	occupied := c.IsOccupied(x, y) || rock.GetPieceChamberXY(x, y, false) != nil

			// // fmt.Printf("There are %v rightmost pieces, this piece is (%v,%v) on rock (%v,%v) which is (%v,%v) on the chamber, occupied=%v\n", len(rock.GetRightmostPieces()), rock.x, rock.y, piece.x, piece.y, x, y, occupied)
			// if occupied {
			// 	// obscured
		}
	}
	return true
}

func (c *Chamber) CanRockMoveDown(rock *Rock) bool {
	// if rock.y-1 < 0 {
	// 	return false
	// }
	// otherwise, is it obscured by anything if it goes left?

	for _, piece := range rock.pieces {
		x := rock.x + piece.x
		y := rock.y - piece.y - 1
		if y < 0 {
			return false
		}
		if !c.IsOccupiedByRockOrEmpty(x, y, rock) {
			// outcome := c.IsOccupied(x, y) || rock.GetPieceChamberXY(x, y, false) != nil
			// // fmt.Printf("bottom piece (%v,%v) which is (%v,%v) in the chamber, outcome=%v\n", piece.x, piece.y, x, y, outcome)
			// if outcome {
			// obscured
			return false
		}
	}
	return true
}

func (c *Chamber) MoveLeft(rock *Rock) {
	c.RemoveRockFromMap(rock)
	rock.x -= 1
	c.AddRockToMap(rock)
}

func (c *Chamber) MoveRight(rock *Rock) {
	c.RemoveRockFromMap(rock)
	rock.x += 1
	c.AddRockToMap(rock)
}

func (c *Chamber) MoveDown(rock *Rock) {
	c.RemoveRockFromMap(rock)
	rock.y -= 1
	c.AddRockToMap(rock)
	y := 0
	for _, r := range c.Rocks {
		y = goutils.Max(y, r.y)
	}
	if y == 0 {
		y = 1
		c.Height = y
	} else {
		c.Height = y + 1

	}

}

func (c *Chamber) RemoveRockFromMap(rock *Rock) {
	remove := make([]string, 0)
	for key := range c.Pieces {
		piece := c.Pieces[key]
		if piece != nil {
			if piece.Rock == rock {
				remove = append(remove, key)
			}
		}
	}

	for _, key := range remove {
		delete(c.Pieces, key)
	}

}

func (c *Chamber) AddRockToMap(rock *Rock) {

	// fmt.Printf("AddRockToMap(%v,%v)\n%v\n", rock.x, rock.y, rock.Debug())
	// reset the pieces in the map
	for row := 0; row < rock.height; row++ {
		for col := 0; col < rock.width; col++ {
			new_x := col + rock.x
			new_y := rock.y - row
			key := fmt.Sprintf("%v_%v", new_x, new_y)
			piece := rock.GetPieceAbsoluteXY(col, row)
			if piece != nil {
				// if rock.Number == 22 {
				// 	piece_exists := piece != nil
				// 	fmt.Printf("AddRockToMap(rock=%v) col,row=%v,%v, new_x=%v, new_y=%v, key=%v, exists=%v\n", rock.Number, col, row, new_x, new_y, key, piece_exists)
				// }
				if c.Pieces[key] == nil {
					c.Pieces[key] = piece
				} else {
					fmt.Println(c.Debug())
					otherRock := c.Pieces[key].Rock
					fmt.Printf("Rock %v is trying to overwrite a piece from rock %v at position [%v] ", rock.Number, otherRock.Number, key)
					panic("foo")
					// os.os.Exit(1)
				}
			}
		}
	}
}

// move the rock, or if it cannot move, return false
func (c *Chamber) Tick(instruction string, VERBOSE bool, VERY_VERBOSE bool) bool {
	if instruction == "<" {
		if c.CanRockMoveLeft(c.CurrentRock) {
			c.MoveLeft(c.CurrentRock)
			if VERY_VERBOSE {
				fmt.Println("Jet of gas pushes rock left:")
				fmt.Println(c.Debug())
			}
		} else {
			if VERY_VERBOSE {
				fmt.Println("Jet of gas pushes rock left, but nothing happens:")
				fmt.Println(c.Debug())
			}
		}
	} else if instruction == ">" {
		if c.CanRockMoveRight(c.CurrentRock) {
			c.MoveRight(c.CurrentRock)
			if VERY_VERBOSE {
				fmt.Println("Jet of gas pushes rock right")
				fmt.Println(c.Debug())
			}
		} else {
			if VERY_VERBOSE {
				fmt.Println("Jet of gas pushes rock right, but nothing happens:")
				fmt.Println(c.Debug())
			}

		}
	}

	if c.CanRockMoveDown(c.CurrentRock) {
		c.MoveDown(c.CurrentRock)
		if VERY_VERBOSE {
			fmt.Println("Rock falls 1 unit:")
			fmt.Println(c.Debug())
		}
		return true
	} else {
		if VERY_VERBOSE {
			fmt.Println("Rock falls 1 unit, causing it to come to rest:")
			fmt.Println(c.Debug())
		}
	}
	return false
}

func (c *Chamber) Run(VERBOSE bool, VERY_VERBOSE bool, breakAfterRock int) {
	rock := c.NewRock()
	c.AddRock(rock)
	if VERBOSE {
		fmt.Printf("\n[%v] %v\n%v\n", 0, "BEGIN", c.Debug())
	}
	rockCount := 1
	index := -1
	for {
		index++

		// if index+2 == breakAfterRock {
		// 	VERBOSE = true
		// 	VERY_VERBOSE = true
		// }
		if index == len(c.Input) {
			index = 0
		}
		instruction := c.Input[index : index+1]
		if !c.Tick(instruction, VERBOSE, VERY_VERBOSE) {
			c.CurrentRock = nil
			if rockCount >= breakAfterRock {
				// if VERBOSE {
				// 	fmt.Println(c.Debug())
				// }
				// fmt.Printf("After %v rocks, size is %v\n", rockCount, c.Height)
				return
			}

			rock := c.NewRock()
			c.AddRock(rock)
			if VERBOSE {
				fmt.Printf("A new rock begins falling\n%v\n", c.Debug())
			}
			rockCount++
		}
		// if VERBOSE {
		// 	fmt.Printf("[%v] after %v, still moving\n%v\n", index, instruction, c.Debug())
		// }
	}
}

func (c *Chamber) Debug() string {
	// |..@@@@.|
	// |.......|
	// |.......|
	// |.......|
	// +-------+
	// for i, rock := range c.Rocks {
	// fmt.Printf("rock[%v], x=%v, y=%v\n", i, rock.x, rock.y)
	// fmt.Println(rock.Debug())
	// }
	line := ""
	for row := c.Height - 1; row >= 0; row-- {
		line = fmt.Sprintf("%v|", line)
		for col := 0; col < c.Width; col++ {
			piece := c.GetRockPieceString(col, row)
			// if piece != "." {
			// fmt.Printf("c.Debug(%v,%v)=%v\n", col, row, piece)
			// }
			line = fmt.Sprintf("%v%v", line, piece)
		}
		line = fmt.Sprintf("%v|", line)
		line = fmt.Sprintf("%v     %v\n", line, row)
	}
	line = fmt.Sprintf("%v+-------+", line)
	fmt.Printf("Rocks=%v, Height=%v\n", len(c.Rocks), c.Height)
	return line
}

func (c *Chamber) GetRockPiece(x int, y int) *Piece {
	key := fmt.Sprintf("%v_%v", x, y)
	return c.Pieces[key]
}

func (c *Chamber) GetRockPieceString(x int, y int) string {
	key := fmt.Sprintf("%v_%v", x, y)
	piece := c.Pieces[key]
	// fmt.Printf("GetRockPieceString(%v,%v)=%v\n", x, y, piece)
	if piece == nil {
		return "."
	} else if piece.Rock == c.CurrentRock {
		return "@"
	} else {
		return "#"
	}
}

func (c *Chamber) IsOccupied(x int, y int) bool {
	return c.GetRockPiece(x, y) != nil
}

func (c *Chamber) IsOccupiedByRockOrEmpty(x int, y int, rock *Rock) bool {
	r := c.GetRock(x, y)
	if r == nil {
		return true
	}
	if r == rock {
		return true
	}
	return false
}

func (c *Chamber) GetRock(x int, y int) *Rock {
	piece := c.GetRockPiece(x, y)
	if piece != nil {
		return piece.Rock
	} else {
		return nil
	}
}

// returns the piece of rock (or air)
// func (c *Chamber) GetRockPiece(x int, y int) string {
// 	// rocks might overlap
// 	rocks := make([]*Rock, 0)
// 	for _, rock := range c.Rocks {
// 		// if rock.Occupies(x, y) {
// 		piece := rock.GetPieceChamberXY(x, y, false)
// 		if piece != nil {
// 			rocks = append(rocks, rock)
// 		}
// 		// }
// 	}
// 	if len(rocks) > 1 {
// 		fmt.Println(">>>>>>")
// 		fmt.Printf("GetRockPiece(%v,%v) has %v rocks - this is WRONG.\n", x, y, len(rocks))
// 		for _, r := range rocks {
// 			fmt.Printf("Rock[%v], (%v,%v)\n", r.Number, r.x, r.y)
// 			fmt.Println(r.Debug())
// 			fmt.Println("")
// 			piece := r.GetPieceChamberXY(x, y, true)
// 			fmt.Printf("%v\n", piece)
// 			fmt.Println("")
// 		}
// 		fmt.Println(">>>>>>")
// 	}

// 	for _, rock := range c.Rocks {
// 		// if rock.Occupies(x, y) {
// 		piece := rock.GetPieceChamberXY(x, y, false)
// 		if piece != nil {
// 			if rock == c.CurrentRock {
// 				// fmt.Printf("%v,%v=%v\n", x, y, "@")
// 				return "@"
// 			} else {
// 				// fmt.Printf("%v,%v=%v\n", x, y, "#")
// 				return "#"
// 			}
// 		}
// 		// }
// 	}

// 	return "."
// 	// for each rockpiece, get the piece at that place - air is overridden by solid

// }
