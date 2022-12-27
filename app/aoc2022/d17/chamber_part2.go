package d17

import (
	"fmt"

	"github.com/simonski/goutils"
)

func (c *Chamber) AddRocks(maxRocks int) []int {
	VERBOSE := c.LOG_LEVEL == 1
	rock := c.NewRock()
	c.AddRock(rock)
	if c.LOG_LEVEL == 1 {
		fmt.Printf("\n[%v] %v\n%v\n", 0, "BEGIN", c.Debug())
	}

	rockHeights := make([]int, 0)
	rockHeights = append(rockHeights, c.Height)
	c.RockCount = 1
	index := -1
	for {
		index++
		if index == len(c.Input) {
			index = 0
		}
		instruction := c.Input[index : index+1]
		if !c.TickPart2(instruction) {

			// we *could* record the height and the rock number - but we'll do that later
			rockHeights = append(rockHeights, c.Height)

			c.CurrentRock = nil
			if c.RockCount >= maxRocks {
				break
			}
			rock := c.NewRock()
			c.AddRock(rock)
			if VERBOSE {
				fmt.Printf("A new rock begins falling\n%v\n", c.Debug())
			}
			c.RockCount++
		}
	}
	return rockHeights
}

func (c *Chamber) Part2_FindFirstKey(rockCount int, keySize int) (int, int) {
	c.AddRocks(rockCount)

	// we want to let it assemble R rocks to H height and then look for repeating patterns.
	cycle := NewCycledetector()

	// now we have N rocks, we can build a buffer and see what we can see
	for row := 0; row < c.Height; row++ {
		line := c.DebugRow(row)
		cycle.Add(line, row, 0)
	}

	// now we have a populated set of data we can perform some querying on.
	// we can detect a cycle by taking N rows from the END then walking forwad looking for repeating cycles
	first_index := -1
	for index := 0; index < cycle.data.Size(); index++ {
		key1 := cycle.Buildkey(index, keySize)
		key2 := cycle.Buildkey(index+keySize, keySize)
		if key1 == key2 {
			first_index = index
			break
		}
	}

	return first_index, keySize

}

func (c *Chamber) Part2_VerifySequences(breakAfterRock int, preambleSize, keySize int) {
	// map of [rockCount]height
	rockHeights := c.AddRocks(breakAfterRock)

	// we want to let it assemble R rocks to H height and then look for repeating patterns.
	cycle := NewCycledetector()
	// now we have N rocks, we can build a buffer and see what we can see
	for row := 0; row < c.Height; row++ {
		line := c.DebugRow(row)
		cycle.Add(line, row, 0)
	}

	// preambleLine := cycle.Buildkey(0, preambleSize)
	key := cycle.Buildkey(preambleSize, keySize)
	key_index := 0
	for index := preambleSize; index <= cycle.data.Size()-keySize; index += keySize {
		this_key := cycle.Buildkey(index, keySize)
		same := key == this_key

		height_low := index
		height_high := index + keySize

		rocks_low, rocks_high := get_rock_count(rockHeights, height_low, height_high)

		fmt.Printf("sequence[%v  %v-%v]: matches %v rocks %v-%v\n", key_index+1, index, index+keySize, same, rocks_low, rocks_high)
		key_index += 1
	}
}

func get_rock_count(rock_heights []int, height_low int, height_high int) (int, int) {
	// calculate the range of rocks in this height
	// results := make([]int, 0)
	low := 100000
	high := 0
	for index, value := range rock_heights {
		if value >= height_low && value <= height_high {
			// results = append(results, index)
			low = goutils.Min(low, index)
			high = goutils.Max(high, index)
		}
		if value > height_high {
			break
		}
	}
	return low, high
}

func (c *Chamber) Part2_FindSequences(breakAfterRock int, maxKeySize int, minKeySize int) {
	// rockHeights := c.AddRocks(breakAfterRock)
	if c.LOG_LEVEL == 1 {
		fmt.Println(c.Debug())
	}

	// we want to let it assemble R rocks to H height and then look for repeating patterns.
	cycle := NewCycledetector()

	// now we have N rocks, we can build a buffer and see what we can see
	for row := 0; row < c.Height; row++ {
		line := c.DebugRow(row)
		cycle.Add(line, row, 0)
	}

	// found_sequence := false
	// now we have a populated set of data we can perform some querying on.
	// we can detect a cycle by taking N rows from the END then walking forwad looking for repeating cycles
	for key_size := maxKeySize; key_size >= minKeySize; key_size-- {

		key_good := false

		// build final key
		length := cycle.data.Size()
		key := cycle.Buildkey(length-key_size, key_size)
		// fmt.Printf("Key size %v is \n\n%v\n\n", key_size, key)
		results := cycle.FindRepeatingKeys(key, key_size)
		// if len(results) > 0 {

		// now we need to see, are the differences the SAME between each repeating key?
		// if they are then this looks like a loop
		difference := 0
		if len(results) > 2 {
			difference = results[1] - results[0]

			// key_good indicates if the difference in position between repeating keys is the SAME for ecah
			// repeat
			key_good = true
			for result_index := 1; result_index+1 < len(results); result_index++ {
				value_1 := results[result_index]
				value_2 := results[result_index+1]
				new_difference := value_2 - value_1
				if new_difference != difference {
					fmt.Printf("The difference between the results varies so this is not a repeating pattern.\n")
					key_good = false
					break
				}
				if result_index+1 == len(results) {
					break
				}
			}
			if key_good {
				fmt.Printf("Key size %v is good, difference is %v.\n", key_size, difference)
			} else {
				fmt.Println("This key is not good.")
			}
			fmt.Printf("There are %v repeating sequences [%v] for key size %v.\n", len(results), results, key_size)
			// found_sequence = true
		} else {
			fmt.Printf("There are no repeating sequences for key size %v\n", key_size)
		}

		// }
	}

	// if found_sequence {
	// 	for index, height := range rockHeights {
	// 		fmt.Printf("%v: %v\n", index, height)
	// 	}
	// }

	// [[72 125 178 231 284 337 390 443 496 549 602 655 708 761 814 867 920 973 1026 1079 1132 1185 1238 1291 1344 1397 1450 1503 1556 1609 1662 1715 1768 1821 1874 1927 1980 2033 2086 2139 2192 2245 2298 2351 2404 2457 2510 2563 2616 2669 2722 2775 2828 2881]]

	// tells me a key size 100 has a repeating pattern where every

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
