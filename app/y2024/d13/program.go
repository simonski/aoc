package d13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 13: Claw Contraption
*/
type Puzzle struct {
	title string
	year  int
	day   int
	input string
	lines []string
	games []*Game
}

type Button struct {
	cost int
	x    int
	y    int
}

type Prize struct {
	x int
	y int
}

type Game struct {
	buttonA *Button
	buttonB *Button
	prize   *Prize
	lines   []string
}

func (g *Game) debug() string {
	l := fmt.Sprintf("%v\n%v\n%v", g.lines[0], g.lines[1], g.lines[2])
	l = fmt.Sprintf("%v\nButtonA x,y=%v,%v, cost=%v, ButtonB x,y=%v,%v, cost=%v", l, g.buttonA.x, g.buttonA.y, g.buttonA.cost, g.buttonB.x, g.buttonB.y, g.buttonB.cost)
	return l
}

// 38839
func (g *Game) Play1() *GameAttempt {

	// the remainder to be filled by buttonB
	remainderX := g.prize.x % g.buttonA.x
	totalX := g.prize.x - remainderX
	maxPressesA := totalX / g.buttonA.x

	// the remainder to be filled by buttonB
	remainderX2 := g.prize.x % g.buttonB.x
	totalX2 := g.prize.x - remainderX2
	maxPressesB := totalX2 / g.buttonB.x

	var attempt *GameAttempt
	for aPresses := maxPressesA; aPresses > 0; aPresses-- {
		for bPresses := maxPressesB; bPresses > 0; bPresses-- {
			ga := NewGameAttempt(g, aPresses, bPresses)
			if ga.isValid() {
				if attempt == nil {
					attempt = ga
				} else {
					if attempt.totalCost() > ga.totalCost() {
						attempt = ga
					}
				}
			}
		}
	}
	return attempt

}

func (g *Game) Play2() *GameAttempt {

	// Button A: X+26, Y+66
	// Button B: X+67, Y+21
	// Prize: X=10000000012748, Y=10000000012176

	// lcm = x1*x2
	// aX = lcm/a.x

	// the remainder to be filled by buttonB

	// lcm := g.buttonA.x * g.buttonB.x // 1742
	// remainder := g.prize.x % lcm     // 1578

	// // then I want to find an lcm which is a button x and the remainder
	// lcm2 := remainder * g.buttonA.x

	// lcm2 := remainder * lcm
	// remainder2 = g.prize.x % lcm2

	// // if I keep adding the lcm to the remainder until I get a value
	// // that yields mod 0 on either buttonA or buttonB

	// // aPresses is the minimum number of presses before we can flip to B presses
	// // so we can say the answer is
	// // (totalX - aPresses) / bSize = bPresses

	// // remainderX := g.prize.x % g.buttonA.x
	// // totalX := g.prize.x - remainderX
	// // maxPressesA := totalX / g.buttonA.x

	// // // the remainder to be filled by buttonB
	// // remainderX2 := g.prize.x % g.buttonB.x
	// // totalX2 := g.prize.x - remainderX2
	// // maxPressesB := totalX2 / g.buttonB.x

	// var attempt *GameAttempt
	// for aPresses := aX; aPresses < maxPresses; aPresses += aX {
	// 	for bPresses := maxPressesB; bPresses > 0; bPresses-- {
	// 		ga := NewGameAttempt(g, aPresses, bPresses)
	// 		if ga.isValid() {
	// 			if attempt == nil {
	// 				attempt = ga
	// 			} else {
	// 				if attempt.totalCost() > ga.totalCost() {
	// 					attempt = ga
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	return nil

}

type GameAttempt struct {
	game     *Game
	aPresses int
	bPresses int
}

func NewGameAttempt(game *Game, aPresses int, bPresses int) *GameAttempt {
	return &GameAttempt{game: game, aPresses: aPresses, bPresses: bPresses}
}

func (ga *GameAttempt) debug() string {
	result := ga.game.debug()
	result = fmt.Sprintf("%v\nIsValid=%v, A Presses=%v, B Presses=%v, (x=%v, y=%v), cost=(%v + %v)=%v", result, ga.isValid(), ga.aPresses, ga.bPresses, ga.X(), ga.Y(), ga.aCost(), ga.bCost(), ga.totalCost())
	return result
}

func (ga *GameAttempt) isValid() bool {
	return ga.X() == ga.game.prize.x && ga.Y() == ga.game.prize.y
}

func (ga *GameAttempt) totalCost() int {
	return (ga.aPresses * ga.game.buttonA.cost) + (ga.bPresses * ga.game.buttonB.cost)
}

func (ga *GameAttempt) aCost() int {
	return ga.aPresses * ga.game.buttonA.cost
}

func (ga *GameAttempt) bCost() int {
	return ga.bPresses * ga.game.buttonB.cost
}

func (ga *GameAttempt) X() int {
	return (ga.aPresses * ga.game.buttonA.x) + (ga.bPresses * ga.game.buttonB.x)
}

func (ga *GameAttempt) Y() int {
	return (ga.aPresses * ga.game.buttonA.y) + (ga.bPresses * ga.game.buttonB.y)
}

func (ga *GameAttempt) IsValid() bool {
	return ga.X() == ga.game.prize.x && ga.Y() == ga.game.prize.y
}

// Button A: X+26, Y+66
// Button B: X+67, Y+21
// Prize: X=12748, Y=12176

func NewGame(lines []string) *Game {
	a := NewButton(lines[0], 3)
	b := NewButton(lines[1], 1)
	prize := NewPrize(lines[2])
	return &Game{buttonA: a, buttonB: b, prize: prize, lines: lines}
}

func NewButton(line string, cost int) *Button {
	line = strings.ReplaceAll(line, "Button A: ", "")
	line = strings.ReplaceAll(line, "Button B: ", "")
	line = strings.ReplaceAll(line, "X+", "")
	line = strings.ReplaceAll(line, "Y+", "")
	line = strings.ReplaceAll(line, " ", "")
	splits := utils.SplitDataToListOfInts(line, ",")
	b := Button{cost: cost, x: splits[0], y: splits[1]}
	return &b
}

func NewPrize(line string) *Prize {
	line = strings.ReplaceAll(line, "Prize: ", "")
	line = strings.ReplaceAll(line, "X=", "")
	line = strings.ReplaceAll(line, "Y=", "")
	line = strings.ReplaceAll(line, " ", "")
	splits := utils.SplitDataToListOfInts(line, ",")
	p := Prize{x: splits[0], y: splits[1]}
	return &p
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.NotStarted
	s.ProgressP2 = utils.NotStarted
	s.DateStarted = "2024-12-14 09:19:41"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("13")
	p := Puzzle{year: iyear, day: iday, title: "Day 13: Claw Contraption"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) Add(value int) {
	for _, g := range puzzle.games {
		g.prize.x += value
		g.prize.y += value
	}
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines

	games := make([]*Game, 0)
	for index := 0; index < len(lines); index++ {
		g := make([]string, 0)
		g = append(g, lines[index])
		g = append(g, lines[index+1])
		g = append(g, lines[index+2])
		game := NewGame(g)
		index += 3
		games = append(games, game)
	}
	puzzle.games = games
}

func (puzzle *Puzzle) Part1() {
	puzzle.Load(TEST_DATA)
	total := 0
	for _, game := range puzzle.games {
		ga := game.Play1()
		if ga != nil {
			fmt.Println(ga.debug())
			fmt.Println("Resolved.\n")
			total += ga.totalCost()

		} else {
			fmt.Println(game.debug())
			fmt.Println("Cannot resolve.\n")
		}
	}
	fmt.Printf("total cost is %v\n", total)
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(TEST_DATA)
	puzzle.Add(10000000000000)
	total := 0
	for _, game := range puzzle.games {
		ga := game.Play2()
		if ga != nil {
			fmt.Println(ga.debug())
			fmt.Println("Resolved.\n")
			total += ga.totalCost()

		} else {
			fmt.Println(game.debug())
			fmt.Println("Cannot resolve.\n")
		}
	}
	fmt.Printf("P2 total cost is %v\n", total)

}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
