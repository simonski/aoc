package d15

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	p := NewPuzzleWithData(TEST_DATA)
	fmt.Printf("There are %v lines.\n", len(p.lines))
}
