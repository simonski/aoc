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
	Rocks           map[int]*Rock
	CurrentRock     *Rock
	MaxPieceY       int
	Width           int
	Height          int
	PieceCache      map[string]*Piece
	Floor           int
	RockCount       int
}

func NewChamber(input string) *Chamber {
	c := Chamber{Input: input}
	c.ROCK_HORIZONTAL = NewRock("H", "####")
	c.ROCK_PLUS = NewRock("P", ".#.,###,.#.")
	c.ROCK_L = NewRock("L", "..#,..#,###")
	c.ROCK_VERTICAL = NewRock("V", "#,#,#,#")
	c.ROCK_SQUARE = NewRock("SQ", "##,##")
	c.Width = 7
	c.PieceCache = make(map[string]*Piece)
	c.Floor = 0
	c.Rocks = make(map[int]*Rock)
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

func (c *Chamber) AddRock(rock *Rock) {
	// reset the height
	maxY := c.Height

	// fmt.Println(rock.Debug())
	// fmt.Printf("AddRock - initial height was %v.\n", maxY)
	// fmt.Println(rock.Debug())

	c.CurrentRock = rock
	rock.Number = len(c.Rocks) + 1
	c.Rocks[rock.Number] = rock

	// make this rock 2 units space left three units above highest rock (or floor)
	rock.x = 2
	// rock.y = maxY + 3 + (rock.height) - 1
	rock.y = maxY + 3 + (rock.height - 1)
	c.Height = rock.y + 1

	c.AddRockToCache(rock)

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
		if !c.IsOccupiedByRockOrEmpty(x, y, rock) {
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
	for _, piece := range rock.pieces {
		x := rock.x + piece.x + 1
		if x > c.Width {
			return false
		}
		y := rock.y - piece.y
		if !c.IsOccupiedByRockOrEmpty(x, y, rock) {
			return false
		}
	}
	return true
}

func (c *Chamber) CanRockMoveDown(rock *Rock) bool {
	for _, piece := range rock.pieces {
		x := rock.x + piece.x
		y := rock.y - piece.y - 1
		if y < 0 {
			return false
		}
		if !c.IsOccupiedByRockOrEmpty(x, y, rock) {
			return false
		}
	}
	return true
}

func (c *Chamber) MoveLeft(rock *Rock) {
	c.RemoveRockFromCache(rock)
	rock.x -= 1
	c.AddRockToCache(rock)
}

func (c *Chamber) MoveRight(rock *Rock) {
	c.RemoveRockFromCache(rock)
	rock.x += 1
	c.AddRockToCache(rock)
}

func (c *Chamber) MoveDown(rock *Rock) {
	c.RemoveRockFromCache(rock)
	rock.y -= 1
	c.AddRockToCache(rock)
	y := 0
	for _, r := range c.Rocks {
		y = goutils.Max(y, r.y)
	}

	// for k, _ := range c.PieceCache {
	// 	splits := strings.Split(k, "_")
	// 	ycandidate, _ := strconv.Atoi(splits[1])
	// 	y = goutils.Max(y, ycandidate)
	// }

	if y == 0 {
		y = 1
		c.Height = y
	} else {
		c.Height = y + 1
	}

}

func (c *Chamber) RemoveRockFromCache(rock *Rock) {
	remove := make([]string, 0)
	for key := range c.PieceCache {
		piece := c.PieceCache[key]
		if piece != nil {
			if piece.Rock == rock {
				remove = append(remove, key)
			}
		}
	}

	for _, key := range remove {
		delete(c.PieceCache, key)
	}

}

func (c *Chamber) AddRockToCache(rock *Rock) {
	c.AddRockToCache_V2(rock)
}

func (c *Chamber) AddRockToCache_V2(rock *Rock) {
	for _, piece := range rock.pieces {
		x := rock.x + piece.x
		y := rock.y - piece.y
		key := fmt.Sprintf("%v_%v", x, y)
		c.MaxPieceY = goutils.Max(c.MaxPieceY, y)
		c.PieceCache[key] = piece
	}

}

func (c *Chamber) AddRockToCache_V1(rock *Rock) {

	// reset the pieces in the map
	for row := c.Floor; row < rock.height; row++ {
		for col := 0; col < rock.width; col++ {
			new_x := col + rock.x
			new_y := rock.y - row
			key := fmt.Sprintf("%v_%v", new_x, new_y)
			piece := rock.GetPieceAbsoluteXY(col, row)
			if piece != nil {
				if c.PieceCache[key] == nil {
					c.PieceCache[key] = piece
				} else {
					fmt.Println(c.Debug())
					otherRock := c.PieceCache[key].Rock
					fmt.Printf("Rock %v is trying to overwrite a piece from rock %v at position [%v] ", rock.Number, otherRock.Number, key)
					panic("foo")
				}
			}
		}
	}
}

func (c *Chamber) DebugRow(row int) string {
	line := ""
	for col := 0; col < c.Width; col++ {
		piece := c.GetRockPieceString(col, row)
		line = fmt.Sprintf("%v%v", line, piece)
	}
	return line
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
	maxY := goutils.Max(c.MaxPieceY, c.Height)
	for row := maxY - 1; row >= c.Floor; row-- {
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
	fmt.Printf("Rocks=%v, Height=%v\n", c.RockCount, c.Height)
	return line
}

func (c *Chamber) GetRockPiece(x int, y int) *Piece {
	key := fmt.Sprintf("%v_%v", x, y)
	return c.PieceCache[key]
}

func (c *Chamber) GetRockPieceString(x int, y int) string {
	key := fmt.Sprintf("%v_%v", x, y)
	piece := c.PieceCache[key]
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
