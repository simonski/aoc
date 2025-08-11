package aoc2021

import (
	"testing"
)

func Test_AOC2021_18_Magnitude_1(t *testing.T) {
	checkMagnitude("[[1,2],[[3,4],5]]", 143, t)
	checkMagnitude("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384, t)
	checkMagnitude("[[[[1,1],[2,2]],[3,3]],[4,4]]", 445, t)
	checkMagnitude("[[[[3,0],[5,3]],[4,4]],[5,5]]", 791, t)
	checkMagnitude("[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137, t)
	checkMagnitude("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488, t)
}

func Test_AOC2021_18_Magnitude_2(t *testing.T) {
	p := NewPair("[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]")
	p.Add("[[[5,[2,8]],4],[5,[[9,9],0]]]")
	p.Reduce()
	p.Add("[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]")
	p.Reduce()
	p.Add("[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]")
	p.Reduce()
	p.Add("[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]")
	p.Reduce()
	p.Add("[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]")
	p.Reduce()
	p.Add("[[[[5,4],[7,7]],8],[[8,3],8]]")
	p.Reduce()
	p.Add("[[9,3],[[9,9],[6,[4,9]]]]")
	p.Reduce()
	p.Add("[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]")
	p.Reduce()
	p.Add("[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]")
	p.Reduce()

	expected_line := `[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]`
	actual_line := p.Line
	if actual_line != expected_line {
		t.Fatalf("Should be magnitude %v, but was %v\n", expected_line, actual_line)
	} else {
		t.Logf("the number reduces properly.\n")
	}

	expected := 4140
	actual := p.Magnitude()
	if actual != expected {
		t.Fatalf("Should be magnitude %v, but was %v\n", expected, actual)
	}
}

func checkMagnitude(line string, expected int, t *testing.T) {
	p := Pair{Line: line}
	p.Reduce()
	actual := p.Magnitude()
	if actual != expected {
		t.Fatalf("%v should be magnitude %v, but was %v\n", line, expected, actual)
	}
}
