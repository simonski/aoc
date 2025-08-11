package aoc2020

import (
	"testing"
)

func Test_AOC2020_06_BoardingPass(t *testing.T) {

	verifyAnswerCounts("abc", 3, t)
	verifyAnswerCounts("b", 1, t)
	verifyAnswerCounts("zza", 2, t)
	verifyAnswerCounts("azaztt", 3, t)
	verifyAnswerCounts("a", 1, t)
	verifyAnswerCounts("aA", 1, t)
	verifyAnswerCounts("a123", 1, t)
	verifyAnswerCounts(":a_", 1, t)
	verifyAnswerCounts(" a ", 1, t)
	verifyAnswerCounts("	a	", 1, t)
	verifyAnswerCounts("", 0, t)
	verifyAnswerCounts("	", 0, t)
	verifyAnswerCounts("  	", 0, t)
	verifyAnswerCounts("  	   ", 0, t)
	verifyAnswerCounts("  	   \\", 0, t)
	verifyAnswerCounts("  	   !", 0, t)

	content := `abcx
	abcy
	abcz`

	verifyAnswerCounts(content, 6, t)

}

func Test_AOC2020_06_BoardingPass2(t *testing.T) {

	verifyAnswerCounts2("abc", 3, t)
	verifyAnswerCounts2("b", 1, t)
	verifyAnswerCounts2("za", 2, t)
	verifyAnswerCounts2("atf", 3, t)
	verifyAnswerCounts2("a", 1, t)
	verifyAnswerCounts2("aA", 1, t)
	verifyAnswerCounts2("a123", 1, t)
	verifyAnswerCounts2(":a_", 1, t)
	verifyAnswerCounts2(" a ", 1, t)
	verifyAnswerCounts2("	a	", 1, t)
	verifyAnswerCounts2("", 0, t)
	verifyAnswerCounts2("	", 0, t)
	verifyAnswerCounts2("  	", 0, t)
	verifyAnswerCounts2("  	   ", 0, t)
	verifyAnswerCounts2("  	   \\", 0, t)
	verifyAnswerCounts2("  	   !", 0, t)

	content := `abcx
	abcy
	abcz`

	verifyAnswerCounts2(content, 3, t)

	content = `abc

	a
	b
	c
	
	ab
	ac
	
	a
	a
	a
	a
	
	b`

	verifyAnswerCounts2(content, 6, t)

}

// verifyQNA helper to check multiple test cases
func verifyAnswerCounts(line string, expected int, t *testing.T) {
	qna := NewQandA(line)
	actual := qna.TotalForAllGroups()
	if actual != expected {
		// t.Errorf("Expected %v, got %v\n\n%v\n\n%v\n\n", expected, actual, line, qna.answers)
		t.Errorf("Expected %v, got %v\n\n%v\n\n", expected, actual, line)
	}
}

// verifyQNA helper to check multiple test cases
func verifyAnswerCounts2(line string, expected int, t *testing.T) {
	qna := NewQandA(line)
	actual := qna.SecondTotalForAllGroups()
	if actual != expected {
		// t.Errorf("Expected %v, got %v\n\n%v\n\n%v\n\n", expected, actual, line, qna.answers)
		t.Errorf("SecondTotalForAllGroups Expected %v, got %v\n\n%v\n\n", expected, actual, line)
	}
}
