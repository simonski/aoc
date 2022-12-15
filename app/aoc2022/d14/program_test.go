package d14

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	p := NewPuzzleWithData(false, "123,02 -> 124,04")
	if len(p.input) != 1 {
		t.Fatalf("Expected 1 line, got %v\n", len(p.input))
	}
	if p.Rows != 124 {
		t.Fatalf("Expected 124 rows, got %v\n", p.Rows)
	}
	if p.Cols != 4 {
		t.Fatalf("Expected 4 rows, got %v\n", p.Cols)
	}

	p = NewPuzzleWithData(false, "123,02 -> 124,04 -> 404,302")
	if len(p.Lines) != 2 {
		t.Fatalf("Expected 2 lines, got %v\n", len(p.Lines))
	}
	if p.Rows != 124 {
		t.Fatalf("Expected 124 rows, got %v\n", p.Rows)
	}
	if p.Cols != 302 {
		t.Fatalf("Expected 302 cols, got %v\n", p.Cols)
	}

	p = NewPuzzleWithData(false, "123,02 -> 124,04 -> 404,302 -> 1,43")
	if len(p.Lines) != 3 {
		t.Fatalf("Expected 3 lines, got %v\n", len(p.Lines))
	}
	if p.Rows != 43 {
		t.Fatalf("Expected 43 rows, got %v\n", p.Rows)
	}
	if p.Cols != 404 {
		t.Fatalf("Expected 404 cols, got %v\n", p.Cols)
	}

	p = NewPuzzleWithData(false, TEST_DATA)
	if len(p.Lines) != 5 {
		t.Fatalf("Expected 5 lines, got %v\n", len(p.Lines))
	}

}

func Test_Debug(t *testing.T) {
	p := NewPuzzleWithData(false, TEST_DATA)
	x, _, max_x, max_y := p.Bounds()
	// fmt.Printf("bounds(%v,%v)->(%v,%v)\n", x, y, max_x, max_y)
	fmt.Println(p.DebugFrame(x, 0, max_x+1, max_y+1))

	t.Fatalf("\nf -rows=%v, cols=%v\n", p.Rows, p.Cols)
}

func Test_DebugSteps(t *testing.T) {
	p := NewPuzzleWithData(false, TEST_DATA)
	x, _, max_x, max_y := p.Bounds()
	// fmt.Printf("bounds(%v,%v)->(%v,%v)\n", x, y, max_x, max_y)
	// fmt.Println(p.DebugFrame(x, 0, max_x+1, max_y+1))

	for index := 0; index < 24; index++ {
		// landed, block := p.Step()
		p.AddSand(max_y)
		fmt.Printf("[%v]\n", index+1)
		fmt.Println(p.DebugFrame(x, 0, max_x+1, max_y+1))
		fmt.Println("")
	}
	t.Fatal("mm")
}

func Test_DebugSteps_Real(t *testing.T) {
	p := NewPuzzleWithData(false, REAL_DATA)
	x, _, max_x, max_y := p.Bounds()
	// fmt.Printf("bounds(%v,%v)->(%v,%v)\n", x, y, max_x, max_y)
	// fmt.Println(p.DebugFrame(x, 0, max_x+1, max_y+1))

	for index := 0; index < 905; index++ {
		// landed, block := p.Step()
		result, _, _ := p.AddSand(max_y)
		// p.AddSand(max_y)
		fmt.Printf("[%v]\n", index+1)
		fmt.Println(p.DebugFrame(x, 0, max_x+1, max_y+1))
		fmt.Println("")
		if !result {
			fmt.Printf("Grains: %v\n", index)
			break
		}
	}
	t.Fatal("mm")
}

func Test_DebugSteps_Test(t *testing.T) {
	p := NewPuzzleWithData(false, TEST_DATA)
	x, _, max_x, max_y := p.Bounds()
	// fmt.Printf("bounds(%v,%v)->(%v,%v)\n", x, y, max_x, max_y)
	// fmt.Println(p.DebugFrame(x, 0, max_x+1, max_y+1))

	for index := 0; index < 905; index++ {
		// landed, block := p.Step()
		result, _, _ := p.AddSand(max_y)
		// p.AddSand(max_y)
		fmt.Printf("[%v]\n", index+1)
		fmt.Println(p.DebugFrame(x, 0, max_x+1, max_y+1))
		fmt.Println("")
		if !result {
			fmt.Printf("Grains: %v\n", index)
			break
		}
	}
	t.Fatal("mm")
}
