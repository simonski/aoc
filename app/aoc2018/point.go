package aoc2018

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	line       string
	position_x int
	position_y int
	velocity_x int
	velocity_y int
	Key        string
}

func (p *Point) Debug() {
	fmt.Printf("line=%v, px=%v, py=%v, vx=%v, vy=%v\n", p.line, p.position_x, p.position_y, p.velocity_x, p.velocity_y)
}

func (p *Point) Remap() {
}

func (p *Point) Step(remap bool) {
	// p.step += 1
	p.position_x += p.velocity_x
	p.position_y += p.velocity_y
	if remap {
		p.Key = fmt.Sprintf("%v.%v", p.position_x, p.position_y)
	}

	// p.position_x = applyrange(p.position_x, p.velocity_x, g.min_x, g.max_x)
	// p.position_y = applyrange(p.position_y, p.velocity_y, g.min_y, g.max_y)
}

func NewPoint(line string) *Point {
	// position=< 9,  1> velocity=< 0,  2>
	splits := strings.Split(line, "velocity")
	position := strings.Split(splits[0], "=")[1]
	position = strings.Replace(position, "position=", "", -1)
	position = strings.Replace(position, "<", "", -1)
	position = strings.Replace(position, ">", "", -1)
	position = strings.Replace(position, " ", "", -1)
	p := strings.Split(position, ",")
	pos_x, _ := strconv.Atoi(p[0])
	pos_y, _ := strconv.Atoi(p[1])

	velocity := strings.Split(splits[1], "=")[1]
	velocity = strings.Replace(velocity, "velocity=", "", -1)
	velocity = strings.Replace(velocity, "<", "", -1)
	velocity = strings.Replace(velocity, ">", "", -1)
	velocity = strings.Replace(velocity, " ", "", -1)
	v := strings.Split(velocity, ",")
	v_x, _ := strconv.Atoi(v[0])
	v_y, _ := strconv.Atoi(v[1])

	return &Point{line: line, position_x: pos_x, position_y: pos_y, velocity_x: v_x, velocity_y: v_y}
}
