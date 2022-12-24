package d17

import (
	"fmt"
	"testing"
)

func Test_RockP(t *testing.T) {
	c := NewChamber(TEST_DATA)

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
	c := NewChamber(TEST_DATA)

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

	if !rock.GetPiece(0, 0).solid {
		t.Fatalf("Piece(0,0) in SQUARE should be solid.")
	}

	if rock.GetPiece(5, 5) != nil {
		t.Fatalf("Piece(5,5) in SQUARE should be nil.")
	}

}

func Test_NewRock(t *testing.T) {
	c := NewChamber(TEST_DATA)
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

func Test_Debug(t *testing.T) {
	c := NewChamber(TEST_DATA)
	r1 := c.NewRock()
	c.AddRock(r1)
	if true {
		fmt.Println(c.Debug())
		t.Fatalf("")
	}

}

func Test_DebugTicks(t *testing.T) {
	c := NewChamber(TEST_DATA)
	c.Run(true, -1)
	if true {
		fmt.Println("")
		fmt.Println(c.Debug())
	}

}

func Test_DebugTicks5(t *testing.T) {
	c := NewChamber(TEST_DATA)
	fmt.Println(c.Debug())
	c.Run(false, 2022)
	fmt.Println(c.Debug())
	fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height())

}
