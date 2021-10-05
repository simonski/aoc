package main

import (
	"testing"
)

func Test_AOC2020_12_TestExample(t *testing.T) {

	// start
	// wp 10, 1
	s := NewShip2(10, 1)
	verifyShipAndWaypoint("0/5 Starting", 0, 0, 10, 1, s, t)
	s.Execute(&Movement{Command: "F", Value: 10})
	verifyShipAndWaypoint("1/5 F10", 100, 10, 110, 11, s, t)

	s.Execute(&Movement{Command: "N", Value: 3})
	verifyShipAndWaypoint("2/5 N3", 100, 10, 110, 14, s, t)

	s.Execute(&Movement{Command: "F", Value: 7})
	verifyShipAndWaypoint("3/5 F7", 170, 38, 180, 42, s, t)

	s.Execute(&Movement{Command: "R", Value: 90})
	verifyShipAndWaypoint("4/5 R90", 170, 38, 174, 28, s, t)

	s.Execute(&Movement{Command: "F", Value: 11})
	verifyShipAndWaypoint("5/5 F11", 214, -72, 218, -82, s, t)
}

func verifyShipAndWaypoint(message string, x_ship int, y_ship int, x_wp int, y_wp int, s *Ship2, t *testing.T) {
	if x_ship != s.X {
		t.Errorf("[%v] ship.x %v!=%v\n", message, x_ship, s.X)
	}
	if y_ship != s.Y {
		t.Errorf("[%v] ship.y %v!=%v\n", message, y_ship, s.Y)
	}
	if x_wp != s.Waypoint.X {
		t.Errorf("[%v] ship.wp.x %v!=%v\n", message, x_wp, s.Waypoint.X)
	}
	if y_wp != s.Waypoint.Y {
		t.Errorf("[%v] ship.wp.y %v!=%v\n", message, y_wp, s.Waypoint.Y)
	}
}
