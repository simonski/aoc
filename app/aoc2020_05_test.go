package app

import (
	"testing"
)

func Test_AOC2020_05_BoardingPass(t *testing.T) {
	CheckBoardingPass("FBFBBFFRLR", 44, 5, 357, t)
	CheckBoardingPass("BFFFBBFRRR", 70, 7, 567, t)
	CheckBoardingPass("FFFBBBFRRR", 14, 7, 119, t)
	CheckBoardingPass("BBFFBBFRLL", 102, 4, 820, t)
}

func CheckBoardingPass(line string, expectRow int, expectCol int, expectSeat int, t *testing.T) {
	pass := NewBoardingPass(line)
	actualRow := pass.GetRow()
	actualCol := pass.GetCol()
	actualSeat := pass.GetSeatId()

	if actualRow != expectRow || actualCol != expectCol || actualSeat != expectSeat {
		t.Errorf("%v should be row %v col %v seat %v, but was %v, %v, %v\n", line, expectRow, expectCol, expectSeat, actualRow, actualCol, actualSeat)
	}
}
