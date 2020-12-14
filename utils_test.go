package main

import (
	"testing"
)

func Test_Utils_Points_R90(t *testing.T) {

	origin := &Point{0, 0}
	p := &Point{1, 1}
	expected_1 := &Point{1, -1}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_1, p, t)

	expected_2 := &Point{-1, -1}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_2, p, t)

	expected_3 := &Point{-1, 1}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_3, p, t)

	expected_4 := &Point{1, 1}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_4, p, t)
}

func Test_Utils_Points_R90_Relative(t *testing.T) {

	origin := &Point{1, 1}
	p := &Point{2, 2}
	expected_1 := &Point{2, 0}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_1, p, t)

	expected_2 := &Point{0, 0}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_2, p, t)

	expected_3 := &Point{0, 2}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_3, p, t)

	expected_4 := &Point{2, 2}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_4, p, t)

}

func Test_Utils_Points_L90(t *testing.T) {

	origin := &Point{0, 0}
	p := &Point{1, 1}
	expected_1 := &Point{-1, 1}
	p.RotateAroundOrigin(-90, origin)
	verifyPoint(expected_1, p, t)

	expected_2 := &Point{-1, -1}
	p.RotateAroundOrigin(-90, origin)
	verifyPoint(expected_2, p, t)

	expected_3 := &Point{1, -1}
	p.RotateAroundOrigin(-90, origin)
	verifyPoint(expected_3, p, t)

	expected_4 := &Point{1, 1}
	p.RotateAroundOrigin(-90, origin)
	verifyPoint(expected_4, p, t)
}

func Test_Utils_Points_R180(t *testing.T) {

	origin := &Point{0, 0}
	p := &Point{1, 1}
	expected_1 := &Point{-1, -1}
	p.RotateAroundOrigin(180, origin)
	verifyPoint(expected_1, p, t)

	expected_2 := &Point{1, 1}
	p.RotateAroundOrigin(180, origin)
	verifyPoint(expected_2, p, t)

}

func Test_Utils_Points_R270(t *testing.T) {

	origin := &Point{0, 0}
	p := &Point{1, 1}
	expected_1 := &Point{-1, 1}
	p.RotateAroundOrigin(270, origin)
	verifyPoint(expected_1, p, t)

	expected_2 := &Point{-1, -1}
	p.RotateAroundOrigin(270, origin)
	verifyPoint(expected_2, p, t)

	expected_3 := &Point{1, -1}
	p.RotateAroundOrigin(270, origin)
	verifyPoint(expected_3, p, t)

	expected_4 := &Point{1, 1}
	p.RotateAroundOrigin(270, origin)
	verifyPoint(expected_4, p, t)

}

func verifyPoint(expected *Point, actual *Point, t *testing.T) {
	if expected.x != actual.x || expected.y != actual.y {
		t.Errorf("verifyPoint %v != %v\n", expected, actual)
	}
}
