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

func Test_4(t *testing.T) {
	g := NewGrid(TEST_DATA_P2_1)
	t.Log("1")
	t.Logf("\n\n%v\n\n", g.Debug())
	t.Log("")

	g = NewGrid(TEST_DATA_P2_2)
	t.Log("2")
	t.Logf("\n\n%v\n\n", g.Debug())
	t.Log("")

	g = NewGrid(TEST_DATA_P2_3)
	t.Log("3")
	t.Logf("\n\n%v\n\n", g.Debug())
	t.Log("")

	g = NewGrid(TEST_DATA_P2_4)
	t.Log("4")
	t.Logf("\n\n%v\n\n", g.Debug())
	t.Log("")

	g = NewGrid(TEST_DATA_P2_5)
	t.Log("5")
	t.Logf("\n\n%v\n\n", g.Debug())
	t.Log("")

	g = NewGrid(TEST_DATA_P2_6)
	t.Log("6")
	t.Logf("\n\n%v\n\n", g.Debug())
	t.Log("")

	g = NewGrid(TEST_DATA_P2_7)
	t.Log("7")
	t.Logf("\n\n%v\n\n", g.Debug())
	t.Log("")

	t.Fatal("x")
}
