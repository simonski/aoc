package d15

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	g := NewGrid(TEST_DATA)
	if len(g.Beacons) != 14 {
		t.Fatalf("There are %v beacons, there should be 14.\n", len(g.Beacons))
	}

	if len(g.Sensors) != 14 {
		t.Fatalf("There are %v Sensors, there should be 14.\n", len(g.Sensors))
	}

	for _, sensor := range g.Sensors {
		if sensor.Beacon == nil {
			t.Fatal("Each sensor requires a beacon.\n")
		}
	}

	if g.Sensors[0].Point.X != 2 {
		t.Fatalf("Sensor[0] should be 2,18, was %v\n", g.Sensors[0])
	}

	if g.Sensors[0].Point.Y != 18 {
		t.Fatalf("Sensor[0] should be 2,18, was %v\n", g.Sensors[0])
	}

	if g.Beacons[0].Point.X != -2 {
		t.Fatalf("Beacon[0] should be -2,15, was %v\n", g.Beacons[0])
	}

	if g.Beacons[0].Point.Y != 15 {
		t.Fatalf("Beacon[0] should be -2,15, was %v\n", g.Beacons[0])
	}

}

func Test_1_DebugGrid(t *testing.T) {
	g := NewGrid(TEST_DATA)
	fmt.Println(g.Debug())
	t.Fatalf("grid is %v x %v = %v cells\n", g.Width(), g.Height(), g.Width()*g.Height())
}

func Test_1_DebugGrid_Real(t *testing.T) {
	g := NewGrid(REAL_DATA)
	// fmt.Println(g.Debug())
	t.Fatalf("grid is %v x %v = %v cells\n", g.Width(), g.Height(), g.Width()*g.Height())
	// 22,427,144,772,599 // 22tn
}

func Test_Distance(t *testing.T) {
	p1 := NewPoint("0,0")
	p2 := NewPoint("10,10")
	d := p1.Distance(p2)
	fmt.Printf("Distance [ %v->%v ] = %v\n", p1, p2, d)

	p1 = NewPoint("10,4")
	p2 = NewPoint("-3,11")
	d = p1.Distance(p2)
	fmt.Printf("Distance [ %v->%v ] = %v\n", p1, p2, d)

	t.Fatalf("Distance")
}

func Test_1_DebugGrid_2(t *testing.T) {
	g := NewGrid(TEST_DATA)
	detectable, undetectable := g.CountCannotsForRow(10)
	if undetectable != 26 {
		t.Fatalf("Cannot count should be 26, was %v, detectable was %v\n.", undetectable, detectable)
	}
}
func Test_1_DebugGrid_3(t *testing.T) {
	g := NewGrid(REAL_DATA)
	detectable, undetectable := g.CountCannotsForRow(2000000)
	t.Fatalf("detectable=%v, undetectable=%v", detectable, undetectable)
}

func Test_CountMissingOnRow1(t *testing.T) {
	g := NewGrid(TEST_DATA)
	missing := g.CountMissing(true, 11)
	t.Fatalf("missing should be %v, was %v\n", 14, missing)
}

func Test_CountMissingOnRow2(t *testing.T) {
	g := NewGrid(REAL_DATA)
	result := make([]int, 0)
	resultrow := 0
	for row := 0; row < 4000000; row++ {
		if row%10000 == 0 {
			fmt.Println(row)
		}
		missing := g.CountMissing(false, row)
		if len(missing) == 1 {
			result = missing
			resultrow = row
			break
		}
	}
	t.Fatalf("row is %v, col is %v\n", resultrow, result)
}
