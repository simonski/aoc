package d18

import (
	"fmt"
	"testing"
)

/*
Style 1:

-- To find the exterior surface area we can start with a 'solid' volume and remove just those that are part of the exterior empty volume of the scanned space.

-- To find the exterior empty cells, we just start at some known exterior voxel and search for all connected empty voxels. Because we have added an exterior boundary of empty voxels to the scan we can guarantee the search will find all connected empty voxels.

Style 2:

Find all adjacent cubes to our _filled” cubes. These are either:
1. Part of an internal pocket. If we flood fill a pocket, it will have a nearby boundary.
2. Part of path to the outside. If we flood fill, we will eventually reach a cube beyond all the droplet bounds.

To solve:
1. For each filled cube, get its adjacent cubes.
	Perform a Breadth-First Search (BFS) from each adacent, if the adjacent cube is empty space.
	INTERNAL: If the BFS only leads to filled cubes, then all paths are blocked, so this cube is internal.
	EXTERNAL: If the BFS leads to cubes that our outside of our bounds, then this cube has a path to the outside. Thus, this cube counts as external.

Store all paths to cache the BFS.
Only increment the surface area count every time we find an adjacent location that has a path out, i.e. only for cubes that are external.

The implementation of the BFS is standard, and we’ve covered this before. The only other thing worth noting here is how we define the bounds that we use to determine if our BFS has reached “outside”. To set these bounds, we simply measure the minimum and maximum values of each of x, y, and z, from our set of filled cubes. If our BFS ever reaches a cube where any of the cube coordinates are beyond these boundaries, then we know we’ve reached “empty space” outside of the perimeter of our droplet.
*/

func Test_ExteriorSurface_Test(t *testing.T) {
	g := NewGrid(TEST_DATA)
	external, _ := g.get_external_surface_area()
	_, not_connected := g.CountSides()
	fmt.Println("TEST")
	fmt.Printf("total_surface_area : %v\nget_external_size  : %v\n", not_connected, external)

	fmt.Println()
	fmt.Println("REAL")
	g2 := NewGrid(REAL_DATA)
	external2, _ := g2.get_external_surface_area()
	_, not_connected2 := g2.CountSides()
	fmt.Printf("total_surface_area : %v\nget_external_size  : %v\n", not_connected2, external2)

}
