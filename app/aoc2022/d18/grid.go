package d18

import (
	"strings"

	"github.com/simonski/goutils"
)

type Grid struct {
	data map[string]*Cube
}

func (g *Grid) Bounds() (int, int, int, int, int, int) {
	max_x := 0
	max_y := 0
	max_z := 0
	min_x := 100
	min_y := 100
	min_z := 100
	for _, c := range g.data {
		max_x = goutils.Max(max_x, c.x)
		max_y = goutils.Max(max_y, c.y)
		max_z = goutils.Max(max_z, c.z)
		min_x = goutils.Min(min_x, c.x)
		min_y = goutils.Min(min_y, c.y)
		min_z = goutils.Min(min_z, c.z)
	}
	return max_x, max_y, max_z, min_x, min_y, min_z
}

func (g *Grid) Add(c *Cube) {
	g.data[c.Key()] = c
}

func (g *Grid) Size() int {
	return len(g.data)
}

func NewGrid(input string) *Grid {
	g := Grid{data: make(map[string]*Cube)}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		c := NewCube(line)
		g.Add(c)
	}
	return &g
}

func (g *Grid) Get(x int, y int, z int) *Cube {
	key := Key(x, y, z)
	return g.data[key]
}

func (g *Grid) GetOrEmpty(x int, y int, z int) *Cube {
	key := Key(x, y, z)
	result := g.data[key]
	if result == nil {
		return &Cube{x: x, y: y, z: z}
	}
	return result
}

func (g *Grid) Contains(x int, y int, z int) bool {
	return g.Get(x, y, z) != nil
}

// returns all neighbouring coordinates that can be directly connected
func (g *Grid) NeighboursConnectable(x int, y int, z int) []*Point3D {
	results := make([]*Point3D, 0)
	results = append(results, NewPoint3D(x-1, y, z))
	results = append(results, NewPoint3D(x+1, y, z))
	results = append(results, NewPoint3D(x, y-1, z))
	results = append(results, NewPoint3D(x, y+1, z))
	results = append(results, NewPoint3D(x, y, z-1))
	results = append(results, NewPoint3D(x, y, z+1))
	return results
}

func (g *Grid) CubesConnected(x int, y int, z int) []*Cube {
	results := make([]*Cube, 0)
	for _, p := range g.NeighboursConnectable(x, y, z) {
		c := g.Get(p.x, p.y, p.z)
		if c != nil {
			results = append(results, c)
		}
	}
	return results
}

// returns number of connected and open sizes
func (g *Grid) CountSides() (int, int) {
	total_connected := 0
	total_not_connected := 0
	for _, c := range g.data {
		connected := len(g.CubesConnected(c.x, c.y, c.z))
		total_connected += connected
		total_not_connected += (6 - connected)
	}
	return total_connected, total_not_connected
}

func (g *Grid) get_external_surface_area() (int, map[*Cube]bool) {
	// """ Determine surface area of all cubes that can reach the outside. """
	cubes_to_outside := make(map[*Cube]bool)   //   # cache cubes we have already identified a path to outside for
	no_path_to_outside := make(map[*Cube]bool) //  set()  # store all internal empty
	surfaces_to_outside := 0

	// # Loop through the cubes and find any that can reach outside
	for _, cube := range g.data {
		for _, p := range g.NeighboursConnectable(cube.x, cube.y, cube.z) {
			neighbour := g.GetOrEmpty(p.x, p.y, p.z)
			if g._has_path_to_outside(neighbour, cubes_to_outside, no_path_to_outside) {
				cubes_to_outside[neighbour] = true
				surfaces_to_outside += 1
			} else {
				no_path_to_outside[neighbour] = true
			}
		}
	}
	return surfaces_to_outside, no_path_to_outside
}

func (g *Grid) _has_path_to_outside(cube *Cube, cubes_to_outside map[*Cube]bool, no_path_to_outside map[*Cube]bool) bool {
	/*
		 Perform BFS to flood fill from this empty cube.
		Param cubes_to_outside is to cache cubes we've seen before, that we know have a path.
		Param internal_cubues is to cache cubes we've seen before, that are internal. """
	*/
	frontier := NewQ()
	frontier.Pushleft(cube)
	explored := make(map[*Cube]bool)
	explored[cube] = true

	max_x, max_y, max_z, min_x, min_y, min_z := g.Bounds()

	for {
		if frontier.Size() == 0 {
			break
		}
		// fmt.Printf("Q size is %v\n", frontier.Size())

		current_cube := frontier.Popleft() // # FIFO for BFS

		// # Check caches
		if cubes_to_outside[current_cube] {
			return true //# We've got out from here before
		}

		if no_path_to_outside[current_cube] {
			continue // # This cube doesn't have a path, so no point checking its neighbours
		}

		if g.Contains(current_cube.x, current_cube.y, current_cube.z) { // } in self.filled_cubes:
			// if filled_cubes[current_cube] { // } in self.filled_cubes:
			continue // # This path is blocked
		}

		// # Check if we've followed a path outside of the bounds
		if current_cube.x > max_x || current_cube.y > max_y || current_cube.z > max_z {
			return true
		}

		if current_cube.x < min_x || current_cube.y < min_y || current_cube.z < min_z {
			return true
		}

		// # We want to look at all neighbours of this empty space
		for _, p := range g.NeighboursConnectable(current_cube.x, current_cube.y, current_cube.z) {
			neighbour := g.GetOrEmpty(p.x, p.y, p.z)
			if !explored[neighbour] {
				frontier.Pushright(neighbour)
				explored[neighbour] = true
			}
		}

		if frontier.Size() == 0 {
			break
		}
	}

	return false

}
