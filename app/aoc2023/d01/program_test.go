package d01

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/simonski/aoc/utils"
)

func Test_1(t *testing.T) {
	p := NewPuzzleWithData(REAL_DATA)
	fmt.Printf("There are %v lines.\n", len(p.lines))

	grandTotal := 0
	for _, line := range p.lines {
		lvar := -1
		rvar := -1
		for index := 0; index < len(line)-1; index++ {
			candidate := line[index : index+1]
			i, isInt := utils.IsInt(candidate)
			if isInt {
				lvar = i
				break
			}
		}

		for index := len(line) - 1; index >= 0; index-- {
			candidate := line[index : index+1]
			i, isInt := utils.IsInt(candidate)
			if isInt {
				rvar = i
				break
			}
		}

		total := fmt.Sprintf("%v%v", lvar, rvar)
		sum, _ := strconv.Atoi(total)
		grandTotal += sum
		fmt.Printf("Line: %v, lvar: %v, rvar: %v, sum: %v\n", line, lvar, rvar, sum)
	}
	fmt.Printf("GrantTotal: %v\n", grandTotal)

}
