package main

import "fmt"

/*
 */

type Ring struct {
	Current *Cup
	Max     *Cup
	Min     *Cup
	Size    int
	Cache   map[int]*Cup
}

func (r *Ring) Play(rounds int, DEBUG bool) {

	// if DEBUG {
	// 	fmt.Printf("First: %v, First.Next %v, First.Prev %v\n", firstCup.Value, firstCup.Next.Value, firstCup.Prev.Value)
	// 	fmt.Printf("Last: %v, Last.Next %v, Last.Prev %v\n", lastCup.Value, lastCup.Next.Value, lastCup.Prev.Value)
	// 	fmt.Printf("Min: %v, Min.Next %v, Min.Prev %v\n", minCup.Value, minCup.Next.Value, minCup.Prev.Value)
	// 	fmt.Printf("Max: %v, Max.Next %v, Max.Prev %v\n", maxCup.Value, maxCup.Next.Value, maxCup.Prev.Value)
	// }
	// TRACE := false
	for round := 0; round < rounds; round++ {
		// if round%1000 == 0 {
		// 	fmt.Printf("%v/%v\n", round, rounds)
		// }
		// if DEBUG {
		// 	fmt.Printf("%v\n", r.Debug(round))
		// }
		current := r.Current
		// if DEBUG {
		// 	fmt.Printf("-- move %v -- Current=%v, Current.Prev=%v, Current.Next=%v\n", round+1, current.Value, current.Prev.Value, current.Next.Value)
		// 	// if TRACE {
		// 	// 	line := r.Debug(round)
		// 	// 	fmt.Printf("cups: %v\n", line)
		// 	// }
		// }

		c1 := current.Next.Detach()
		c2 := current.Next.Detach()
		c3 := current.Next.Detach()
		// if TRACE {
		// 	fmt.Printf("-- move %v -- Current=%v, (detached %v,%v,%v) new Current.Prev=%v, new Current.Next=%v\n", round+1, current.Value, c1.Value, c2.Value, c3.Value, current.Prev.Value, current.Next.Value)
		// }

		lookFor := current.Value - 1
		// fmt.Printf("Current=%v, lookFor=%v, after detach %v,%v,%v Current.Next = %v.%v\n", current.Value, lookFor, c1.Value, c2.Value, c3.Value, current.Value, current.Next.Value)
		//
		for {
			// if TRACE {
			// 	fmt.Printf("Look for %v\n", lookFor)
			// }
			if lookFor == c1.Value || lookFor == c2.Value || lookFor == c3.Value {
				// if TRACE {
				// 	fmt.Printf("Lookfor %v subtracting 1 as it is one of the values\n", lookFor)
				// }
				lookFor -= 1
			} else {
				if lookFor < r.Min.Value {
					// if TRACE {
					// 	fmt.Printf("Look for %v too small setting to max value %v\n", lookFor, r.Max.Value)
					// }
					// fmt.Printf("Lookfor is < min vale, make it max")
					lookFor = r.Max.Value
					continue
				}
				// if TRACE {
				// 	fmt.Printf("Look for %v found a good value\n", lookFor)
				// }
				// fmt.Printf("Look for is NOT one of the values,it's fine.")
				break
			}
		}
		// fmt.Printf("Look for is %v\n", lookFor)
		destinationCup := r.Current
		if lookFor == r.Min.Value {
			destinationCup = r.Min
		} else if lookFor == r.Max.Value {
			destinationCup = r.Max
		} else {
			// if TRACE {
			// 	fmt.Printf("lookFor is %v, searching for cup.\n", lookFor)
			// }
			destinationCup = r.Find(lookFor)
			// if TRACE {
			// 	fmt.Printf("lookFor is %v, searching for cup, found %v.\n", lookFor, destinationCup.Value)
			// }
		}

		// if TRACE {
		// 	fmt.Printf("-- move %v -- Current=%v, Destination=%v  destination.Prev=%v, destination.Next=%v\n", round+1, current.Value, destinationCup.Value, destinationCup.Prev.Value, destinationCup.Next.Value)
		// }

		// fmt.Printf("Destination Cup (%v) going to attach right %v\n", destinationCup.Value, c1.Value)
		destinationCup.AttachRight(c1)
		// line1 := r.Debug(round)
		// fmt.Printf("destination attach right %v.Next=%v\n", destinationCup.Value, c1.Value)
		c1.AttachRight(c2)
		// line2 := r.Debug(round)
		// fmt.Printf("c1 attach c2 right %v.Next=%v\n", c1.Value, c2.Value)
		c2.AttachRight(c3)
		// line3 := r.Debug(round)
		// fmt.Printf("c2 attach c3 right %v.Next=%v\n", c2.Value, c3.Value)
		// if DEBUG {
		// 	// line := r.Debug(round)
		// 	// fmt.Printf("-- move %v -- \n", round+1)
		// 	// fmt.Printf("cups: %v\n", line)
		// 	fmt.Printf("pick up: %v, %v, %v\n", c1.Value, c2.Value, c3.Value)
		// 	fmt.Printf("destination: %v\n", destinationCup.Value)
		// 	// if TRACE {
		// 	// 	fmt.Printf("After Attach c1 (%v): '%v'\n", c1.Value, line1)
		// 	// 	fmt.Printf("After Attach c2 (%v): '%v'\n", c2.Value, line2)
		// 	// 	fmt.Printf("After Attach c3 (%v): '%v'\n", c3.Value, line3)
		// 	// }
		// 	fmt.Printf("\n")

		// }
		r.Current = current.Next
	}
	// fmt.Printf(r.Debug(rounds))
}

// Debug unrolls the Ring starting at Current - round
func (r *Ring) Debug(round int) string {
	start := r.Current
	for index := 0; index < round; index++ {
		start = start.Prev
	}
	// now walk forward printing out the cups until we loop
	line := ""
	looped := false
	current := start
	for {
		if current == start && looped {
			break
		} else {
			looped = true
		}
		if current == r.Current {
			line += fmt.Sprintf("(%v) ", current.Value)
		} else {
			line += fmt.Sprintf("%v ", current.Value)
		}
		current = current.Next
	}
	return line
}

func (r *Ring) Find(value int) *Cup {

	result, exists := r.Cache[value]
	if exists {
		return result
	}

	candidate := r.Current
	for {
		if candidate.Value != value {
			candidate = candidate.Next
		} else {
			break
		}
	}

	r.Cache[value] = candidate
	return candidate
}

type Cup struct {
	Value int
	Next  *Cup
	Prev  *Cup
}

// Detach removes the cup from the Ring of Cups, returning itself
func (c *Cup) Detach() *Cup {
	// fmt.Printf("%v.Detach() left is %v, right is %v\n", c.Value, c.Prev.Value, c.Next.Value)
	next := c.Next
	prev := c.Prev
	next.Prev = prev
	prev.Next = next
	c.Next = nil
	c.Prev = nil
	return c
}

// Attaches / Inserts the passed cup to the right on this cup
func (c *Cup) AttachRight(toAttach *Cup) {
	// fmt.Printf("%v.AttachRight(%v)\n", c.Value, toAttach.Value)
	next := c.Next
	c.Next = toAttach
	toAttach.Prev = c
	toAttach.Next = next
	next.Prev = toAttach
}

func NewRing(data []int) *Ring {
	minCup := &Cup{Value: len(data)}
	maxCup := &Cup{Value: 0}
	lastCup := &Cup{}
	firstCup := &Cup{Value: -1}
	cache := make(map[int]*Cup)
	for index, v := range data {
		cup := &Cup{Value: v}
		if cup.Value < minCup.Value {
			minCup = cup
		}
		if cup.Value > maxCup.Value {
			maxCup = cup
		}
		lastCup.Next = cup
		cup.Prev = lastCup
		lastCup = cup
		if index == 0 {
			firstCup = cup
		}
		cache[cup.Value] = cup
	}

	// close the loop
	lastCup.Next = firstCup
	firstCup.Prev = lastCup

	r := Ring{Current: firstCup, Min: minCup, Max: maxCup, Size: len(data), Cache: cache}

	return &r
}
