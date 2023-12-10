package d10

import (
	"testing"
)

func Test_1(t *testing.T) {
	g := NewGrid(TEST_DATA_1)
	t.Logf("\n%v\n", g.Debug())
}

func Test_2(t *testing.T) {
	g := NewGrid(TEST_DATA_2)
	t.Logf("\n%v\n", g.Debug())
}

func Test_3(t *testing.T) {
	g := NewGrid(TEST_DATA_3)
	t.Logf("\n%v\n", g.Debug())
}
