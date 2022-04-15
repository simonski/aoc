package aoc2021

import (
	"fmt"
	"strings"
	"testing"
)

func Test_AOC2021_19_Split1(t *testing.T) {
	p := Pair{}
	value := "10"
	expected := "[5,5]"
	actual := p.Split(value)
	if actual != expected {
		t.Errorf("%v should split to %v, but was %v\n", value, expected, actual)
	}
}

func Test_AOC2021_19_Split2(t *testing.T) {
	p := Pair{}
	value := "11"
	expected := "[5,6]"
	actual := p.Split(value)
	if actual != expected {
		t.Errorf("%v should split to %v, but was %v\n", value, expected, actual)
	}
}

func Test_AOC2021_19_Split3(t *testing.T) {
	p := Pair{}
	value := "12"
	expected := "[6,6]"
	actual := p.Split(value)
	if actual != expected {
		t.Errorf("%v should split to %v, but was %v\n", value, expected, actual)
	}
}

func Test_AOC2021_19_SplitAt(t *testing.T) {
	value := "[12,4]"
	p := NewPair(value)
	expected := "[[6,6],4]"
	p.Reduce()
	actual := p.Line
	if actual != expected {
		t.Errorf("%v should split to %v, but was %v\n", value, expected, actual)
	}
}

func Test_AOC2021_19_FindPair_1(t *testing.T) {
	// [[[[[9,8],1],2],3],4] becomes [[[[0,9],2],3],4] (the 9 has no regular number to its left, so it is not added to any regular number).
	line := "[[[[[9,8],1],2],3],4]"
	p := NewPair(line)
	// expected := "[[[[0,9],2],3],4]"
	index := p.FindPairAtIndex(4, line)
	expected := "[9,8]"
	if index != expected {
		t.Errorf("%v should be %v, but was %v\n", p.Line, expected, index)
	}
}

func Test_AOC2021_19_FindPair_2(t *testing.T) {
	// [[[[[9,8],1],2],3],4] becomes [[[[0,9],2],3],4] (the 9 has no regular number to its left, so it is not added to any regular number).
	line := "[[[[[9,8],1],2],3],4]"
	p := NewPair(line)
	// expected := "[[[[0,9],2],3],4]"
	index := p.FindPairAtIndex(3, line)
	expected := "[[9,8],1]"
	if index != expected {
		t.Errorf("%v should be %v, but was %v\n", p.Line, expected, index)
	}
}

func Test_AOC2021_19_FindPairToExplode_1(t *testing.T) {
	line := "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"
	p := NewPair(line)
	// expected := "[[[[0,9],2],3],4]"
	index := p.FindIndexOfFifthPair(line)
	expected := 10
	if index != expected {
		t.Errorf("%v should be %v, but was %v\n", p.Line, expected, index)
	}

	line = "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"

	p = NewPair(line)
	index = p.FindIndexOfFifthPair(line)
	expected = 24
	if index != expected {
		t.Errorf("%v should be %v, but was %v\n", p.Line, expected, index)
	}

}

func Test_AOC2021_19_Explode_1(t *testing.T) {
	// [[[[[9,8],1],2],3],4] becomes [[[[0,9],2],3],4] (the 9 has no regular number to its left, so it is not added to any regular number).
	line := "[[[[[9,8],1],2],3],4]"
	p := NewPair(line)
	index := p.FindIndexOfFifthPair(line)

	expected := 4
	if index != expected {
		t.Errorf("%v should be %v, but was %v\n", p.Line, expected, index)
	}

	actual, _ := p.ExplodeLeftMostPair(index, line)
	expected_exploded := "[[[[0,9],2],3],4]"
	// actual := p.Line
	if actual != expected_exploded {
		t.Errorf("%v should be %v, but was %v\n", p.Line, expected, actual)
	}
}

func Test_AOC2021_19_Reduce_1(t *testing.T) {
	// [[[[[9,8],1],2],3],4] becomes [[[[0,9],2],3],4] (the 9 has no regular number to its left, so it is not added to any regular number).
	p := NewPair("[[[[[9,8],1],2],3],4]")
	expected := "[[[[0,9],2],3],4]"
	p.Reduce()
	actual := p.Line
	if actual != expected {
		t.Errorf("%v should be %v, but was %v\n", p.Line, expected, actual)
	}
}

func Test_AOC2021_19_Reduce_2(t *testing.T) {
	p := NewPair("[7,[6,[5,[4,[3,2]]]]]")
	expected := "[7,[6,[5,[7,0]]]]"
	p.Reduce()
	actual := p.Line
	if actual != expected {
		t.Errorf("%v should be %v, but was %v\n", p.Line, expected, actual)
	}
}

func Test_AOC2021_19_Reduce_3(t *testing.T) {
	p := NewPair("[[6,[5,[4,[3,2]]]],1]")
	expected := "[[6,[5,[7,0]]],3]"
	p.Reduce()
	actual := p.Line
	if actual != expected {
		t.Errorf("%v should be %v, but was %v\n", p.Line, expected, actual)
	}
}

func Test_AOC2021_19_Reduce_4(t *testing.T) {
	p := NewPair("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
	expected1 := "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"
	p.DoReduce("none")
	actual1 := p.Line
	if actual1 != expected1 {
		t.Errorf("Test_AOC2021_18_Reduce_4 Step 1 %v should be %v, but was %v\n", p.Line, expected1, actual1)
	}

	expected2 := "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"
	p.DoReduce("none")
	actual2 := p.Line
	if actual2 != expected2 {
		t.Errorf("Test_AOC2021_18_Reduce_4 Step 2 %v should be %v, but was %v\n", p.Line, expected1, actual2)
	}
}

func Test_AOC2021_19_Reduce_5(t *testing.T) {
	p := NewPair("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
	expected := "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"
	p.Reduce()
	actual := p.Line
	if actual != expected {
		t.Errorf("Test_AOC2021_18_Reduce_5 %v should be %v, but was %v\n", p.Line, expected, actual)
	}
}

func Test_AOC2021_19_FindIndexOfRegularNumberGreaterOrEqualToTen(t *testing.T) {

	p := NewPair("")
	data := "[[[[5,0],[7,4]],[15,5]],[6,6]]"
	index := p.FindIndexOfRegularNumberGreaterOrEqualToTen(data)
	if index != 17 {
		t.Errorf("FindIndexOfRegularNumberGreaterOrEqualToTen expected 17, got %v.\n", index)
	}

	data = "11[[[[5,0],[7,4]],[15,5]],[6,6]]"
	index = p.FindIndexOfRegularNumberGreaterOrEqualToTen(data)
	if index != 0 {
		t.Errorf("FindIndexOfRegularNumberGreaterOrEqualToTen expected 0, got %v.\n", index)
	}

	data = "[,[,[,[11[[[[5,0],[7,4]],[15,5]],[6,6]]"
	index = p.FindIndexOfRegularNumberGreaterOrEqualToTen(data)
	if index != 7 {
		t.Errorf("FindIndexOfRegularNumberGreaterOrEqualToTen expected 7, got %v.\n", index)
	}

	data = "[,[,[,[1[[[[5,0],[7,4]],[5,5]],[6,6]]"
	index = p.FindIndexOfRegularNumberGreaterOrEqualToTen(data)
	if index != -1 {
		t.Errorf("FindIndexOfRegularNumberGreaterOrEqualToTen expected -1, got %v.\n", index)
	}

}

func Test_AOC2021_19_Add_FindIndexOfFifthPair(t *testing.T) {

	data := "[[[[5,0],[7,4]],[5,5]],[6,6]]"
	p := NewPair("")
	index := p.FindIndexOfFifthPair(data)
	if index != -1 {
		t.Errorf("FindIndexOfFifthPair expected 10, got %v.\n", index)
	}

}

func Test_AOC2021_19_Add_1(t *testing.T) {

	p1 := NewPair("[1,2]")
	p2 := "[[3,4],5]"
	//  becomes [[1,2],[[3,4],5]].

	p1.Add(p2)
	expected := "[[1,2],[[3,4],5]]"
	actual := p1.Line
	if actual != expected {
		t.Errorf("Test_AOC2021_18_Add_1 %v +  %v should be %v, but was %v\n", p1.Line, p2, actual, expected)
	}

}

// FAILING FROM HERE

func Test_AOC2021_19_Add_0(t *testing.T) {

	style := "none"
	line := "[[[[4,3],4],4],[7,[[8,4],9]]]"
	to_add := "[1,1]"
	p := NewPair(line)
	p.Add(to_add)
	expected_1 := "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"
	if p.Line != expected_1 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 1: got %v (add), expected %v\n", p.Line, expected_1)
	}

	r, _, style := p.DoReduce(style)
	if r != 1 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 2: (explode) expected 2 reductions, got %v\n", r)
	}
	expected_2 := "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]"
	if p.Line != expected_2 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 2: got %v expected %v\n", p.Line, expected_2)
	}

	fmt.Println()
	fmt.Println("Step3")
	r, _, style = p.DoReduce(style)
	if r != 1 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 3: expected 1 reductions, got %v\n", r)
	}
	expected_3 := "[[[[0,7],4],[15,[0,13]]],[1,1]]"
	if p.Line != expected_3 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 3: (explode) got %v expected %v\n", p.Line, expected_3)
	}

	fmt.Println()
	fmt.Println("Step4")
	r, _, style = p.DoReduce(style)
	if r != 1 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 4: expected 1 reductions, got %v\n", r)
	}
	expected_4 := "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"
	if p.Line != expected_4 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 4: (split) got %v expected %v\n", p.Line, expected_4)
	}

	fmt.Println()
	fmt.Println("Step5")
	r, _, style = p.DoReduce(style)
	if r != 1 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 5: expected 1 reductions, got %v\n", r)
	}
	expected_5 := "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"
	if p.Line != expected_5 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 5: (split) got %v expected %v\n", p.Line, expected_5)
	}
	fmt.Println()

	r, _, style = p.DoReduce(style)
	if r != 1 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 6: expected 1 reductions, got %v\n", r)
	}
	expected_6 := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	if p.Line != expected_6 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 6: (explode) got %v expected %v\n", p.Line, expected_6)
	}

	// r = p.DoReduce()
	// if r != 1 {
	// 	t.Fatalf("Test_AOC2021_18_Add_0 Step 7: expected 1 reductions, got %v\n", r)
	// }
	// expected_7 := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	// if p.Line != expected_7 {
	// 	t.Fatalf("Test_AOC2021_18_Add_0 Step 7: got %v expected %v\n", p.Line, expected_7)
	// }

	// r = p.DoReduce()
	// if r != 1 {
	// 	t.Fatalf("Test_AOC2021_18_Add_0 Step 8: expected 1 reductions, got %v\n", r)
	// }
	// expected_8 := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	// if p.Line != expected_8 {
	// 	t.Fatalf("Test_AOC2021_18_Add_0 Step 8: got %v expected %v\n", p.Line, expected_8)
	// }

	r, _, _ = p.DoReduce("none")
	if r != 0 {
		t.Fatalf("Test_AOC2021_18_Add_0 Step 9: expected 0 reductions, got %v\n", r)
	}

}

func Test_AOC2021_19_Add_2(t *testing.T) {

	expected := "[[[[1,1],[2,2]],[3,3]],[4,4]]"
	data := `[1,1]
	[2,2]
	[3,3]
	[4,4]`

	splits := strings.Split(data, "\n")

	p := NewPair(splits[0])
	for index := 1; index < len(splits); index++ {
		line := splits[index]
		p.Add(line)
	}
	actual := p.Line
	if actual != expected {
		t.Errorf("expected %v, got %v\n", expected, actual)
	}

}

func Test_AOC2021_19_Add_3(t *testing.T) {

	expected := "[[[[3,0],[5,3]],[4,4]],[5,5]]"
	data := `[1,1]
	[2,2]
	[3,3]
	[4,4]
	[5,5]`

	splits := strings.Split(data, "\n")

	p := NewPair(splits[0])
	for index := 1; index < len(splits); index++ {
		line := splits[index]
		p.Add(line)
		p.Reduce()
	}
	actual := p.Line
	if actual != expected {
		t.Errorf("expected %v, got %v\n", expected, actual)
	}

}

func Test_AOC2021_19_Add_4(t *testing.T) {

	expected := "[[[[5,0],[7,4]],[5,5]],[6,6]]"
	data := `[1,1]
	[2,2]
	[3,3]
	[4,4]
	[5,5]
	[6,6]`

	splits := strings.Split(data, "\n")

	p := NewPair(splits[0])
	for index := 1; index < len(splits); index++ {
		line := splits[index]
		p.Add(line)
		p.Reduce()
	}
	actual := p.Line
	if actual != expected {
		t.Errorf("expected %v, got %v\n", expected, actual)
	}

}

func Test_AOC2021_19_Add_5(t *testing.T) {

	expected := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"

	data := `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
	[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
	[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
	[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
	[7,[5,[[3,8],[1,4]]]]
	[[2,[2,2]],[8,[8,1]]]
	[2,9]
	[1,[[[9,3],9],[[9,0],[0,7]]]]
	[[[5,[7,4]],7],1]
	[[[[4,2],2],6],[8,7]]`

	splits := strings.Split(data, "\n")

	expected_at_step := [9]string{
		"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		"[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
		"[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
		"[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
		"[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
		"[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
		"[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"}

	line1 := splits[0]
	line1 = strings.ReplaceAll(line1, " ", "")
	line1 = strings.ReplaceAll(line1, "\t", "")
	p := NewPair(line1)
	for index := 1; index < len(splits); index++ {
		line := splits[index]
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "\t", "")

		original := p.Line
		p.Add(line)
		t.Logf("\n\nStep[%v] (before reduce),\n\n   \"%v\" \n+  \"%v\"\n=  \"%v\"\n\n", index-1, original, line, p.Line)
		p.Reduce()
		expected := expected_at_step[index-1]
		actual := p.Line
		if actual != expected {
			t.Fatalf("\nStep[%v]\nexpected [%v],\ngot      %v\n", index-1, expected, actual)
		} else {
			t.Logf("Step[%v] passed\n", index-1)
		}
	}
	actual := p.Line
	if actual != expected {
		t.Fatalf("expected %v, got %v\n", expected, actual)
	}

}
