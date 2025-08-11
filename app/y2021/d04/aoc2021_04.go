package aoc2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
--- Day 4: Giant Squid ---
You're already almost 1.5km (almost a mile) below the surface of the ocean, already so deep that you can't see any sunlight. What you can see, however, is a giant squid that has attached itself to the outside of your submarine.

Maybe it wants to play bingo?

Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the chosen number is marked on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in any row or any column of a board are marked, that board wins. (Diagonals don't count.)

The submarine has a bingo subsystem to help passengers (currently, you and the giant squid) pass the time. It automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For example:

7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
Finally, 24 is drawn:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
At this point, the third board wins because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: 14 21 17 24 4).

The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.

To guarantee victory against the giant squid, figure out which board will win first. What will your final score be if you choose that board?


*/

func (app *Application) Y2021D04_Summary() *utils.Summary {
	s := utils.NewSummary(2021, 4)
	s.Name = "Giant Squid"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

// rename this to the year and day in question
func (app *Application) Y2021D04P1() {
	playGameD1(DAY_2021_04_TEST_DATA)
	playGameD1(DAY_2021_04_DATA)
}

type Game struct {
	numbersToCall []int
	boards        []*Board
}

func (game *Game) PlayFirst() (*Board, int) {
	for index := 0; index < len(game.numbersToCall); index++ {
		numberToCall := game.numbersToCall[index]
		// fmt.Printf("numbersToCall(%v), numberToCall=%v\n", game.numbersToCall, numberToCall)
		for _, board := range game.boards {
			board.Call(numberToCall)
			if board.IsComplete() {
				return board, numberToCall
			}
		}
	}
	return nil, -1
}

func (game *Game) PlayLast() (*Board, int) {
	lastWinningBoard := game.boards[0]
	lastNumberToCall := -1
	for index := 0; index < len(game.numbersToCall); index++ {
		numberToCall := game.numbersToCall[index]
		for _, board := range game.boards {
			if !board.IsComplete() {
				board.Call(numberToCall)
				if board.IsComplete() {
					lastWinningBoard = board
					lastNumberToCall = numberToCall
				}
			}
		}
	}
	return lastWinningBoard, lastNumberToCall
}

func NewGame(data string) *Game {
	lines := strings.Split(data, "\n")
	boards := make([]*Board, 0)
	numbersToCall := utils.SplitDataToListOfInts(lines[0], ",")
	// avoid the blank line then read the boards
	for index := 2; index < len(lines); index += 6 {
		boardLines := lines[index : index+5]
		board := NewBoard(boardLines)
		boards = append(boards, board)
	}

	game := Game{boards: boards, numbersToCall: numbersToCall}
	return &game
}

type Board struct {
	values [5][5]int
	marked [5][5]bool
}

func (board *Board) Call(number int) bool {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if board.values[row][col] == number {
				board.marked[row][col] = true
				return true
			}
		}
	}
	return false
}

func (board *Board) IsComplete() bool {
	// check rows
	for row := 0; row < 5; row++ {
		result := true
		for col := 0; col < 5; col++ {
			if !board.marked[row][col] {
				result = false
				break
			}
		}
		if result {
			return result
		}
	}

	// check cols
	for col := 0; col < 5; col++ {
		result := true
		for row := 0; row < 5; row++ {
			if !board.marked[row][col] {
				result = false
				break
			}
		}
		if result {
			return result
		}
	}

	return false

}

func (board *Board) SumOfUnmarked() int {
	// check rows
	total := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !board.marked[row][col] {
				total += board.values[row][col]
			}
		}
	}
	return total
}

/*
returns a string
n n n n n
n n n n n
n n n n n
n n n n n
n n n n n
*/
func (b *Board) Debug() string {
	result := ""
	for row := 0; row < 5; row++ { // I dont understand multidimensional array lengths ingo
		for col := 0; col < 5; col++ {
			result += fmt.Sprintf("%v ", b.values[row][col])
		}
		result += "\n"
	}
	return result
}

func NewBoard(data []string) *Board {
	b := Board{values: [5][5]int{{}}, marked: [5][5]bool{{}}}
	for row, line := range data {
		line = strings.Trim(line, " ")
		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.ReplaceAll(line, " ", ",")
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "\n", "")
		splits := strings.Split(line, ",")
		for col, number := range splits {
			value, _ := strconv.Atoi(number)
			b.values[row][col] = value
			b.marked[row][col] = false // unnecessary as defaults but whatever
		}
	}
	// fmt.Println()
	return &b
}

func playGameD1(data string) {
	game := NewGame(data)
	winningBoard, lastNumber := game.PlayFirst()
	score := winningBoard.SumOfUnmarked() * lastNumber
	fmt.Printf("The winning store is %v\n", score)
}

func playGameD2(data string) {
	game := NewGame(data)
	winningBoard, lastNumber := game.PlayLast()
	score := winningBoard.SumOfUnmarked() * lastNumber
	fmt.Printf("The last winning store is %v\n", score)
}

// rename this to the year and day in question
func (app *Application) Y2021D04P2() {
	playGameD2(DAY_2021_04_TEST_DATA)
	playGameD2(DAY_2021_04_DATA)
}

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
func (app *Application) Y2021D04() {
	app.Y2021D04P1()
	app.Y2021D04P2()
}
