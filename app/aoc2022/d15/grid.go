package d15

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/simonski/goutils"
)

type Point struct {
	X      int
	Y      int
	Beacon *Beacon
	Sensor *Sensor
}

func NewPoint(input string) *Point {
	p := Point{}
	splits := strings.Split(input, ",")
	x, _ := strconv.Atoi(splits[0])
	y, _ := strconv.Atoi(splits[1])
	p.X = x
	p.Y = y
	return &p
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v,%v)", p.X, p.Y)
}

func (p *Point) Distance(p2 *Point) int {
	return p.DistanceTo(p2.X, p2.Y)
}

func (p *Point) DistanceTo(x int, y int) int {
	xdist := p.X - x
	if xdist < 0 {
		xdist *= -1
	}

	ydist := p.Y - y
	if ydist < 0 {
		ydist *= -1
	}
	return xdist + ydist
}

type Grid struct {
	data       map[string]*Point
	MinX       int
	MaxX       int
	MaxY       int
	MinY       int
	Beacons    []*Beacon
	Sensors    []*Sensor
	x_detected map[int]bool
}

func NewGrid(data string) *Grid {
	g := Grid{}
	g.data = make(map[string]*Point)
	g.MinX = 1000000
	g.MinY = 1000000
	g.Beacons = make([]*Beacon, 0)
	g.Sensors = make([]*Sensor, 0)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		line = strings.ReplaceAll(line, "Sensor at x=", "")
		line = strings.ReplaceAll(line, "y=", "")
		line = strings.ReplaceAll(line, "closest beacon is at x=", "")
		line = strings.ReplaceAll(line, " ", "")
		splits := strings.Split(line, ":") // should now be 20,1:30,4 or [ 20,1 43,2 ]
		sensor := NewSensor(splits[0])
		beacon := NewBeacon(splits[1])
		sensor.Beacon = beacon
		sensor.Strength = sensor.Point.Distance(beacon.Point)
		g.AddSensor(sensor)
		g.AddBeacon(beacon)
		// fmt.Printf("line was '%v', sensor=%v, beacon=%v, new bounds: %v,%v -> %v,%v\n", line, sensor, beacon, g.MinX, g.MinY, g.MaxX, g.MaxY)
	}
	return &g

}

func (g *Grid) Bounds() (int, int, int, int) {
	return g.MinX, g.MinY, g.MaxX, g.MaxY
}

func (g *Grid) Width() int {
	if g.MinX < 0 {
		return g.MaxX + (g.MinX * -1)
	} else {
		return g.MaxX - g.MinX
	}
}

func (g *Grid) Height() int {
	if g.MinY < 0 {
		return g.MaxY + (g.MinY * -1)
	} else {
		return g.MaxY - g.MinY
	}
}

func (g *Grid) Debug() string {
	results := ""
	for y := g.MinY; y < g.MaxY; y++ {
		line := ""
		for x := g.MinX; x < g.MaxX; x++ {
			p := g.Get(x, y)
			if p == nil {
				line = fmt.Sprintf("%v.", line)
			} else {
				if p.Beacon != nil {
					line = fmt.Sprintf("%vB", line)
				} else if p.Sensor != nil {
					line = fmt.Sprintf("%vS", line)
				}
			}
		}
		if results != "" {
			results = fmt.Sprintf("%v\n%v", results, line)
		} else {
			results = line
		}
	}
	return results
}

func (g *Grid) AddSensor(sensor *Sensor) {
	g.Sensors = append(g.Sensors, sensor)
	g.MaxX = goutils.Max(g.MaxX, sensor.Point.X)
	g.MinX = goutils.Min(g.MinX, sensor.Point.X)
	g.MaxY = goutils.Max(g.MaxY, sensor.Point.Y)
	g.MinY = goutils.Min(g.MinY, sensor.Point.Y)
	g.Put(sensor.Point.X, sensor.Point.Y, sensor.Point)
}

func (g *Grid) AddBeacon(beacon *Beacon) {
	g.Beacons = append(g.Beacons, beacon)
	g.MaxX = goutils.Max(g.MaxX, beacon.Point.X)
	g.MinX = goutils.Min(g.MinX, beacon.Point.X)
	g.MaxY = goutils.Max(g.MaxY, beacon.Point.Y)
	g.MinY = goutils.Min(g.MinY, beacon.Point.Y)
	g.Put(beacon.Point.X, beacon.Point.Y, beacon.Point)
}

func (g *Grid) Put(x int, y int, p *Point) {
	key := fmt.Sprintf("%v_%v", x, y)
	g.MaxX = goutils.Max(g.MaxX, y)
	g.MinX = goutils.Min(g.MinX, x)
	g.MaxY = goutils.Max(g.MaxY, y)
	g.MinY = goutils.Min(g.MinY, y)
	g.data[key] = p
}

func (g *Grid) Get(x int, y int) *Point {
	key := fmt.Sprintf("%v_%v", x, y)
	return g.data[key]
}

func (g *Grid) Contains(x int, y int, p *Point) bool {
	return g.Get(x, y) != nil
}

func (g *Grid) CountCannotsForRow(row int) (int, int) {
	// go over each point left to right

	// remove any sensors whose distance is too great - that is,
	// it cannot reach the point even under the best of circumstances (a straight line to it)
	sensors := make([]*Sensor, 0)
	for _, sensor := range g.Sensors {
		strength := sensor.Strength
		distance := sensor.Point.DistanceTo(sensor.Point.X, row)
		if distance <= strength {
			// it has the distance
			sensors = append(sensors, sensor)
			fmt.Printf("Sensor %v strength %v could detect (distance is %v)\n", sensor, sensor.Strength, distance)
		} else {
			fmt.Printf("Sensor %v strength %v cannot detect (distance is %v)\n", sensor, sensor.Strength, distance)
		}
	}
	// sensors := g.Sensors
	fmt.Printf("We will search %v out of %v sensors.\n", len(sensors), len(g.Sensors))
	// fmt.Println(g.Debug())

	// line := ""
	could_be_beacon := 0
	could_not_be_beacon := 0
	iteration := 0
	totalIterations := g.Width() * len(sensors)
	for x := g.MinX; x <= g.MaxX; x++ {
		p := g.Get(x, row)
		within_signal_strength := false
		if p != nil {

		} else {
			// if p == nil {
			// could be
			for _, sensor := range sensors {
				iteration++
				fmt.Printf("%v/%v\n", iteration, totalIterations)
				distance := sensor.Point.DistanceTo(x, row)
				if distance <= sensor.Strength {
					// fmt.Printf("Sensor %v can detect (strength=%v) checking point (%v,%v), distance=%v\n", sensor, sensor.Strength, x, row, distance)
					within_signal_strength = true
					break
				} else {
					// fmt.Printf("Sensor %v cannot detect (strength=%v) checking point (%v,%v), distance=%v\n", sensor, sensor.Strength, x, row, distance)

				}

			}
		}

		if within_signal_strength {
			// fmt.Printf("(%v,%v) IS detectable by a sensor\n", x, row)
			could_not_be_beacon += 1
			// line += "."
		} else {
			// fmt.Printf("(%v,%v) IS NOT detectable by any sensor\n", x, row)
			could_be_beacon += 1
			// line += "#"
		}

		// line += "."
		// } else if p != nil && p.Sensor != nil {
		// 	// it's a sensor
		// 	line += "S"
		// } else if p != nil && p.Beacon != nil {
		// 	// it's a beacon
		// 	line += "B"
		// }

	}
	return could_be_beacon, could_not_be_beacon // , line
}

func (g *Grid) CountCannotsForRow_V2(row int) (int, int, map[int]bool, int) {
	// go over each point left to right

	// remove any sensors whose distance is too great - that is,
	// it cannot reach the point even under the best of circumstances (a straight line to it)
	// sensors := make([]*Sensor, 0)
	// for _, sensor := range g.Sensors {
	// 	strength := sensor.Strength
	// 	distance := sensor.Point.DistanceTo(sensor.Point.X, row)
	// 	if distance <= strength {
	// 		// it has the distance
	// 		sensors = append(sensors, sensor)
	// 	} else {
	// 	}
	// }

	result := make(map[int]bool)

	min_x := math.MaxInt
	for _, sensor := range g.Sensors {

		// for each sensor, make a straight line to the row as the shortest distance
		// then we can go left by the different and right by the difference
		// x = sensor.Point.X
		distance := sensor.Point.DistanceTo(sensor.Point.X, row)
		strength := sensor.Strength
		difference := strength - distance
		if difference <= 0 {
			continue
		}
		difference = goutils.Abs(difference)
		for x := sensor.Point.X - difference; x < sensor.Point.X+difference; x++ {
			p := g.Get(x, row)
			if p == nil {
				result[x] = true /// true means we can detect it, so it cannot be a signal
			}
		}
		min_x = goutils.Min(min_x, sensor.Point.X-difference)
	}

	known := 0
	unknown := 0
	for _, value := range result {
		if value {
			known++
		} else {
			unknown++
		}
	}
	// known -= removeSensors

	// for x := sensor.Point.X - difference; x < sensor.Point.X+difference; x++ {
	// 	p := g.Get(x, row)
	// 	if p == nil {
	// 		result[x] = true /// true means we can detect it, so it cannot be a signal
	// 	}
	// }

	// if within_signal_strength {
	// 	// fmt.Printf("(%v,%v) IS detectable by a sensor\n", x, row)
	// 	could_not_be_beacon += 1
	// 	// line += "."
	// } else {
	// 	// fmt.Printf("(%v,%v) IS NOT detectable by any sensor\n", x, row)
	// 	could_be_beacon += 1
	// 	// line += "#"
	// }

	// line += "."
	// } else if p != nil && p.Sensor != nil {
	// 	// it's a sensor
	// 	line += "S"
	// } else if p != nil && p.Beacon != nil {
	// 	// it's a beacon
	// 	line += "B"
	// }

	return unknown, known, result, min_x // could_be_beacon, could_not_be_beacon // , line
}

type Segment struct {
	MinX int
	MaxX int
}

func NewSegment(minx int, maxx int) *Segment {
	return &Segment{MinX: minx, MaxX: maxx}
}

func (s *Segment) String() string {
	return fmt.Sprintf("(%v <-> %v)", s.MinX, s.MaxX)
}

func (s *Segment) IsInsideOf(other *Segment) bool {
	return s.MinX >= s.MinX && s.MaxX <= other.MaxX
}

func (s *Segment) OverlapsToLeftOf(other *Segment) bool {
	return s.MinX < other.MinX && s.MaxX >= other.MinX && s.MaxX <= other.MaxX
}

func (s *Segment) OverlapsToRightOf(other *Segment) bool {
	return s.MinX >= s.MinX && s.MinX <= other.MaxX && s.MaxX > other.MaxX
}

func (s *Segment) IsOverlappedBy(other *Segment) bool {
	return s.MinX > other.MinX && s.MaxX < other.MaxX
}

func (s *Segment) OverlapsWith(other *Segment) bool {
	return s.IsInsideOf(other) || s.OverlapsToLeftOf(other) || s.OverlapsToRightOf(other) || s.IsOverlappedBy(other)
}

func (s *Segment) Integrate(VERBOSE bool, other *Segment) bool {
	if other.IsInsideOf(s) {
		// nothing required
		if VERBOSE {
			fmt.Println("s2 is inside s1, do nothing return true")
		}
		return true
	}
	if s.IsOverlappedBy(other) {
		if VERBOSE {
			fmt.Println("s1 is ovelapped by s2, set s1 bounds to be s2 bounds, return true")
		}
		s.MinX = other.MinX
		s.MaxX = other.MaxX
		return true
	}
	if other.OverlapsToLeftOf(s) {
		if VERBOSE {
			fmt.Println("s1 is ovelapped to the left by s2, set s1 left bounds to be s2 left bounds, return true")
		}
		s.MinX = other.MinX
		return true
	}
	if other.OverlapsToRightOf(s) {
		if VERBOSE {
			fmt.Println("s1 is ovelapped to the right by s2, set s1 right bounds to be s2 right bounds, return true")
		}
		s.MaxX = other.MaxX
		return true
	}
	if VERBOSE {
		fmt.Println("s1 is not overlapped at all by s2, return false")
	}
	return false
}

func (g *Grid) CountMissing(VERBOSE bool, row int) []int {
	// for this row, go over each point left to right
	// for each sensor that can reach the row
	// figure the range of X it CANNOT reach (min-x, max-x)
	// if a sensor overlaps, merge the segments
	// eventually we have all sensors and segments or ranges that are inaccesable.
	// now you can calculate the # missing and their as the gaps

	// remove any sensors whose distance is too great anyway - that is,
	// it cannot reach the point even under the best of circumstances (a straight line to it)

	segments := g.GetSegments(VERBOSE, row)

	// each segment shoudl be its own region
	if VERBOSE {
		line := ""
		for index := 0; index < len(segments)-1; index++ {
			segment := segments[index]
			line += segment.String() + ", "
		}
		fmt.Println(line)
	}

	result := make([]int, 0)
	for index := 0; index < len(segments)-1; index++ {
		s1 := segments[index]
		s2 := segments[index+1]
		if VERBOSE {
			fmt.Printf("segment[%v], min=%v, max=%v, segment[%v], min=%v, max=%v\n", index, s1.MinX, s1.MaxX, index+1, s2.MinX, s2.MaxX)
		}
		if s2.MinX-s1.MaxX > 2 {
			break
		}

		for x := s1.MaxX + 1; x < s2.MinX; x++ {
			if VERBOSE {
				fmt.Printf("%v,%v\n", x, row)
			}
			if g.Get(x, row) == nil {
				fmt.Printf("Candidate (%v,%v) has no entry in grid, adding.\n", x, row)
				result = append(result, x)
			} else {
				fmt.Printf("Candidate (%v,%v) has an entry in grid %v, adding.\n", x, row, g.Get(x, row))
			}
		}
	}

	return result
}

func (g *Grid) GetSegments(VERBOSE bool, row int) []*Segment {
	// for this row, go over each point left to right
	// for each sensor that can reach the row
	// figure the range of X it CANNOT reach (min-x, max-x)
	// if a sensor overlaps, merge the segments
	// eventually we have all sensors and segments or ranges that are inaccesable.
	// now you can calculate the # missing and their as the gaps

	// remove any sensors whose distance is too great anyway - that is,
	// it cannot reach the point even under the best of circumstances (a straight line to it)

	segments := make([]*Segment, 0)
	for _, sensor := range g.Sensors {
		strength := sensor.Strength
		sx := sensor.Point.X
		// sy := sensor.Point.X
		// sy := sensor.Point.Y
		distance := sensor.Point.DistanceTo(sx, row)
		if distance <= strength {
			// it has has the distance
			difference := strength - distance
			sensor_min_x := sx - difference
			sensor_max_x := sx + difference
			// if VERBOSE {
			// 	fmt.Printf("Sensor (%v,%v), strength=%v, min_x=%v, max_x=%v\n", sx, sy, strength, sensor_min_x, sensor_max_x)
			// }
			segment := NewSegment(sensor_min_x, sensor_max_x)
			segments = append(segments, segment)

		} else {
			// it does not have the distance

		}

	}

	// now we have lots of segments for this row we can try to integrate them all together
	// now sort the segemnts by minX
	sort.Slice(segments, func(i int, j int) bool {
		s1 := segments[i]
		s2 := segments[j]
		return s1.MinX < s2.MinX
	})

	if VERBOSE {
		fmt.Printf("Segments[%v] (before integration) : %v\n", row, segments)
	}
	// now integrate until we can no longer integrate
	if len(segments) > 1 {
		for {
			indexToRemove := -1
			integrated := false
			for index := 0; index < len(segments)-1; index++ {
				s1 := segments[index]
				s2 := segments[index+1]
				if s1.Integrate(VERBOSE, s2) {
					// remove s2 and start again
					integrated = true
					indexToRemove = index + 1
				} else {
					// move on attempting to integrate
				}
			}
			if integrated {
				// remove the index we integrated and start again
				segments = RemoveSegmentAtIndex(segments, indexToRemove)
			} else {
				break
			}
		}
	}

	if VERBOSE {
		fmt.Printf("Segments[%v] (afer integration)  : %v\n", row, segments)
	}

	return segments

}

func RemoveSegmentAtIndex(segments []*Segment, removeIndex int) []*Segment {
	result := make([]*Segment, 0)
	for index := 0; index < len(segments); index++ {
		if index != removeIndex {
			result = append(result, segments[index])
		}
	}
	return result
}

func (g *Grid) CountGaps(VERBOSE bool, row int) (int, int) {
	// for this row, go over each point left to right
	// for each sensor that can reach the row
	// figure the range of X it CANNOT reach (min-x, max-x)
	// if a sensor overlaps, merge the segments
	// eventually we have all sensors and segments or ranges that are inaccesable.
	// now you can calculate the # missing and their as the gaps

	// remove any sensors whose distance is too great anyway - that is,
	// it cannot reach the point even under the best of circumstances (a straight line to it)

	segments := g.GetSegments(VERBOSE, row)

	// each segment should be its own region
	if VERBOSE {
		fmt.Printf("Segments for row %v: %v\n", row, segments)
	}

	// now count the 'gaps'
	result := 0
	col := -1
	for index := 0; index < len(segments)-1; index++ {
		s1 := segments[index]
		s2 := segments[index+1]
		gap := s2.MinX - s1.MaxX - 1
		col = s1.MaxX + 1
		if gap > 0 {
			result += gap
		}
		if VERBOSE {
			fmt.Printf("segment[%v], min=%v, max=%v, segment[%v], min=%v, max=%v\n", index, s1.MinX, s1.MaxX, index+1, s2.MinX, s2.MaxX)
			fmt.Printf("gap (%v - %v) = %v\n", s2.MinX, s1.MaxX, gap)
		}

	}

	return result, col
}

func (g *Grid) CountCannotBePresent(VERBOSE bool, row int) int {
	// for this row, go over each point left to right
	// for each sensor that can reach the row
	// figure the range of X it CANNOT reach (min-x, max-x)
	// if a sensor overlaps, merge the segments
	// eventually we have all sensors and segments or ranges that are inaccesable.
	// now you can calculate the # missing and their as the gaps

	// remove any sensors whose distance is too great anyway - that is,
	// it cannot reach the point even under the best of circumstances (a straight line to it)

	segments := g.GetSegments(VERBOSE, row)

	// each segment should be its own region
	if VERBOSE {
		fmt.Printf("Segments for row %v: %v\n", row, segments)
	}

	// now count the total length occupied - this is the area that cannot contain something
	total := 0
	for index := 0; index < len(segments); index++ {
		s1 := segments[index]
		length := goutils.Abs(s1.MinX - s1.MaxX)
		total += length
		if VERBOSE {
			fmt.Printf("segment[%v], min=%v, max=%v, length=%v, total_so_far=%v\\n", s1, s1.MinX, s1.MaxX, length, total)
		}

	}

	return total
}
