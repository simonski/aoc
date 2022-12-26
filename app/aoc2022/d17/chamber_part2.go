package d17

import (
	"fmt"
)

func (c *Chamber) TickPart2(instruction string, VERBOSE bool, VERY_VERBOSE bool) bool {
	if instruction == "<" {
		if c.CanRockMoveLeft(c.CurrentRock) {
			c.MoveLeft(c.CurrentRock)
			if VERY_VERBOSE {
				fmt.Printf("Jet of gas pushes rock left (current rock is %v, (%v,%v):\n", c.CurrentRock.Name, c.CurrentRock.x, c.CurrentRock.y)
				fmt.Println(c.Debug())
			}
		} else {
			if VERY_VERBOSE {
				fmt.Printf("Jet of gas pushes rock left, but nothing happens (current rock is %v, (%v,%v):\n", c.CurrentRock.Name, c.CurrentRock.x, c.CurrentRock.y)
				fmt.Println(c.Debug())
			}
		}
	} else if instruction == ">" {
		if c.CanRockMoveRight(c.CurrentRock) {
			c.MoveRight(c.CurrentRock)
			if VERY_VERBOSE {
				fmt.Printf("Jet of gas pushes rock right (current rock is %v, (%v,%v):\n", c.CurrentRock.Name, c.CurrentRock.x, c.CurrentRock.y)
				fmt.Println(c.Debug())
			}
		} else {
			if VERY_VERBOSE {
				fmt.Printf("Jet of gas pushes rock right, but nothing happens: (current rock is %v, (%v,%v):\n", c.CurrentRock.Name, c.CurrentRock.x, c.CurrentRock.y)
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

func (c *Chamber) RunPart2(VERBOSE bool, VERY_VERBOSE bool, breakAfterRock int) {
	rock := c.NewRock()
	c.AddRock(rock)
	if VERBOSE {
		fmt.Printf("\n[%v] %v\n%v\n", 0, "BEGIN", c.Debug())
	}
	c.RockCount = 1
	index := -1
	for {
		index++

		if index == len(c.Input) {
			index = 0
		}
		instruction := c.Input[index : index+1]
		if !c.TickPart2(instruction, VERBOSE, VERY_VERBOSE) {
			c.CurrentRock = nil
			if c.RockCount >= breakAfterRock {
				return
			}

			rock := c.NewRock()
			c.AddRock(rock)
			if VERBOSE {
				fmt.Printf("A new rock begins falling\n%v\n", c.Debug())
			}
			c.RockCount++

			if c.RockCount%1000 == 0 {
				max := 1000000000000
				pct := (100 / max) * c.RockCount
				fmt.Printf("%v%% (%v/%v) - %v rocks in cache, %v pieces in cache\n", pct, c.RockCount, max, len(c.Rocks), len(c.PieceCache))
				if VERBOSE {
					fmt.Println(c.Debug())
				}

				// go from the "first" rock and work out when we get to a seal
				col1 := false
				col2 := false
				col3 := false
				col4 := false
				col5 := false
				col6 := false
				col7 := false

				closingLine := -1

				// walk up from the bottom finding the first row that "closes" horizontally.

				// see if we can find a line that means we can drop the number of rocks and pieces we have

				for row := c.Height - 1; row > c.Floor; row-- {
					col1 = c.IsOccupied(0, row) || c.IsOccupied(0, row-1)
					col2 = c.IsOccupied(1, row) || c.IsOccupied(1, row-1)
					col3 = c.IsOccupied(2, row) || c.IsOccupied(2, row-1)
					col4 = c.IsOccupied(3, row) || c.IsOccupied(3, row-1)
					col5 = c.IsOccupied(4, row) || c.IsOccupied(4, row-1)
					col6 = c.IsOccupied(5, row) || c.IsOccupied(5, row-1)
					col7 = c.IsOccupied(6, row) || c.IsOccupied(6, row-1)

					if col1 && col2 && col3 && col4 && col5 && col6 && col7 {
						closingLine = row - 1
						break
					}
				}

				for row := c.Floor; row < c.Height; row++ {
					value := c.DebugRow(row)
					if value == "#######" {
						fmt.Printf("row [%v] = %v\n", row, value)
					}
				}

				if closingLine > -1 {

					// if c.RockCount >= 9000 {
					// 	VERBOSE = true
					// 	// VERY_VERBOSE = true
					// }

					oldFloor := c.Floor
					if VERBOSE {
						fmt.Printf("Terminating line is %v, old floor was %v, setting new floor to be %v\n", closingLine, oldFloor, closingLine-1)
						fmt.Printf("Remove all pieces and rocks from the floor to just under this new floor")
					}
					c.Floor = closingLine - 1
					// get rid of anything "lower" than closingLine-1

					for index := oldFloor; index < closingLine; index++ {
						for col := 0; col < 7; col++ {
							key := fmt.Sprintf("%v_%v", col, index)
							// fmt.Printf("Removing piece %v\n", key)
							delete(c.PieceCache, key)
						}
					}

					remove_rock := make([]int, 0)
					for rock_key, rock := range c.Rocks {
						if rock.y < closingLine {
							remove_rock = append(remove_rock, rock_key)
						}
					}
					for _, key := range remove_rock {
						delete(c.Rocks, key)
					}

					if VERBOSE {
						fmt.Printf("After cleaning, graph is\n\n%v\n", c.Debug())
					}
				}
				// os.Exit(1)

			}

		}
	}
}
