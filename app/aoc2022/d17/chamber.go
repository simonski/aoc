package d17

import (
	"fmt"
)

type Chamber struct {
	rockTick        int
	Input           string
	ROCK_HORIZONTAL *Rock
	ROCK_VERTICAL   *Rock
	ROCK_PLUS       *Rock
	ROCK_L          *Rock
	ROCK_SQUARE     *Rock
	RocksXY         map[string]*Rock
	Rocks           []*Rock
	CurrentRock     *Rock
	Width           int
	Pieces          map[string]*Piece
}

func NewChamber(input string) *Chamber {
	c := Chamber{Input: input}
	c.ROCK_HORIZONTAL = NewRock("####")
	c.ROCK_PLUS = NewRock(".#.,###,.#.")
	c.ROCK_L = NewRock("..#,..#,###")
	c.ROCK_VERTICAL = NewRock("#,#,#,#")
	c.ROCK_SQUARE = NewRock("##,##")
	c.Width = 7
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

func (c *Chamber) Height() int {
	if len(c.Rocks) == 0 {
		return 0
	}
	y := 0
	for _, rock := range c.Rocks {
		if rock.y > y {
			y = rock.y
		}
	}
	return y + 1
}

func (c *Chamber) AddRock(rock *Rock) {
	// reset the height
	maxY := c.Height() - 1

	// fmt.Println(rock.Debug())
	// fmt.Printf("AddRock - initial height was %v.\n", maxY)
	// fmt.Println(rock.Debug())

	c.CurrentRock = rock
	c.Rocks = append(c.Rocks, rock)
	rock.Number = len(c.Rocks)

	// make this rock 2 units space left three units above highest rock (or floor)
	rock.x = 2
	// rock.y = maxY + 3 + (rock.height) - 1
	rock.y = maxY + 3 + (rock.height)
	// rock.y =
	// fmt.Printf("AddRock - rock x,y is (%v,%v), width=%v, height=%v\n", rock.x, rock.y, rock.width, rock.height)
	// fmt.Printf("AddRock - c height is now %v, rock x=%v, rock y=%v, rock width=%v, rock height=%v\n", c.Height(), rock.x, rock.y, rock.width, rock.height)

}

func (c *Chamber) CanRockMoveLeft(rock *Rock) bool {
	if rock.x == 0 {
		return false // just can't go left anymore
	}
	// otherwise, is it obscured by anything if it goes left?
	for _, piece := range rock.GetLeftmostPieces(c) {
		x := rock.x - piece.x - 1
		y := rock.y - piece.y
		if c.IsOccupied(x, y) || rock.GetPieceChamberXY(x, y, false) != nil {
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
	for _, piece := range rock.GetRightmostPieces(c) {
		x := rock.x + piece.x + 1
		if x > c.Width {
			return false
		}
		y := rock.y - piece.y
		occupied := c.IsOccupied(x, y) || rock.GetPieceChamberXY(x, y, false) != nil

		// fmt.Printf("There are %v rightmost pieces, this piece is (%v,%v) on rock (%v,%v) which is (%v,%v) on the chamber, occupied=%v\n", len(rock.GetRightmostPieces()), rock.x, rock.y, piece.x, piece.y, x, y, occupied)
		if occupied {
			// obscured
			return false
		}
	}
	return true
}

func (c *Chamber) CanRockMoveDown(rock *Rock) bool {
	// if rock.y-1 < 0 {
	// 	return false
	// }
	// otherwise, is it obscured by anything if it goes left?

	for _, piece := range rock.GetBottomPieces(c) {
		x := rock.x + piece.x
		y := rock.y - piece.y - 1
		if y < 0 {
			return false
		}
		outcome := c.IsOccupied(x, y) || rock.GetPieceChamberXY(x, y, false) != nil
		// fmt.Printf("bottom piece (%v,%v) which is (%v,%v) in the chamber, outcome=%v\n", piece.x, piece.y, x, y, outcome)
		if outcome {
			// obscured
			return false
		}
	}
	return true
}

// move the rock, or if it cannot move, return false
func (c *Chamber) Tick(instruction string) bool {
	if instruction == "<" {
		if c.CanRockMoveLeft(c.CurrentRock) {
			c.CurrentRock.x -= 1
		}
	} else if instruction == ">" {
		if c.CanRockMoveRight(c.CurrentRock) {
			c.CurrentRock.x += 1
		}
	}

	if c.CanRockMoveDown(c.CurrentRock) {
		c.CurrentRock.y -= 1
		return true
	}

	return false
}

func (c *Chamber) Run(VERBOSE bool, breakAfterRock int) {
	rock := c.NewRock()
	c.AddRock(rock)
	if VERBOSE {
		fmt.Printf("\n[%v] %v\n%v\n", 0, "BEGIN", c.Debug())
	}
	rockCount := 1
	index := -1
	for {
		index++
		if index == len(c.Input) {
			index = 0
		}
		instruction := c.Input[index : index+1]
		if !c.Tick(instruction) {
			c.CurrentRock = nil
			if rockCount >= breakAfterRock {
				if VERBOSE {
					fmt.Println(c.Debug())
				}
				fmt.Printf("After %v rocks, size is %v\n", rockCount, c.Height())
				return
			}

			rock := c.NewRock()
			c.AddRock(rock)

			if VERBOSE {
				fmt.Println(c.Debug())
			}
			rockCount++
		}
		// if VERBOSE {
		// 	fmt.Printf("[%v] %v\n%v\n", index, instruction, c.Debug())
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
	for row := c.Height() - 1; row >= 0; row-- {
		line = fmt.Sprintf("%v|", line)
		for col := 0; col < c.Width; col++ {
			piece := c.GetRockPiece(col, row)
			line = fmt.Sprintf("%v%v", line, piece)
		}
		line = fmt.Sprintf("%v|", line)
		line = fmt.Sprintf("%v     %v\n", line, row)
	}
	line = fmt.Sprintf("%v+-------+", line)
	fmt.Printf("Rocks=%v, Height=%v\n", len(c.Rocks), c.Height())
	return line
}

func (c *Chamber) IsOccupied(x int, y int) bool {
	return c.GetRockPiece(x, y) != "."
}

// returns the piece of rock (or air)
func (c *Chamber) GetRockPiece(x int, y int) string {
	// rocks might overlap
	rocks := make([]*Rock, 0)
	for _, rock := range c.Rocks {
		// if rock.Occupies(x, y) {
		piece := rock.GetPieceChamberXY(x, y, false)
		if piece != nil {
			rocks = append(rocks, rock)
		}
		// }
	}
	if len(rocks) > 1 {
		fmt.Println(">>>>>>")
		fmt.Printf("GetRockPiece(%v,%v) has %v rocks - this is WRONG.\n", x, y, len(rocks))
		for _, r := range rocks {
			fmt.Printf("Rock[%v], (%v,%v)\n", r.Number, r.x, r.y)
			fmt.Println(r.Debug())
			fmt.Println("")
			piece := r.GetPieceChamberXY(x, y, true)
			fmt.Printf("%v\n", piece)
			fmt.Println("")
		}
		fmt.Println(">>>>>>")
	}

	for _, rock := range c.Rocks {
		// if rock.Occupies(x, y) {
		piece := rock.GetPieceChamberXY(x, y, false)
		if piece != nil {
			if rock == c.CurrentRock {
				// fmt.Printf("%v,%v=%v\n", x, y, "@")
				return "@"
			} else {
				// fmt.Printf("%v,%v=%v\n", x, y, "#")
				return "#"
			}
		}
		// }
	}

	return "."
	// for each rockpiece, get the piece at that place - air is overridden by solid

}
