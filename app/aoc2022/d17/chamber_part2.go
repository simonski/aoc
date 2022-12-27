package d17

import (
	"fmt"
)

func (c *Chamber) RunPart2(breakAfterRock int, findFloor bool, floorSize int) {
	VERBOSE := c.LOG_LEVEL == 1
	rock := c.NewRock()
	c.AddRock(rock)
	if c.LOG_LEVEL == 1 {
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
		if !c.TickPart2(instruction) {
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

			if c.RockCount%floorSize == 0 {
				max := 1000000000000
				pct := (100 / max) * c.RockCount
				fmt.Printf("%v%% (%v/%v) - %v rocks in cache, %v pieces in cache\n", pct, c.RockCount, max, len(c.Rocks), len(c.PieceCache))
				if VERBOSE {
					fmt.Println(c.Debug())
				}

				if findFloor {
					c.ResetFloor()
				}

			}

		}
	}
}

func (c *Chamber) TickPart2(instruction string) bool {
	VERY_VERBOSE := c.LOG_LEVEL == 2
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
