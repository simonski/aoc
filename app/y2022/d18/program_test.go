package d18

import (
	"fmt"
	"testing"
)

func Test_1x(t *testing.T) {
	p := NewPuzzleWithData(TEST_DATA)
	fmt.Printf("There are %v lines.\n", len(p.lines))
}
