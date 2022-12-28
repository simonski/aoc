package d18

import (
	"strconv"
	"strings"
)

type Cube struct {
	x int
	y int
	z int
}

func (c *Cube) Key() string {
	return Key(c.x, c.y, c.z)
}

func (c *Cube) String() string {
	return c.Key()
}

func NewCube(line string) *Cube {
	splits := strings.Split(line, ",")
	x, _ := strconv.Atoi(splits[0])
	y, _ := strconv.Atoi(splits[1])
	z, _ := strconv.Atoi(splits[2])
	c := Cube{x: x, y: y, z: z}
	return &c
}
