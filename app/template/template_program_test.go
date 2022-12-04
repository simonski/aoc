package TOKEN_PACKAGE

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	p := NewPuzzle(TEST_DATA)
	fmt.Printf("There are %v lines.\n", len(p.lines))
}
