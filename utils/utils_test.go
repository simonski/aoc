package utils

import (
	"testing"
)

func Test_Utils_Points_R90(t *testing.T) {

	origin := &Point2D{0, 0}
	p := &Point2D{1, 1}
	expected_1 := &Point2D{1, -1}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_1, p, t)

	expected_2 := &Point2D{-1, -1}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_2, p, t)

	expected_3 := &Point2D{-1, 1}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_3, p, t)

	expected_4 := &Point2D{1, 1}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_4, p, t)
}

func Test_Utils_Points_R90_Relative(t *testing.T) {

	origin := &Point2D{1, 1}
	p := &Point2D{2, 2}
	expected_1 := &Point2D{2, 0}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_1, p, t)

	expected_2 := &Point2D{0, 0}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_2, p, t)

	expected_3 := &Point2D{0, 2}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_3, p, t)

	expected_4 := &Point2D{2, 2}
	p.RotateAroundOrigin(90, origin)
	verifyPoint(expected_4, p, t)

}

func Test_Utils_Points_L90(t *testing.T) {

	origin := &Point2D{0, 0}
	p := &Point2D{1, 1}
	expected_1 := &Point2D{-1, 1}
	p.RotateAroundOrigin(-90, origin)
	verifyPoint(expected_1, p, t)

	expected_2 := &Point2D{-1, -1}
	p.RotateAroundOrigin(-90, origin)
	verifyPoint(expected_2, p, t)

	expected_3 := &Point2D{1, -1}
	p.RotateAroundOrigin(-90, origin)
	verifyPoint(expected_3, p, t)

	expected_4 := &Point2D{1, 1}
	p.RotateAroundOrigin(-90, origin)
	verifyPoint(expected_4, p, t)
}

func Test_Utils_Points_R180(t *testing.T) {

	origin := &Point2D{0, 0}
	p := &Point2D{1, 1}
	expected_1 := &Point2D{-1, -1}
	p.RotateAroundOrigin(180, origin)
	verifyPoint(expected_1, p, t)

	expected_2 := &Point2D{1, 1}
	p.RotateAroundOrigin(180, origin)
	verifyPoint(expected_2, p, t)

}

func Test_Utils_Points_R270(t *testing.T) {

	origin := &Point2D{0, 0}
	p := &Point2D{1, 1}
	expected_1 := &Point2D{-1, 1}
	p.RotateAroundOrigin(270, origin)
	verifyPoint(expected_1, p, t)

	expected_2 := &Point2D{-1, -1}
	p.RotateAroundOrigin(270, origin)
	verifyPoint(expected_2, p, t)

	expected_3 := &Point2D{1, -1}
	p.RotateAroundOrigin(270, origin)
	verifyPoint(expected_3, p, t)

	expected_4 := &Point2D{1, 1}
	p.RotateAroundOrigin(270, origin)
	verifyPoint(expected_4, p, t)

}

func verifyPoint(expected *Point2D, actual *Point2D, t *testing.T) {
	if expected.X != actual.X || expected.Y != actual.Y {
		t.Errorf("verifyPoint %v != %v\n", expected, actual)
	}
}

func Test_decimal_to_binary(t *testing.T) {
	value := 1
	actual := Decimal_to_binary(1)
	expected := "000000000000000000000000000000000001"
	if actual != expected {
		t.Errorf("decimal_to_binary, for %v expected %v got %v\n", value, expected, actual)
	}

	actual = Decimal_to_binary(8)
	expected = "000000000000000000000000000000001000"
	if actual != expected {
		t.Errorf("decimal_to_binary, for %v expected %v got %v\n", value, expected, actual)
	}

}

func Test_binary_to_decimal(t *testing.T) {
	value := "000000000000000000000000000000000001"
	actual := Binary_to_decimal(value)
	expected := int64(1)
	if actual != expected {
		t.Errorf("binary_to_decimal, for %v expected %v got %v\n", value, expected, actual)
	}

	value = "000000000000000000000000000000001000"
	actual = Binary_to_decimal(value)
	expected = int64(8)
	if actual != expected {
		t.Errorf("binary_to_decimal, for %v expected %v got %v\n", value, expected, actual)
	}

}
