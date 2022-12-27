package d17

import (
	"fmt"
	"testing"
)

func Test_RockP(t *testing.T) {
	c := NewChamber(TEST_DATA, 0)

	plus := c.ROCK_PLUS
	if plus.height != 3 {
		t.Fatalf("ROCK_PLUS should be 3 height, was %v\n", plus.height)
	}
	if plus.width != 3 {
		t.Fatalf("ROCK_PLUS should be 3 width, was %v\n", plus.width)
	}
	pc := plus.Clone()
	fmt.Println(pc.Debug())
	if !pc.Equals(plus) {
		t.Fatalf("CLoned PLUS does not equal its clone but should.")
	}

	if pc.Equals(c.ROCK_L) {
		t.Fatalf("CLoned PLUS must not equal ROCK_L")
	}

	if pc.Equals(c.ROCK_HORIZONTAL) {
		t.Fatalf("CLoned PLUS must not equal ROCK_L")
	}
}

func Test_RockH(t *testing.T) {
	c := NewChamber(TEST_DATA, 0)

	rock := c.ROCK_HORIZONTAL
	if rock.height != 1 {
		t.Fatalf("ROCK_HORIZONTAL should be 1 height, was %v\n", rock.height)
	}
	if rock.width != 4 {
		t.Fatalf("ROCK_HORIZONTAL should be 4 width, was %v\n", rock.width)
	}
	pc := rock.Clone()
	fmt.Println(pc.Debug())
	if !pc.Equals(rock) {
		t.Fatalf("CLoned H does not equal its clone but should.")
	}

	if pc.Equals(c.ROCK_L) {
		t.Fatalf("CLoned PLUS must not equal ROCK_L")
	}

	if pc.Equals(c.ROCK_SQUARE) {
		t.Fatalf("CLoned H must not equal ROCK_SQUARE")
	}

	if rock.GetPieceChamberXY(0, 0, false) == nil {
		t.Fatalf("Piece(0,0) in SQUARE should be solid.")
	}

	if rock.GetPieceChamberXY(5, 5, false) != nil {
		t.Fatalf("Piece(5,5) in SQUARE should be nil.")
	}
}

func Test_NewRock(t *testing.T) {
	c := NewChamber(TEST_DATA, 0)
	r1 := c.NewRock()
	if !r1.Equals(c.ROCK_HORIZONTAL) {
		fmt.Println(r1.Debug())
		t.Fatalf("First rock should be horizontal.")
	}

	r2 := c.NewRock()
	if !r2.Equals(c.ROCK_PLUS) {
		fmt.Println(r2.Debug())
		t.Fatalf("2nd rock should be PLUS.")
	}

	r3 := c.NewRock()
	if !r3.Equals(c.ROCK_L) {
		fmt.Println(r3.Debug())
		t.Fatalf("3rd rock should be L.")
	}

	r4 := c.NewRock()
	if !r4.Equals(c.ROCK_VERTICAL) {
		fmt.Println(r4.Debug())
		t.Fatalf("4th rock should be V.")
	}

	r5 := c.NewRock()
	if !r5.Equals(c.ROCK_SQUARE) {
		fmt.Println(r5.Debug())
		t.Fatalf("5th rock should be SQUARE.")
	}

	r6 := c.NewRock()
	if !r6.Equals(c.ROCK_HORIZONTAL) {
		fmt.Println(r6.Debug())
		t.Fatalf("6th rock should be SQUARE.")
	}

	r7 := c.NewRock()
	if !r7.Equals(c.ROCK_PLUS) {
		fmt.Println(r7.Debug())
		t.Fatalf("7th rock should be PLUS (tick=%v).", c.rockTick)
	}
}

func verifyPieceAt(c *Chamber, x int, y int, expect string, t *testing.T) {
	if c.GetRockPieceString(x, y) != expect {
		t.Fatalf("Piece at %v,%v should be %v\n", x, y, expect)
	}
}

func Test_Debug(t *testing.T) {
	c := NewChamber(TEST_DATA, 0)
	r1 := c.NewRock()
	c.AddRock(r1)
	fmt.Println(c.Debug())
	verifyPieceAt(c, 0, 0, ".", t)
	verifyPieceAt(c, 2, 3, "#", t)
}

func Test_DebugTicks(t *testing.T) {
	c := NewChamber(TEST_DATA, 0)
	c.RunPart1(-1, false, FLOOR_SIZE)
	if true {
		fmt.Println("")
		fmt.Println(c.Debug())
	}
}

func requireRock(rock1 *Rock, rock2 *Rock, t *testing.T) {
	if !rock1.Equals(rock2) {
		t.Fatal("Rock differs.")
	}
}
func Test_DebugRocks(t *testing.T) {
	c := NewChamber(TEST_DATA, 0)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
	requireRock(c.NewRock(), c.ROCK_HORIZONTAL, t)
	requireRock(c.NewRock(), c.ROCK_PLUS, t)
	requireRock(c.NewRock(), c.ROCK_L, t)
	requireRock(c.NewRock(), c.ROCK_VERTICAL, t)
	requireRock(c.NewRock(), c.ROCK_SQUARE, t)
}

func Test_Rock_equals(t *testing.T) {
	c := NewChamber(REAL_DATA, 0)
	horiz := c.ROCK_HORIZONTAL.Clone()
	vert := c.ROCK_VERTICAL.Clone()
	l := c.ROCK_L.Clone()
	plus := c.ROCK_PLUS.Clone()
	sq := c.ROCK_SQUARE.Clone()
	if !horiz.Equals(c.ROCK_HORIZONTAL) {
		t.Fatal("horiz is not equal.\n")
	}
	if len(horiz.pieces) != 4 {
		t.Fatalf("horiz shoudl have 4 pieces, has %v\n", len(horiz.pieces))
	}

	if !sq.Equals(c.ROCK_SQUARE) {
		t.Fatal("square is not equal.\n")
	}
	if len(sq.pieces) != 4 {
		t.Fatalf("sq shoudl have 4 pieces, has %v\n", len(sq.pieces))
	}

	if !vert.Equals(c.ROCK_VERTICAL) {
		t.Fatal("vertical is not equal.\n")
	}
	if len(vert.pieces) != 4 {
		t.Fatalf("vert shoudl have 4 pieces, has %v\n", len(vert.pieces))
	}

	if !plus.Equals(c.ROCK_PLUS) {
		t.Fatal("plus is not equal.\n")
	}
	if len(plus.pieces) != 9 {
		t.Fatalf("plus shoudl have 9 pieces, has %v\n", len(plus.pieces))
	}

	if !l.Equals(c.ROCK_L) {
		t.Fatal("l is not equal.\n")
	}
	if len(l.pieces) != 9 {
		t.Fatalf("vert should have 9 pieces, has %v\n", len(l.pieces))
	}
}

func requirePieceXY(rock *Rock, x int, y int, solid bool, t *testing.T) {
	p := rock.GetPieceAbsoluteXY(x, y)
	if p == nil {
		t.Fatalf("rock should contain a piece at %v,%v\n", x, y)
	}
	if p.x != x || p.y != y {
		t.Fatalf("requirePiece, x=%v, should be %v, y=%v, should be %v", p.x, x, p.y, y)
	}
}
func requireNilPieceXY(rock *Rock, x int, y int, t *testing.T) {
	p := rock.GetPieceAbsoluteXY(x, y)
	if p != nil {
		t.Fatalf("rock should not contain a piece at %v,%v\n", x, y)
	}
}

func Test_Piece_Horizontal(t *testing.T) {
	c := NewChamber(REAL_DATA, 0)
	h := c.ROCK_HORIZONTAL.Clone()
	requirePieceXY(h, 0, 0, true, t)
	requirePieceXY(h, 1, 0, true, t)
	requirePieceXY(h, 2, 0, true, t)
	requirePieceXY(h, 3, 0, true, t)
	requireNilPieceXY(h, 4, 0, t)
}

func Test_Piece_Vertical(t *testing.T) {
	c := NewChamber(REAL_DATA, 0)
	h := c.ROCK_VERTICAL.Clone()
	requirePieceXY(h, 0, 0, true, t)
	requirePieceXY(h, 0, 1, true, t)
	requirePieceXY(h, 0, 2, true, t)
	requirePieceXY(h, 0, 3, true, t)
	requireNilPieceXY(h, 1, 3, t)
}

func Test_Piece_Square(t *testing.T) {
	c := NewChamber(REAL_DATA, 0)
	h := c.ROCK_SQUARE.Clone()
	requirePieceXY(h, 0, 0, true, t)
	requirePieceXY(h, 1, 0, true, t)
	requirePieceXY(h, 0, 1, true, t)
	requirePieceXY(h, 1, 1, true, t)
	requireNilPieceXY(h, 2, 2, t)
}

func Test_Piece_L(t *testing.T) {
	c := NewChamber(REAL_DATA, 0)
	h := c.ROCK_L.Clone()
	requirePieceXY(h, 2, 0, true, t)
	requirePieceXY(h, 2, 1, true, t)
	requirePieceXY(h, 2, 2, true, t)
	requirePieceXY(h, 1, 2, true, t)
	requirePieceXY(h, 0, 2, true, t)

	requireNilPieceXY(h, 0, 0, t)
	requireNilPieceXY(h, 1, 0, t)
	requireNilPieceXY(h, 1, 1, t)
	requireNilPieceXY(h, 1, 1, t)
}

func Test_Piece_Plus(t *testing.T) {
	c := NewChamber(REAL_DATA, 0)
	h := c.ROCK_PLUS.Clone()
	requireNilPieceXY(h, 0, 0, t)
	requirePieceXY(h, 1, 0, true, t)
	requireNilPieceXY(h, 2, 0, t)

	requirePieceXY(h, 0, 1, true, t)
	requirePieceXY(h, 1, 1, true, t)
	requirePieceXY(h, 2, 1, true, t)

	requireNilPieceXY(h, 0, 2, t)
	requirePieceXY(h, 1, 2, true, t)
	requireNilPieceXY(h, 2, 2, t)
}

// func Test_Edges_Horizontal(t *testing.T) {
// 	c := NewChamber(REAL_DATA)
// 	left := c.ROCK_HORIZONTAL.GetLeftmostPieces(c)
// 	bottom := c.ROCK_HORIZONTAL.GetBottomPieces(c)
// 	right := c.ROCK_HORIZONTAL.GetRightmostPieces(c)

// 	if len(left) != 1 {
// 		t.Fatalf("HORIZONTAL should have 1 left most but has %v\n", len(left))
// 	}
// 	if len(right) != 1 {
// 		t.Fatalf("HORIZONTAL should have 1 right most but has %v\n%v\n", len(right), right)
// 	}
// 	if len(bottom) != 4 {
// 		t.Fatalf("HORIZONTAL should have 4 bottom most but has %v\n", len(bottom))
// 	}
// }

// func Test_Edges_Vertical(t *testing.T) {
// 	c := NewChamber(REAL_DATA)
// 	left := c.ROCK_VERTICAL.GetLeftmostPieces(c)
// 	right := c.ROCK_VERTICAL.GetRightmostPieces(c)
// 	bottom := c.ROCK_VERTICAL.GetBottomPieces(c)
// 	if len(left) != 4 {
// 		t.Fatalf("ROCK_VERTICAL should have 4 left most but has %v\n", len(left))
// 	}
// 	if len(bottom) != 1 {
// 		t.Fatalf("ROCK_VERTICAL should have 1 bottom most but has %v\n", len(bottom))
// 	}
// 	if len(right) != 4 {
// 		t.Fatalf("ROCK_VERTICAL should have 4 right most but has %v\n", len(right))
// 	}
// }

// func Test_Edges_Square(t *testing.T) {
// 	c := NewChamber(REAL_DATA)
// 	left := c.ROCK_SQUARE.GetLeftmostPieces(c)
// 	right := c.ROCK_SQUARE.GetRightmostPieces(c)
// 	bottom := c.ROCK_SQUARE.GetBottomPieces(c)
// 	if len(left) != 2 {
// 		t.Fatalf("ROCK_SQUARE should have 2 left most but has %v\n", len(left))
// 	}
// 	if len(bottom) != 2 {
// 		t.Fatalf("ROCK_SQUARE should have 2 bottom most but has %v\n", len(bottom))
// 	}
// 	if len(right) != 2 {
// 		t.Fatalf("ROCK_SQUARE should have 2 right most but has %v\n", len(right))
// 	}
// }

// func Test_Edges_L(t *testing.T) {
// 	c := NewChamber(REAL_DATA)
// 	left := c.ROCK_L.GetLeftmostPieces(c)
// 	bottom := c.ROCK_L.GetRightmostPieces(c)
// 	right := c.ROCK_L.GetBottomPieces(c)
// 	if len(left) != 3 {
// 		t.Fatalf("ROCK_L should have 3 left most but has %v\n", len(left))
// 	}
// 	if len(bottom) != 3 {
// 		t.Fatalf("ROCK_L should have 3 bottom most but has %v\n", len(bottom))
// 	}
// 	if len(right) != 3 {
// 		t.Fatalf("ROCK_L should have 3 right most but has %v\n", len(right))
// 	}
// }

// func Test_Edges_Plus(t *testing.T) {
// 	c := NewChamber(REAL_DATA)
// 	left := c.ROCK_PLUS.GetLeftmostPieces(c)
// 	bottom := c.ROCK_PLUS.GetRightmostPieces(c)
// 	right := c.ROCK_PLUS.GetBottomPieces(c)
// 	if len(left) != 3 {
// 		t.Fatalf("ROCK_PLUS should have 3 left most but has %v\n", len(left))
// 	}
// 	if len(bottom) != 3 {
// 		t.Fatalf("ROCK_PLUS should have 3 bottom most but has %v\n", len(bottom))
// 	}
// 	if len(right) != 3 {
// 		t.Fatalf("ROCK_PLUS should have 3 right most but has %v\n", len(right))
// 	}
// }

func Test_Part1_Test(t *testing.T) {
	c := NewChamber(TEST_DATA, 0)
	fmt.Println(c.Debug())
	c.RunPart1(2022, false, FLOOR_SIZE)
	fmt.Println(c.Debug())
	fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
}

func Test_Part1_Test_FindFloor(t *testing.T) {
	c := NewChamber(TEST_DATA, 0)
	fmt.Println(c.Debug())
	// c.Run(true, true, 41)
	c.RunPart1(2022, true, FLOOR_SIZE)
	fmt.Println(c.Debug())
	fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
}

func Test_Part1_Real(t *testing.T) {
	c := NewChamber(REAL_DATA, 0)
	fmt.Println(c.Debug())
	floorSize := 4000
	c.RunPart1(2022, false, floorSize)
	fmt.Println(c.Debug())
	fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
}

func Test_Part1_Real_FindFloor(t *testing.T) {
	c := NewChamber(REAL_DATA, 0)
	fmt.Println(c.Debug())
	c.RunPart1(2022, true, FLOOR_SIZE)
	fmt.Println(c.Debug())
	fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)

}

func Test_Part2_Test(t *testing.T) {
	c := NewChamber(TEST_DATA, 0)
	fmt.Println(c.Debug())
	c.RunPart2(1000000000000, true, FLOOR_SIZE)
	fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
}

func Test_Part2_Real(t *testing.T) {
	c := NewChamber(REAL_DATA, 0)
	fmt.Println(c.Debug())
	c.RunPart2(1000000000000, true, FLOOR_SIZE)
	fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
}

func Test_Part1(t *testing.T) {
	// this test is so I can verify as I work on part2 I haven't broken anything
	c := NewChamber(TEST_DATA, 0)
	c.RunPart1(2022, false, FLOOR_SIZE)
	if c.Height != 3068 {
		fmt.Println(c.Debug())
		t.Fatalf("Part1 TEST_DATA should be 3068, was %v\n", c.Height)
	}

	c2 := NewChamber(REAL_DATA, 0)
	c2.RunPart1(2022, false, FLOOR_SIZE)
	if c.Height != 3119 {
		fmt.Println(c2.Debug())
		t.Fatalf("Part1 REAL_DATA should be 3119, was %v\n", c2.Height)
	}

}
