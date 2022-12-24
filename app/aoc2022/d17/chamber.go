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
	} else { // if tick == 5 {
		return c.ROCK_SQUARE.Clone()
		// } else {
		// 	return c.ROCK_HORIZONTAL.Clone()
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
	maxY := c.Height()
	// fmt.Printf("AddRock - initial height was %v.\n", maxY)

	// fmt.Printf("Rock is \n\n%v\n\n", rock.Debug())

	c.CurrentRock = rock
	c.Rocks = append(c.Rocks, rock)

	// make this rock 2 units space left three units above highest rock (or floor)
	rock.x = 2
	rock.y = maxY + 3 + (rock.height) - 1
	rock.y = maxY + 2 + (rock.height)
	// fmt.Printf("AddRock - rock x is 2, rock height is %v, rock y will be %v\n", rock.height, rock.y)
	// fmt.Printf("AddRock - c height is now %v, rock x=%v, rock y=%v, rock width=%v, rock height=%v\n", c.Height(), rock.x, rock.y, rock.width, rock.height)
	fmt.Print("\n\nADDROCK\n\n")

}

func (c *Chamber) CanRockMoveLeft(rock *Rock) bool {
	if rock.x == 0 {
		return false // just can't go left anymore
	}
	// otherwise, is it obscured by anything if it goes left?
	for _, piece := range rock.GetLeftmostPieces() {
		x := rock.x - piece.x - 1
		y := rock.y - piece.y
		if c.IsOccupied(x, y) {
			// obscured
			return false
		}
	}
	return true
}

func (c *Chamber) CanRockMoveRight(rock *Rock) bool {
	if rock.x+rock.width == c.Width {
		return false
	}
	// otherwise, is it obscured by anything if it goes left?
	for _, piece := range rock.GetRightmostPieces() {
		x := rock.x + piece.x + 1
		y := rock.y - piece.y
		occupied := c.IsOccupied(x, y)

		// fmt.Printf("There are %v rightmost pieces, this piece is (%v,%v) on rock (%v,%v) which is (%v,%v) on the chamber, occupied=%v\n", len(rock.GetRightmostPieces()), rock.x, rock.y, piece.x, piece.y, x, y, occupied)
		if occupied {
			// obscured
			return false
		}
	}
	return true
}

func (c *Chamber) CanRockMoveDown(rock *Rock) bool {
	if rock.y-1 < 0 {
		return false
	}
	// otherwise, is it obscured by anything if it goes left?
	for _, piece := range rock.GetBottomPieces() {
		x := rock.x + piece.x
		y := rock.y - piece.y - 1
		outcome := c.IsOccupied(x, y)
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
			// fmt.Println("MOVE LEFT")
			c.CurrentRock.x -= 1
		} else {
			// fmt.Println("CANNOT MOVE LEFT")

		}
	} else if instruction == ">" {
		if c.CanRockMoveRight(c.CurrentRock) {
			// fmt.Println("MOVE RIGHT")
			c.CurrentRock.x += 1
		} else {
			// fmt.Println("CANNOT MOVE RIGHT")
		}
	}

	if c.CanRockMoveDown(c.CurrentRock) {
		// fmt.Println("MOVE DOWN")
		c.CurrentRock.y -= 1
		return true
	} else {
		// fmt.Println("CANNOT MOVE DOWN, COME TO REST AND NEW ROCK WILL BE REQUIRED.")
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
			// if VERBOSE {
			// 	fmt.Println(c.Debug())
			// }
			if rockCount+1 > breakAfterRock {
				if VERBOSE {
					fmt.Println(c.Debug())
				}
				fmt.Printf("After %v rocks, size is %v\n", rockCount, c.Height())
				break
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
	line := ""
	for row := c.Height(); row >= 0; row-- {
		line = fmt.Sprintf("%v|", line)
		for col := 0; col < c.Width; col++ {
			piece := c.GetRockPiece(col, row)
			line = fmt.Sprintf("%v%v", line, piece)
		}
		line = fmt.Sprintf("%v|\n", line)
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
	for _, rock := range c.Rocks {
		if rock.Occupies(x, y) {
			piece := rock.GetPiece(x, y)
			if piece != nil && piece.solid {
				if rock == c.CurrentRock {
					return "@"
				} else {
					return "#"
				}
			}
		}
	}
	return "."
	// for each rockpiece, get the piece at that place - air is overridden by solid

}
