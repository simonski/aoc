package aoc2021

type PositionXYZ struct {
	X int
	Y int
	Z int
}
type Beacon struct {
	PositionXYZ
}

type Scanner struct {
	PositionXYZ
	PositionCalculated bool
}

type Ocean struct {
	Beacons  []*Beacon
	Scanners []*Scanner
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
