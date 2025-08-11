package d09

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	p := NewPuzzleWithData(TEST_DATA_1)
	fmt.Printf("There are %v lines.\n", len(p.lines))
}
