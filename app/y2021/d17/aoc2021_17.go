package aoc2021

import (
	"fmt"
	"os"
	"sort"

	"github.com/simonski/aoc/utils"
	goutils "github.com/simonski/goutils"
)

/*
--- Day 17: Trick Shot ---
You finally decode the Elves' message. HI, the message says. You continue searching for the sleigh keys.

Ahead of you is what appears to be a large ocean trench. Could the keys have fallen into it? You'd better send a probe to investigate.

The probe launcher on your submarine can fire the probe with any integer velocity in the x (forward) and y (upward, or downward if negative) directions. For example, an initial x,y velocity like 0,10 would fire the probe straight up, while an initial velocity like 10,-1 would fire the probe forward at a slight downward angle.

The probe's x,y position starts at 0,0. Then, it will follow some trajectory by moving in steps. On each step, these changes occur in the following order:

The probe's x position increases by its x velocity.
The probe's y position increases by its y velocity.
Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if it is greater than 0, increases by 1 if it is less than 0, or does not change if it is already 0.
Due to gravity, the probe's y velocity decreases by 1.
For the probe to successfully make it into the trench, the probe must be on some trajectory that causes it to be within a target area after any step. The submarine computer has already calculated this target area (your puzzle input). For example:

target area: x=20..30, y=-10..-5
This target area means that you need to find initial x,y velocity values such that after any step, the probe's x position is at least 20 and at most 30, and the probe's y position is at least -10 and at most -5.

Given this target area, one initial velocity that causes the probe to be within the target area after any step is 7,2:

.............#....#............
.......#..............#........
...............................
S........................#.....
...............................
...............................
...........................#...
...............................
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTT#TT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
In this diagram, S is the probe's initial position, 0,0. The x coordinate increases to the right, and the y coordinate increases upward. In the bottom right, positions that are within the target area are shown as T. After each step (until the target area is reached), the position of the probe is marked with #. (The bottom-right # is both a position the probe reaches and a position in the target area.)

Another initial velocity that causes the probe to be within the target area after any step is 6,3:

...............#..#............
...........#........#..........
...............................
......#..............#.........
...............................
...............................
S....................#.........
...............................
...............................
...............................
.....................#.........
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................T#TTTTTTTTT
....................TTTTTTTTTTT
Another one is 9,0:

S........#.....................
.................#.............
...............................
........................#......
...............................
....................TTTTTTTTTTT
....................TTTTTTTTTT#
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
One initial velocity that doesn't cause the probe to be within the target area after any step is 17,-4:

S..............................................................
...............................................................
...............................................................
...............................................................
.................#.............................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT..#.............................
....................TTTTTTTTTTT................................
...............................................................
...............................................................
...............................................................
...............................................................
................................................#..............
...............................................................
...............................................................
...............................................................
...............................................................
...............................................................
...............................................................
..............................................................#
The probe appears to pass through the target area, but is never within it after any step. Instead, it continues down and to the right - only the first few steps are shown.

If you're going to fire a highly scientific probe out of a super cool probe launcher, you might as well do it with style. How high can you make the probe go while still reaching the target area?

In the above example, using an initial velocity of 6,9 is the best you can do, causing the probe to reach a maximum y position of 45. (Any higher initial y velocity causes the probe to overshoot the target area entirely.)

Find the initial velocity that causes the probe to reach the highest y position and still eventually be within the target area after any step. What is the highest y position it reaches on this trajectory?
*/

func (app *Application) Y2021D17_Summary() *utils.Summary {
	s := utils.NewSummary(2021, 17)
	s.Name = "Trick Shot"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

type Hit struct {
	VelocityAtStart int
	VelocityAtEnd   int
	Position        int
	Step            int
	MaxValue        int
	MaxValueStep    int
}

func (hit *Hit) CalculateMax() int {
	max := 0
	v := hit.VelocityAtStart
	p := 0
	for step := 0; step <= hit.Step; step++ {
		p += v
		v -= 1
		max = goutils.Max(max, p)
	}
	return max

}

func (hit *Hit) Debug() string {
	return fmt.Sprintf("Step %v, VelocityAtStart %v, VelocityAtEnd %v, Position %v, MaxValue: %v, MaxValueStep: %v", hit.Step, hit.VelocityAtStart, hit.VelocityAtEnd, hit.Position, hit.MaxValue, hit.MaxValueStep)

}
func NewHit(velocityAtStart int, velocityAtEnd, position int, step int) *Hit {
	return &Hit{VelocityAtStart: velocityAtStart, VelocityAtEnd: velocityAtEnd, Position: position, Step: step}
}

type HitDB struct {
	IndexByPosition   map[int][]*Hit
	IndexByStep       map[int][]*Hit
	ZeroVelocitySteps []int
}

func NewHitDB() *HitDB {
	db := &HitDB{IndexByPosition: make(map[int][]*Hit), IndexByStep: make(map[int][]*Hit), ZeroVelocitySteps: make([]int, 0)}
	return db
}

func (db *HitDB) Add(hit *Hit) {
	arr := db.IndexByPosition[hit.Position]
	if arr == nil {
		arr = make([]*Hit, 0)
	}
	arr = append(arr, hit)
	db.IndexByPosition[hit.Position] = arr

	arr = db.IndexByStep[hit.Step]
	if arr == nil {
		arr = make([]*Hit, 0)
	}
	arr = append(arr, hit)
	db.IndexByStep[hit.Step] = arr

	if hit.VelocityAtEnd == 0 {
		db.ZeroVelocitySteps = append(db.ZeroVelocitySteps, hit.Step)
	}
}

func (db *HitDB) Debug(title string) {
	keys := make([]int, 0, len(db.IndexByStep))
	for k := range db.IndexByStep {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	fmt.Println(keys)

	for _, key := range keys {
		arr := db.IndexByStep[key]
		// fmt.Printf("%v contains %v entries.\n", key, len(arr))
		for _, hit := range arr {
			fmt.Printf("%v %v\n", title, hit.Debug())
		}
	}

	// now retain only x-steps and y-steps that match

	// fmt.Println()

	// if len(my) == 0 {

	// 	fmt.Print("Nothing in y.\n")
	// } else {
	// 	for key, value := range my {
	// 		fmt.Printf("y: velocity %v: %v\n", key, value)
	// }
	// }

	// both maps look like [velocity]=[position=],

}

func Test_d17_part1(min_x int, max_x int, initial_xv int, max_y int, min_y int, initial_yv int, x_drag int, y_drag int) (*HitDB, *HitDB) {

	x_db := create_velo_map(min_x, max_x, x_drag, initial_xv, true)
	y_db := create_velo_map(min_y, max_y, y_drag, initial_yv, false)
	return x_db, y_db

}

// find the minimum and maximum x velocities that get is until the zone
func create_velo_map(min int, max int, drag int, initial int, clamp_to_zero bool) *HitDB {

	db := NewHitDB()
	for starting_velocity := initial; starting_velocity >= 0; starting_velocity -= 1 {
		position := 0
		step := 0
		v := starting_velocity
		max_value := 0
		max_value_step := 0
		for {
			step += 1
			position += v

			if clamp_to_zero {
				if v > 0 {
					v -= drag
				}
			} else {
				if !clamp_to_zero {
					v -= drag
				}
			}

			if position > max_value {
				max_value = goutils.Max(max_value, position)
				max_value_step = step
			}

			// if v == 0 {
			in_range := position >= min && position <= max
			if in_range {
				// velo_map[starting_velocity] = fmt.Sprintf("position=%v,current_velocity=%v,steps=%v", position, v, step)

				hit := NewHit(starting_velocity, v, position, step)
				hit.MaxValue = max_value
				hit.MaxValueStep = max_value_step

				db.Add(hit)
				// fmt.Printf("v=%v, position=%v, step=%v\n", v, position, step)
			}
			// }
			if clamp_to_zero {
				if position > max {
					break
				}
				if v == 0 {
					break
				}
			} else {
				if position < min {
					break
				}
			}

			// fmt.Printf("position=%v, velocity=%v, step=%v\n", position, v, step)
		}
	}
	return db

}

// x = >= 29 && <= 73

// x + (n*x_diff) >= 29 && <= 73

func (p *Point) Copy() *Point {
	p2 := &Point{}
	p2.X = p.X
	p2.Y = p.Y
	return p2
}

func (p *Point) StepD17(x_velocity int, y_velocity int, x_drag int, y_drag int) {
	p.X += x_velocity
	if p.X > 0 {
		p.X -= 1
	} else {
		p.X += 1
	}

	p.Y += y_velocity
	if p.Y > 0 {
		p.Y -= 1
	} else {
		p.Y += 1
	}

}

func winding_y(title string, initial_velocity int, min_target int, max_target int) []*H {
	p := 0
	steps := 0
	results := make([]*H, 0)
	v := initial_velocity
	// for v := initial_velocity; v >= min_velocity; v -= 1 {
	for {
		steps += 1
		p += v
		if p >= min_target && p <= max_target {
			fmt.Printf("%v: v=%v, final_v=%v, position=%v, steps=%v\n", title, initial_velocity, v, p, steps)
			h := &H{Initial_velocity: initial_velocity, Final_velocity: v, Position: p, Steps: steps}
			results = append(results, h)

		}
		v -= 1

		// if max_target > 0 {
		// 	if p > max_target {
		// 		break
		// 	}
		// }

		// if p < max_target < 0 {
		if p < min_target {
			break
		}
		// }
	}
	return results
}

func winding_x(title string, initial_velocity int, min_target int, max_target int, max_steps int) []*H {
	p := 0
	steps := 0
	results := make([]*H, 0)
	v := initial_velocity
	for step := 0; step <= max_steps; step++ {
		steps += 1
		p += v
		if p >= min_target && p <= max_target {
			fmt.Printf("%v: v=%v, final_v=%v, position=%v, steps=%v\n", title, initial_velocity, v, p, steps)
			h := &H{Initial_velocity: initial_velocity, Final_velocity: v, Position: p, Steps: steps}
			results = append(results, h)

		}
		v -= 1

		if steps > max_steps {
			break
		}

		// if max_target > 0 {
		// 	if p > max_target {
		// 		break
		// 	}
		// }

		// if max_target < 0 {
		// 	if p < min_target {
		// 		break
		// 	}
		// }
	}
	return results
}

type H struct {
	Initial_velocity int
	Final_velocity   int
	Position         int
	Steps            int
}

type HDB struct {
	X []*H
	Y []*H
}

func (db *HDB) AddX(h *H) {
	db.X = append(db.X, h)
	// fmt.Printf("X now contains %v\n", len(db.X))
}

func (db *HDB) AddY(h *H) {
	db.Y = append(db.Y, h)
	// fmt.Printf("Y now contains %v\n", len(db.Y))
}

func (db *HDB) GetYHits(x *H) []*H {
	results := make([]*H, 0)

	for _, y := range db.Y {
		if x.Final_velocity > 0 && y.Steps == x.Steps {
			// fmt.Printf("GetYHits(x=%v), y=%v, match exact.\n", x, y)
			results = append(results, y)
		} else if x.Final_velocity == 0 && y.Steps >= x.Steps {
			// fmt.Printf("GetYHits(x=%v), y=%v, match on 0.\n", x, y)
			results = append(results, y)
		}
	}
	return results
}

func (db *HDB) Foo() {
	m := make(map[string]bool)
	// fmt.Printf("There are %v entries in db.x to do a range on.\n", len(db.X))
	for _, x := range db.X {
		// fmt.Printf("xhit to check is %v\n", x)
		yhits := db.GetYHits(x)
		for _, yhit := range yhits {
			key := fmt.Sprintf("%v,%v", x.Initial_velocity, yhit.Initial_velocity)
			// fmt.Printf("%v\n", key)
			m[key] = true
		}
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k)
	}

	// for k, _ := range m {
	// 	if m[k] {
	// 		fmt.Printf("%v\n", k)
	// 	}
	// }

	// if m["23,-10"] {
	// 	fmt.Println("23,-10 exists!\n")
	// } else {
	// 	fmt.Println("23,-10 does not exist!\n")
	// }
	fmt.Printf("There are %v hits.\n", len(m))

}

func (db *HDB) Debug() string {
	line := ""
	fmt.Printf("There are %v entries in db.x to do a debug on.\n", len(db.X))
	for _, h := range db.X {
		line += fmt.Sprintf("X: v=%v, final_v=%v, position=%v, steps=%v\n", h.Initial_velocity, h.Final_velocity, h.Position, h.Steps)
	}
	for _, h := range db.X {
		line += fmt.Sprintf("Y: v=%v, final_v=%v, position=%v, steps=%v\n", h.Initial_velocity, h.Final_velocity, h.Position, h.Steps)
	}

	return line
}

func DemoD17_test() *HDB {

	db := &HDB{X: make([]*H, 0), Y: make([]*H, 0)}

	target_min := -10
	target_max := -5
	min_velocity := -10

	// calculate Y first
	for v := 45; v >= min_velocity; v -= 1 {
		results := winding_y("Y", v, target_min, target_max)
		for _, hit := range results {
			db.AddY(hit)
		}
	}

	max_steps := 0
	for _, yhit := range db.Y {
		max_steps = goutils.Max(max_steps, yhit.Steps)
	}
	fmt.Printf("Max steps for Y is %v\n", max_steps)

	// calculate X next *and* allow it to tick over to the maxumum number of steps (in case we count back into the target)
	min_velocity = -1000
	target_min = 20
	target_max = 30

	for v := 3000; v >= min_velocity; v -= 1 {
		hits := winding_x("X", v, target_min, target_max, max_steps)
		for _, hit := range hits {
			db.AddX(hit)
		}
	}

	fmt.Printf("foo\n")

	// so actually we can go "back" on x, also, so we should overshoot x by some

	return db

}

func demoD17_real() *HDB {

	db := &HDB{X: make([]*H, 0), Y: make([]*H, 0)}

	target_min := -248
	target_max := -194
	min_velocity := -248

	// calculate Y first
	for v := 30628; v >= min_velocity; v -= 1 {
		results := winding_y("Y", v, target_min, target_max)
		for _, hit := range results {
			db.AddY(hit)
		}
	}

	max_steps := 0
	for _, yhit := range db.Y {
		max_steps = goutils.Max(max_steps, yhit.Steps)
	}
	fmt.Printf("Max steps for Y is %v\n", max_steps)

	// calculate X next *and* allow it to tick over to the maxumum number of steps (in case we count back into the target)
	min_velocity = -1000
	target_min = 29
	target_max = 73

	for v := 3000; v >= min_velocity; v -= 1 {
		hits := winding_x("X", v, target_min, target_max, max_steps)
		for _, hit := range hits {
			db.AddX(hit)
		}
	}

	fmt.Printf("foo\n")

	// so actually we can go "back" on x, also, so we should overshoot x by some

	return db

}

// rename this to the year and day in question
func (app *Application) Y2021D17P1() {

	// db := demoD17_test()
	db := demoD17_real()
	db.Debug()

	db.Foo()
	os.Exit(1)

	// fmt.Printf("ArithSeq(248): %v\n", utils.ArithmeticProgression(1, 248))
	// fmt.Printf("ArithSeq(247): %v\n", utils.ArithmeticProgression(1, 247))
	// fmt.Printf("ArithSeq(249): %v\n", utils.ArithmeticProgression(1, 249))
	// fmt.Printf("ArithSeq(100): %v\n", utils.ArithmeticProgression(1, 100))
	// fmt.Printf("ArithSeq(50): %v\n", utils.ArithmeticProgression(1, 50))
	// fmt.Printf("ArithSeq(25): %v\n", utils.ArithmeticProgression(1, 25))
	// os.Exit(1)

	// highpoint, means by the time we get to 0 we are moving at highpoint speed

	// min_x := 29
	// max_x := 73
	// initial_xv := 300
	// max_y := -194
	// min_y := -248

	// // y-difference is 54

	// initial_yv := 5000
	// x_drag := 1
	// y_drag := 1

	// 3081  - 91000
	// 19701 - 45000
	min_x := 29
	max_x := 79
	max_y := -194
	min_y := -248
	initial_xv := 40
	initial_yv := 40
	x_drag := 1
	y_drag := 1

	use_test_data := false
	if use_test_data {
		min_x = 20
		max_x = 30
		max_y = -5
		min_y = -10
		initial_xv = 40
		initial_yv = 40
	} else {
		//		target area: x=29..73, y=-248..-194
		min_x = 29
		max_x = 79
		max_y = -194
		min_y = -248
		initial_xv = 1000
		initial_yv = 1000
	}

	x_db1, y_db1 := Test_d17_part1(min_x, max_x, initial_xv, max_y, min_y, initial_yv, x_drag, y_drag)
	x_db1.Debug("x")
	y_db1.Debug("y")

	// retain entries that hit on step
	x_db2 := NewHitDB()
	y_db2 := NewHitDB()
	for step := range x_db1.IndexByStep {
		if y_db1.IndexByStep[step] != nil {
			x_db2.IndexByStep[step] = x_db1.IndexByStep[step]
			y_db2.IndexByStep[step] = y_db1.IndexByStep[step]
		}
	}

	for x_step := range x_db1.ZeroVelocitySteps {
		for y_step := range y_db1.IndexByStep {
			if y_step >= x_step {
				y_db2.IndexByStep[y_step] = y_db1.IndexByStep[y_step]
			}
		}
	}

	fmt.Println()

	// point is the y steps have to intersect either exactly the x-steps
	// or they have to be an x-step which is a 0 rated x (almost certainly will be one of these)

	keys := make([]int, 0, len(y_db2.IndexByStep))
	for k := range y_db2.IndexByStep {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, step := range keys {
		entries := y_db2.IndexByStep[step]
		for _, entry := range entries {
			fmt.Printf("[%v] %v\n", step, entry.Debug())
		}
	}

	os.Exit(0)

	// now we have an x_db2 that contains only matching hits
	// and an y_db2 that contains matching and > y steps for x velocities that end at zero meaning we can slow
	// our x-velocity first but keep going higher
	// in this case we can then just ask any of the ys what their highest velocity was
	maximum_y := 0
	var max_entry *Hit
	for _, arr := range y_db2.IndexByStep {
		for _, entry := range arr {
			m := entry.CalculateMax()
			if m > max_y {
				maximum_y = m
				max_entry = entry

			}
			fmt.Printf("%v, maxY: %v\n", maximum_y, entry.Debug())
		}

	}

	fmt.Printf("Max y is %v, entry is %v\n", maximum_y, max_entry.Debug())
	// for any x-0 step, retain ANY y-enties that have a step >= this

}

// rename this to the year and day in question
func (app *Application) Y2021D17P2() {

	min_x := 29
	max_x := 79
	max_y := -194
	min_y := -248
	initial_xv := 40
	initial_yv := 40
	x_drag := 1
	y_drag := 1

	use_test_data := false
	if use_test_data {
		min_x = 20
		max_x = 30
		max_y = -5
		min_y = -10
		initial_xv = 40
		initial_yv = 40
	} else {
		//		target area: x=29..73, y=-248..-194
		min_x = 29
		max_x = 79
		max_y = -194
		min_y = -248
		initial_xv = 1000
		initial_yv = 1000
	}

	x_db1, y_db1 := Test_d17_part1(min_x, max_x, initial_xv, max_y, min_y, initial_yv, x_drag, y_drag)
	x_db1.Debug("x")
	y_db1.Debug("y")

	// retain entries that hit on step
	x_db2 := NewHitDB()
	y_db2 := NewHitDB()
	for step := range x_db1.IndexByStep {
		if y_db1.IndexByStep[step] != nil {
			x_db2.IndexByStep[step] = x_db1.IndexByStep[step]
			y_db2.IndexByStep[step] = y_db1.IndexByStep[step]
		}
	}

	for x_step := range x_db1.ZeroVelocitySteps {
		for y_step := range y_db1.IndexByStep {
			if y_step >= x_step {
				y_db2.IndexByStep[y_step] = y_db1.IndexByStep[y_step]
			}
		}
	}

	fmt.Println()

	// point is the y steps have to intersect either exactly the x-steps
	// or they have to be an x-step which is a 0 rated x (almost certainly will be one of these)

	keys := make([]int, 0, len(y_db2.IndexByStep))
	for k := range y_db2.IndexByStep {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, step := range keys {
		entries := y_db2.IndexByStep[step]
		for _, entry := range entries {
			fmt.Printf("[%v] %v\n", step, entry.Debug())
		}
	}

	os.Exit(0)

	// now we have an x_db2 that contains only matching hits
	// and an y_db2 that contains matching and > y steps for x velocities that end at zero meaning we can slow
	// our x-velocity first but keep going higher
	// in this case we can then just ask any of the ys what their highest velocity was
	maximum_y := 0
	var max_entry *Hit
	for _, arr := range y_db2.IndexByStep {
		for _, entry := range arr {
			m := entry.CalculateMax()
			if m > max_y {
				maximum_y = m
				max_entry = entry

			}
			fmt.Printf("%v, maxY: %v\n", maximum_y, entry.Debug())
		}

	}

	fmt.Printf("Max y is %v, entry is %v\n", maximum_y, max_entry.Debug())
	// for any x-0 step, retain ANY y-enties that have a step >= this

}

const D17_DATA = `23,-10
25,-9
27,-5
29,-6
22,-6
21,-7
9,0
27,-7
24,-5
25,-7
26,-6
25,-5
6,8
11,-2
20,-5
29,-10
6,3
28,-7
8,0
30,-6
29,-8
20,-10
6,7
6,4
6,1
14,-4
21,-6
26,-10
7,-1
7,7
8,-1
21,-9
6,2
20,-7
30,-10
14,-3
20,-8
13,-2
7,3
28,-8
29,-9
15,-3
22,-5
26,-8
25,-8
25,-6
15,-4
9,-2
15,-2
12,-2
28,-9
12,-3
24,-6
23,-7
25,-10
7,8
11,-3
26,-7
7,1
23,-9
6,0
22,-10
27,-6
8,1
22,-8
13,-4
7,6
28,-6
11,-4
12,-4
26,-9
7,4
24,-10
23,-8
30,-8
7,0
9,-1
10,-1
26,-5
22,-9
6,5
7,5
23,-6
28,-10
10,-2
11,-1
20,-9
14,-2
29,-7
13,-3
23,-5
24,-8
27,-9
30,-7
28,-5
21,-10
7,9
6,6
21,-5
27,-10
7,2
30,-9
21,-8
22,-7
24,-9
20,-6
6,9
29,-5
8,-2
27,-8
30,-5
24,-7`
