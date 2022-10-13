package aoc2021

import "fmt"

type Scanner struct {
	Beacons []*Beacon
}

// a scanner shoudl be able to provide 24 variations of beacon placement by
// order of beacon position x,y,z - 4 varieties
// sign of each - 6 varieties; 24 varieties of beacon

type Ocean struct {
	Beacons  []*Beacon
	Scanners []*Scanner
}

type Beacon struct {
	X int
	Y int
	Z int
}

func (b *Beacon) Debug() string {
	return fmt.Sprintf("X: %v, Y: %v, Z: %v", b.X, b.Y, b.Z)
}

func (b *Beacon) Permutations() []*Beacon {
	results := make([]*Beacon, 0)
	results = append(results, &Beacon{X: b.X, Y: b.Y, Z: b.Z})
	results = append(results, &Beacon{X: b.X, Y: b.Z, Z: b.Y})
	results = append(results, &Beacon{X: b.Y, Y: b.X, Z: b.Z})
	results = append(results, &Beacon{X: b.Y, Y: b.Z, Z: b.X})
	results = append(results, &Beacon{X: b.Z, Y: b.X, Z: b.Y})
	results = append(results, &Beacon{X: b.Z, Y: b.Y, Z: b.X})

	r := make([]*Beacon, 0)
	for _, b := range results {
		r = append(r, b.Rotate()...)
		// for _, p := range b.Rotate() {
		// 	r = append(r, p)
		// }
	}
	// r2 := v[0], v[2], v[1]
	// r3 := v[1], v[0], v[2]
	// r4 := v[1], v[2], v[0]
	// r5 := v[2], v[0], v[1]
	// r6 := v[2], v[1], v[0]

	return r
}

func (b *Beacon) Rotate() []*Beacon {
	// 1 x > -y, y > x, z->z
	// 2 x > -x, y -> -y, z -> z
	// 3 x > y, y -> -z, z -> z
	results := make([]*Beacon, 0)
	b1 := b
	b2 := &Beacon{b.Y, -b.X, b.Z}
	b3 := &Beacon{-b.X, -b.Y, b.Z}
	b4 := &Beacon{-b.Y, b.X, b.Z}

	results = append(results, b1)
	results = append(results, b2)
	results = append(results, b3)
	results = append(results, b4)
	return results

}

/*
For example, if a scanner is at x,y,z coordinates 500,0,-500 and there are beacons at -500,1000,-1500 and 1501,0,-500, the scanner could report that the first beacon is at -1000,1000,-1000 (relative to the scanner) but would not detect the second beacon at all.
*/

/*
Each scanner is capable of detecting all beacons in a large cube centered on the scanner; beacons that are at most 1000 units away from the scanner in each of the three axes (x, y, and z) have their precise position determined relative to the scanner.
*/

/*
Unfortunately, while each scanner can report the positions of all detected beacons relative to itself, the scanners do not know their own position. You'll need to determine the positions of the beacons and scanners yourself.
*/

/*
The scanners and beacons map a single contiguous 3d region. This region can be reconstructed by finding pairs of scanners that have overlapping detection regions such that there are at least 12 beacons that both scanners detect within the overlap. By establishing 12 common beacons, you can precisely determine where the scanners are relative to each other, allowing you to reconstruct the beacon map one scanner at a time.
*/

func (s *Scanner) Scan() []*Beacon {
	results := make([]*Beacon, 0)
	return results
}
