package d13

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/simonski/aoc/utils"
// )

// /*
// Day 13: Claw Contraption
// */

// type Puzzle struct {
// 	title string
// 	year  int
// 	day   int
// 	input string
// 	lines []string

// 	games []*Game
// }

// type Button struct {
// 	cost int
// 	x    int
// 	y    int
// }

// type Prize struct {
// 	x int
// 	y int
// }

// type Game struct {
// 	buttonA *Button
// 	buttonB *Button
// 	prize   *Prize
// }

// // Button A: X+26, Y+66
// // Button B: X+67, Y+21
// // Prize: X=12748, Y=12176

// func NewGame(lines []string) *Game {
// 	a := NewButton(lines[0], 3)
// 	b := NewButton(lines[1], 1)
// 	prize := NewPrize(lines[2])
// 	return &Game{buttonA: a, buttonB: b, prize: prize}
// }

// func NewButton(line string, cost int) *Button {
// 	line = strings.ReplaceAll(line, "Button A: ", "")
// 	line = strings.ReplaceAll(line, "Button B: ", "")
// 	line = strings.ReplaceAll(line, "X+", "")
// 	line = strings.ReplaceAll(line, "Y+", "")
// 	line = strings.ReplaceAll(line, " ", "")
// 	splits := utils.SplitDataToListOfInts(line, ",")
// 	b := Button{cost: cost, x: splits[0], y: splits[1]}
// 	return &b
// }

// func NewPrize(line string) *Prize {
// 	line = strings.ReplaceAll(line, "Prize: ", "")
// 	line = strings.ReplaceAll(line, "X=", "")
// 	line = strings.ReplaceAll(line, "Y=", "")
// 	line = strings.ReplaceAll(line, " ", "")
// 	splits := utils.SplitDataToListOfInts(line, ",")
// 	p := Prize{x: splits[0], y: splits[1]}
// 	return &p
// }

// func (puzzle *Puzzle) GetSummary() utils.Summary {
// 	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
// 	s.ProgressP1 = utils.NotStarted
// 	s.ProgressP2 = utils.NotStarted
// 	s.DateStarted = "2024-12-14 09:19:41"
// 	return s
// }

// func NewPuzzleWithData(input string) *Puzzle {
// 	iyear, _ := strconv.Atoi("2024")
// 	iday, _ := strconv.Atoi("13")
// 	p := Puzzle{year: iyear, day: iday, title: "Day 13: Claw Contraption"}
// 	p.Load(input)
// 	return &p
// }

// func NewPuzzle() *Puzzle {
// 	return NewPuzzleWithData(REAL_DATA)
// }

// func (puzzle *Puzzle) Load(input string) {
// 	lines := strings.Split(input, "\n")
// 	puzzle.input = input
// 	puzzle.lines = lines

// 	games := make([]*Game, 0)
// 	for index := 0; index < len(lines); index++ {
// 		g := make([]string, 0)
// 		g = append(g, lines[index])
// 		g = append(g, lines[index+1])
// 		g = append(g, lines[index+2])
// 		game := NewGame(g)
// 		index += 3
// 		games = append(games, game)
// 	}
// 	puzzle.games = games
// }

// func (puzzle *Puzzle) tryX(button1 *Button, button2 *Button, game *Game) (bool, int, int, int, int, int) {

// 	xRemainder := game.prize.x % button1.x
// 	maxAPresses := (game.prize.x - xRemainder) / button1.x
// 	presses1 := 0
// 	presses2 := 0
// 	// cost1 := 0
// 	// cost2 := 0

// 	// maxAPresses is the highest number of A presses that yields a remainder that can be used.
// 	// this yields a remainder which we can see if the B presses fill in
// 	// if it does, then add it to the candidates
// 	// once we have all the candidates, we can work out the best cost
// 	candidates := make([][]int, 0)

// 	for index := maxAPresses; index > 0; index-- {
// 		presses1 = index
// 		button1_x := presses1 * button1.x
// 		button1_y := presses1 * button1.y
// 		xRemainder := game.prize.x - button1_x
// 		yRemainder := game.prize.y - button1_y
// 		if xRemainder%button2.x == 0 && yRemainder%button2.y == 0 {
// 			presses2 = xRemainder / button2.x
// 			pair := make([]int, 0)
// 			pair = append(pair, presses1)
// 			pair = append(pair, presses2)
// 			candidates = append(candidates, pair)
// 		}
// 	}

// 	if len(candidates) == 0 {
// 		return false, -1, -1, -1, -1, -1
// 	}

// 	minCost := 10000000000000000
// 	pressesA := 0
// 	pressesB := 0
// 	costA := 0
// 	costB := 0
// 	for _, candidate := range candidates {
// 		aCost := button1.cost * candidate[0]
// 		bCost := button2.cost * candidate[1]
// 		if aCost+bCost < minCost {
// 			minCost = aCost + bCost
// 			pressesA = candidate[0]
// 			pressesB = candidate[1]
// 			costA = aCost
// 			costB = bCost
// 		}
// 	}
// 	return true, pressesA, pressesB, costA, costB, minCost
// }

// func (puzzle *Puzzle) tryY(button1 *Button, button2 *Button, game *Game) (bool, int, int, int, int, int) {

// 	yRemainder := game.prize.y % button1.y
// 	maxAPresses := (game.prize.y - yRemainder) / button1.y
// 	presses1 := 0
// 	presses2 := 0
// 	// cost1 := 0
// 	// cost2 := 0

// 	// maxAPresses is the highest number of A presses that yields a remainder that can be used.
// 	// this yields a remainder which we can see if the B presses fill in
// 	// if it does, then add it to the candidates
// 	// once we have all the candidates, we can work out the best cost
// 	candidates := make([][]int, 0)

// 	for index := maxAPresses; index > 0; index-- {
// 		presses1 = index
// 		button1_x := presses1 * button1.x
// 		button1_y := presses1 * button1.y
// 		xRemainder := game.prize.x - button1_x
// 		yRemainder := game.prize.y - button1_y
// 		if xRemainder%button2.x == 0 && yRemainder%button2.y == 0 {
// 			presses2 = xRemainder / button2.x
// 			pair := make([]int, 0)
// 			pair = append(pair, presses1)
// 			pair = append(pair, presses2)
// 			candidates = append(candidates, pair)
// 		}
// 	}

// 	minCost := 10000000000000000
// 	if len(candidates) == 0 {
// 		return false, -1, -1, -1, -1, minCost
// 	}

// 	pressesA := 0
// 	pressesB := 0
// 	costA := 0
// 	costB := 0

// 	cost := minCost
// 	for _, candidate := range candidates {
// 		aCost := button1.cost * candidate[0]
// 		bCost := button2.cost * candidate[1]
// 		if aCost+bCost < minCost {
// 			minCost = aCost + bCost
// 			pressesA = candidate[0]
// 			pressesB = candidate[1]
// 			costA = aCost
// 			costB = bCost
// 		}
// 	}
// 	cost = costA + costB
// 	return true, pressesA, pressesB, costA, costB, cost
// }

// func (puzzle *Puzzle) Part1() {
// 	puzzle.Load(REAL_DATA)
// 	// puzzle.Load(TEST_DATA)

// 	// 115077 too high
// 	//  93685 too high
// 	//  86368 NOPE
// 	//  75478 NOPE
// 	//  63457 NOPE
// 	//  79288 NOPE
// 	//  68787 NOPE

// 	fmt.Printf("There are %v games.\n", len(puzzle.games))

// 	total := 0
// 	for gindex, game := range puzzle.games {
// 		fmt.Println()
// 		fmt.Printf("[%v], Button A: X+%v, Y+%v, cost=%v\n", gindex, game.buttonA.x, game.buttonA.y, game.buttonA.cost)
// 		fmt.Printf("[%v], Button B: X+%v, Y+%v, cost=%v\n", gindex, game.buttonB.x, game.buttonB.y, game.buttonB.cost)
// 		fmt.Printf("[%v], Prize X=%v, Y=%v\n", gindex, game.prize.x, game.prize.y)

// 		found1, pressesA1, pressesB1, costA1, costB1, cost1 := puzzle.tryX(game.buttonA, game.buttonB, game)
// 		found2, pressesB2, pressesA2, costB2, costA2, cost2 := puzzle.tryY(game.buttonA, game.buttonB, game)

// 		found3, _, _, _, _, cost3 := puzzle.tryX(game.buttonB, game.buttonA, game)
// 		found4, _, _, _, _, cost4 := puzzle.tryY(game.buttonB, game.buttonA, game)

// 		var found bool
// 		var pressesA int
// 		var pressesB int
// 		var costA int
// 		var costB int
// 		cost := 100000000000000000

// 		if found1 && found2 && cost1 < cost2 {
// 			found = found1
// 			pressesA = pressesA1
// 			pressesB = pressesB1
// 			costA = costA1
// 			costB = costB1
// 			cost = cost1
// 		} else if found1 && found2 && cost1 > cost2 {
// 			found = found2
// 			pressesA = pressesA2
// 			pressesB = pressesB2
// 			costA = costA2
// 			costB = costB2
// 			cost = cost2
// 		} else if found1 {
// 			found = found1
// 			pressesA = pressesA1
// 			pressesB = pressesB1
// 			costA = costA1
// 			costB = costB1
// 			cost = cost1
// 		} else if found2 {
// 			found = found2
// 			pressesA = pressesA2
// 			pressesB = pressesB2
// 			costA = costA2
// 			costB = costB2
// 			cost = cost2
// 		} else {
// 			found = false
// 		}

// 		if found1 || found2 || found3 || found4 {
// 			if cost1 > 0 {
// 				cost = utils.MinInt(cost, cost1)
// 			}
// 			if cost2 > 0 {
// 				cost = utils.MinInt(cost, cost2)
// 			}
// 			if cost3 > 0 {
// 				cost = utils.MinInt(cost, cost3)
// 			}
// 			if cost4 > 0 {
// 				cost = utils.MinInt(cost, cost4)
// 			}
// 		}

// 		// found2, aPresses2, bPresses2, costa2, costb2 := puzzle.tryY(game.buttonA, game.buttonB, game)

// 		// costa1 = aPresses1 * game.buttonA.cost
// 		// costb1 = bPresses1 * game.buttonB.cost

// 		// costa2 = aPresses2 * game.buttonA.cost
// 		// costb2 = bPresses2 * game.buttonB.cost

// 		// total := costa1 + costb1
// 		// total2 := costa2 + costb2

// 		// a_prefix := " "
// 		// b_prefix := " "
// 		// if found1 && found2 {
// 		// 	if total1 <= total2 {
// 		// 		a_prefix = "*"
// 		// 		total += total1
// 		// 	} else {
// 		// 		b_prefix = "*"
// 		// 		total += total2
// 		// 	}
// 		// 	fmt.Printf("Works (A%v), A presses=%v, B presses=%v, costA=%v, costB=%v, total=%v\n", a_prefix, aPresses1, bPresses1, costa1, costb1, total1)
// 		// 	fmt.Printf("Works (B%v), A presses=%v, B presses=%v, costA=%v, costB=%v, total=%v\n", b_prefix, aPresses2, bPresses2, costa2, costb2, total2)
// 		// } else if found1 {
// 		// 	a_prefix = "*"
// 		// 	fmt.Printf("Works (A%v), A presses=%v, B presses=%v, costA=%v, costB=%v, total=%v\n", a_prefix, aPresses1, bPresses1, costa1, costb1, total1)
// 		// 	fmt.Printf("Works (B%v), A presses=%v, B presses=%v, costA=%v, costB=%v, total=%v\n", b_prefix, aPresses2, bPresses2, costa2, costb2, total2)
// 		// 	total += total1
// 		if found {
// 			total += cost
// 			fmt.Printf("Works A presses=%v, B presses=%v, costA=%v, costB=%v, cost=%v, total=%v\n", pressesA, pressesB, costA, costB, cost, total)
// 			// total += total2
// 		} else {
// 			fmt.Println("Not found.")
// 		}

// 	}
// 	fmt.Printf("total=%v\n", total)

// 	// aPresses := (game.prize.x - xRemainder) / game.buttonA.x
// 	// aDiff := aPresses * game.buttonA.x
// 	// yDiff := aPresses * game.buttonA.y

// 	// aPresses := (game.prize.y - yRemainder) / game.buttonA.y

// 	// bPresses := xRemainder / game.buttonB.x
// 	// bPresses := yRemainder / game.buttonB.y

// 	// aCost := aPresses * game.buttonA.cost
// 	// bCost := bPresses * game.buttonB.cost
// 	// cost := aCost + bCost

// 	// // using ButtonA first
// 	// xRemainder := game.prize.x % game.buttonB.x
// 	// yRemainder := game.prize.y % game.buttonB.y

// 	// bXPresses := (game.prize.x - xRemainder) / game.buttonB.x
// 	// bYPresses := (game.prize.y - yRemainder) / game.buttonB.y

// 	// aXPresses := xRemainder / game.buttonA.x
// 	// aYPresses := yRemainder / game.buttonA.y
// }

// func (puzzle *Puzzle) Part2() {
// 	puzzle.Load(REAL_DATA)
// }

// func (puzzle *Puzzle) Run() {
// 	puzzle.Part1()
// 	puzzle.Part2()
// }
