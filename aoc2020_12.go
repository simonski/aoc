package main

/*

https://adventofcode.com/2020/day/12

--- Day 12: Rain Risk ---
Your ferry made decent progress toward the island, but the storm came in faster than anyone expected. The ferry needs to take evasive actions!

Unfortunately, the ship's navigation computer seems to be malfunctioning; rather than giving a route directly to safety, it produced extremely circuitous instructions. When the captain uses the PA system to ask if anyone can help, you quickly volunteer.

The navigation instructions (your puzzle input) consists of a sequence of single-character actions paired with integer input values. After staring at them for a few minutes, you work out what they probably mean:

Action N means to move north by the given value.
Action S means to move south by the given value.
Action E means to move east by the given value.
Action W means to move west by the given value.
Action L means to turn left the given number of degrees.
Action R means to turn right the given number of degrees.
Action F means to move forward by the given value in the direction the ship is currently facing.
The ship starts by facing east. Only the L and R actions change the direction the ship is facing. (That is, if the ship is facing east and the next instruction is N10, the ship would move north 10 units, but would still move east if the following action were F.)

For example:

F10
N3
F7
R90
F11
These instructions would be handled as follows:

F10 would move the ship 10 units east (because the ship starts by facing east) to east 10, north 0.
N3 would move the ship 3 units north to east 10, north 3.
F7 would move the ship another 7 units east (because the ship is still facing east) to east 17, north 3.
R90 would cause the ship to turn right by 90 degrees and face south; it remains at east 17, north 3.
F11 would move the ship 11 units south to east 17, south 8.
At the end of these instructions, the ship's Manhattan distance (sum of the absolute values of its east/west position and its north/south position) from its starting position is 17 + 8 = 25.

Figure out where the navigation instructions lead. What is the Manhattan distance between that location and the ship's starting position?


*/
import (
	"fmt"
	"strconv"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_12 is the entrypoint
func AOC_2020_12(cli *goutils.CLI) {
	// AOC_2020_12_part1_attempt1(cli)
	AOC_2020_12_part2_attempt1(cli)
}

func AOC_2020_12_part1_attempt1(cli *goutils.CLI) {
	filename := cli.GetFileExistsOrDie("-input")
	p := NewPathFromFile(filename)
	// p.Debug()

	s := NewShip()
	for index, m := range p.movements {
		s.Execute(m)
		fmt.Printf("[%v] %v%v  -> Ship[x=%v, y=%v, a=%v]\n", index, m.Command, m.Value, s.x, s.y, s.angle)
	}
}

func AOC_2020_12_part2_attempt1(cli *goutils.CLI) {
	filename := cli.GetFileExistsOrDie("-input")
	p := NewPathFromFile(filename)
	s := NewShip2(10, 1)
	fmt.Printf("START\n  Ship2(%v,%v %v) WP (%v,%v)\n", s.x, s.y, s.angle, s.waypoint.x, s.waypoint.y)
	for index, m := range p.movements {
		s.Execute(m)
		wp := s.waypoint
		fmt.Printf("[%v] %v%v  -> Ship(%v,%v %v), -> Wp(%v,%v)\n", index, m.Command, m.Value, s.x, s.y, s.angle, wp.x, wp.y)
	}
}

type Path struct {
	movements []*Movement
}

func (p *Path) Debug() {
	for index, m := range p.movements {
		fmt.Printf("[%v] %v%v\n", index, m.Command, m.Value)
	}
}

type Movement struct {
	Command string
	Value   int
}

type Ship struct {
	x     int
	y     int
	angle int
}

type Ship2 struct {
	x        int
	y        int
	angle    int
	waypoint *Waypoint
}

type Waypoint struct {
	x int
	y int
}

func NewShip() *Ship {
	return &Ship{x: 0, y: 0, angle: 90}
}

func (s *Ship) Execute(m *Movement) {
	if m.Command == "N" {
		s.y += m.Value
	}
	if m.Command == "S" {
		s.y -= m.Value
	}
	if m.Command == "E" {
		s.x += m.Value
	}
	if m.Command == "W" {
		s.x -= m.Value
	}
	if m.Command == "R" {
		s.angle += m.Value
		if s.angle >= 360 {
			s.angle = s.angle - 360
		}
	}
	if m.Command == "L" {
		s.angle -= m.Value
		if s.angle < 0 {
			s.angle = 360 + s.angle
		}
	}
	if m.Command == "F" {
		angle := s.angle
		if angle == 0 {
			s.y += m.Value
		} else if angle == 180 {
			s.y -= m.Value
		} else if angle == 90 {
			s.x += m.Value
		} else {
			s.x -= m.Value
		}
	}
}
func NewShip2(waypoint_x int, waypoint_y int) *Ship2 {
	wp := Waypoint{x: waypoint_x, y: waypoint_y}
	return &Ship2{x: 0, y: 0, angle: 90, waypoint: &wp}
}

func (s *Ship2) Execute(m *Movement) {
	if m.Command == "N" {
		s.waypoint.y += m.Value
	}
	if m.Command == "S" {
		s.waypoint.y -= m.Value
	}
	if m.Command == "E" {
		s.waypoint.x += m.Value
	}
	if m.Command == "W" {
		s.waypoint.x -= m.Value
	}

	if m.Command == "R" {
		wp := &Point2D{s.waypoint.x, s.waypoint.y}
		origin := &Point2D{s.x, s.y}
		wp.RotateAroundOrigin(m.Value, origin)
		s.waypoint.x = wp.x
		s.waypoint.y = wp.y

	} else if m.Command == "L" {
		wp := &Point2D{s.waypoint.x, s.waypoint.y}
		origin := &Point2D{s.x, s.y}
		wp.RotateAroundOrigin(-m.Value, origin)
		s.waypoint.x = wp.x
		s.waypoint.y = wp.y
	}

	if m.Command == "F" {
		x_diff := s.waypoint.x - s.x
		y_diff := s.waypoint.y - s.y

		x := (x_diff * m.Value)
		y := (y_diff * m.Value)

		fmt.Printf("Movement is %v%v, x/y diff is %v,%v, total changes will be adding %v,%v\n", m.Command, m.Value, x_diff, y_diff, x, y)

		s.x += x
		s.y += y

		fmt.Printf("New s.x/s.y %v,%v\n", s.x, s.y)

		s.waypoint.x = s.x + x_diff
		s.waypoint.y = s.y + y_diff
	}
}

func NewPathFromFile(filename string) *Path {
	lines := load_file_to_strings(filename)
	data := make([]*Movement, 0)
	for _, line := range lines {
		command := line[0:1]
		value := line[1:]
		ivalue, _ := strconv.Atoi(value)
		movement := &Movement{Command: command, Value: ivalue}
		data = append(data, movement)
	}
	path := Path{movements: data}
	return &path
}
