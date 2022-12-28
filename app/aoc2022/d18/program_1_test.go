package d18

import (
	"testing"
)

func Test_1(t *testing.T) {
	g := NewGrid(TEST_DATA)
	if g.Size() != 13 {
		t.Fatalf("Grid size should be 13, was %v\n", g.Size())
	}
}

func Test_Bounds_Test(t *testing.T) {
	g := NewGrid(TEST_DATA)
	x, y, z, _, _, _ := g.Bounds()
	if x != 3 || y != 3 || z != 6 {
		t.Fatalf("Grid bounds should be (3,3,6), was (%v,%v,%v)\n", x, y, z)
	}
}

func Test_Connected_Test(t *testing.T) {
	g := NewGrid(TEST_DATA)
	_, not_connected := g.CountSides()
	if not_connected != 64 {
		t.Fatalf("Grid connected sides should be 64, was %v\n", not_connected)
	}
}

func Test_2_real(t *testing.T) {
	g := NewGrid(REAL_DATA)
	if g.Size() != 2825 {
		t.Fatalf("Grid size should be 13, was %v\n", g.Size())
	}
}

func Test_Bounds_Real(t *testing.T) {
	g := NewGrid(REAL_DATA)
	x, y, z, _, _, _ := g.Bounds()
	if x != 21 || y != 20 || z != 21 {
		t.Fatalf("Grid bounds should be (21,20,21), was (%v,%v,%v)\n", x, y, z)
	}
}

func Test_Connected__real(t *testing.T) {
	g := NewGrid(REAL_DATA)
	_, not_connected := g.CountSides()
	if not_connected != 64 {
		t.Fatalf("Grid connected sides should be 64, was %v\n", not_connected)
	}
}
