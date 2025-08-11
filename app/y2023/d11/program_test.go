package d11

import (
	"testing"
)

func Test_1(t *testing.T) {
	u := NewUniverse(TEST_DATA_1)
	t.Logf("\n%v", u.Debug())

	u.Expand()
	t.Logf("\n%v", u.Debug())

	pairs := u.Pairs()
	if len(pairs) != 36 {
		t.Fatalf("Expected 36 pairs, got %v\n", len(pairs))
	}

	total := 0
	for _, p := range pairs {
		cell1 := p[0]
		cell2 := p[1]
		distance := cell1.Distance(cell2)
		total += distance
	}
	if total != 375 {
		t.Fatalf("Expected 374 distance, got %v\n", total)
	}

}
