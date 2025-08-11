package aoc2021

import (
	"fmt"
	"strings"

	utils "github.com/simonski/aoc/utils"
)

/*
--- Day 12: Passage Pathing ---
With your submarine's subterranean subsystems subsisting suboptimally, the only way you're getting out of this cave anytime soon is by finding a path yourself. Not just a path - the only way to know if you've found the best path is to find all of them.

Fortunately, the sensors are still mostly working, and so you build a rough map of the remaining caves (your puzzle input). For example:

start-A
start-b
A-c
A-b
b-d
A-end
b-end
This is a list of how all of the caves are connected. You start in the cave named start, and your destination is the cave named end. An entry like b-d means that cave b is connected to cave d - that is, you can move between them.

So, the above cave system looks roughly like this:

    start
    /   \
c--A-----b--d
    \   /
     end
Your goal is to find the number of distinct paths that start at start, end at end, and don't visit small caves more than once. There are two types of caves: big caves (written in uppercase, like A) and small caves (written in lowercase, like b). It would be a waste of time to visit any small cave more than once, but big caves are large enough that it might be worth visiting them multiple times. So, all paths you find should visit small caves at most once, and can visit big caves any number of times.

Given these rules, there are 10 paths through this example cave system:

start,A,b,A,c,A,end
start,A,b,A,end
start,A,b,end
start,A,c,A,b,A,end
start,A,c,A,b,end
start,A,c,A,end
start,A,end
start,b,A,c,A,end
start,b,A,end
start,b,end

(Each line in the above list corresponds to a single path; the caves visited by that path are listed in the order they are visited and separated by commas.)

Note that in this cave system, cave d is never visited by any path: to do so, cave b would need to be visited twice (once on the way to cave d and a second time when returning from cave d), and since cave b is small, this is not allowed.

Here is a slightly larger example:

dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
The 19 paths through it are as follows:

start,HN,dc,HN,end
start,HN,dc,HN,kj,HN,end
start,HN,dc,end
start,HN,dc,kj,HN,end
start,HN,end
start,HN,kj,HN,dc,HN,end
start,HN,kj,HN,dc,end
start,HN,kj,HN,end
start,HN,kj,dc,HN,end
start,HN,kj,dc,end
start,dc,HN,end
start,dc,HN,kj,HN,end
start,dc,end
start,dc,kj,HN,end
start,kj,HN,dc,HN,end
start,kj,HN,dc,end
start,kj,HN,end
start,kj,dc,HN,end
start,kj,dc,end
Finally, this even larger example has 226 paths through it:

fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
How many paths through this cave system are there that visit small caves at most once?

*/

func (app *Application) Y2021D12_Summary() *utils.Summary {
	s := utils.NewSummary(2021, 12)
	s.Name = "Passage Pathing"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

type Cave struct {
	Name       string
	Caves      []*Cave
	VisitCount int
}

type Path struct {
	Caves []*Cave
}

func (p *Path) Debug() string {
	line := ""
	for index, c := range p.Caves {
		line += c.Name
		if index+1 < len(p.Caves) {
			line += ","
		}
	}
	return line
}

func (p *Path) CanVisit(maxVisits int, c *Cave) bool {
	if c.IsStart() {
		return false
	} else if c.IsEnd() {
		return true
	} else if c.IsLarge() {
		return true
	} else {
		// it is small - have we visited it before
		// return c.VisitCount == 0
		// ok at most 1 small cave can appear twice
		counter := make(map[string]int)
		doubleCount := 0
		for _, candidate := range p.Caves {
			if candidate.IsSmall() {
				count := counter[candidate.Name]
				count++
				counter[candidate.Name] = count
				if count >= maxVisits {
					doubleCount++
					if doubleCount > 1 {
						return false
					}
				}
			}
		}
		return true
	}
}

func (p *Path) Last() *Cave {
	return p.Caves[len(p.Caves)-1]
}

func (p *Path) Size() int {
	return len(p.Caves) - 1
}

func (p *Path) IsEnd() bool {
	return p.Last().IsEnd()
}

func (p *Path) Add(c *Cave) {
	p.Caves = append(p.Caves, c)
}

func (p *Path) Copy() *Path {
	p2 := NewPath()
	for _, cave := range p.Caves {
		p2.Add(cave)
	}
	return p2
}

func NewPath() *Path {
	return &Path{Caves: make([]*Cave, 0)}
}

func (c *Cave) IsStart() bool {
	return c.Name == "start"
}

func (c *Cave) IsEnd() bool {
	return c.Name == "end"
}

func (c *Cave) Add(destination *Cave) {
	c.Caves = append(c.Caves, destination)
	destination.Caves = append(destination.Caves, c)
}

func (c *Cave) IsLarge() bool {
	return strings.ToUpper(c.Name) == c.Name
}

func (c *Cave) IsSmall() bool {
	return !c.IsLarge()
}

type CaveSystem struct {
	Caves map[string]*Cave
	All   []*Cave
	Start *Cave
	End   *Cave
}

func (cs *CaveSystem) Size() int {
	return len(cs.Caves)
}

func (cs *CaveSystem) PathFind(maxVisits int) []*Path { // returns an array of paths
	// for each child, branch
	// as we go to a child, mark it visited
	paths := make([]*Path, 0)
	start := cs.Start
	for _, cave := range cs.Start.Caves {
		// cs.Reset()
		path := NewPath()
		paths = append(paths, path)
		path.Add(start)
		path.Add(cave)
		paths = cs.WalkPath(0, maxVisits, path, paths)
	}
	return paths
}

func DebugPath(path []*Cave) string {
	line := path[0].Name
	for index := 1; index < len(path); index++ {
		cave := path[index]
		line = fmt.Sprintf("%v,%v", line, cave.Name)
	}
	return line
}

/*
Your goal is to find the number of distinct paths that start at start, end at end, and don't visit small caves more than once. There are two types of caves: big caves (written in uppercase, like A) and small caves (written in lowercase, like b). It would be a waste of time to visit any small cave more than once, but big caves are large enough that it might be worth visiting them multiple times. So, all paths you find should visit small caves at most once, and can visit big caves any number of times.
*/

func (cs *CaveSystem) WalkPath(depth int, maxVisits int, path *Path, paths []*Path) []*Path {
	cave := path.Last()
	// fmt.Printf("WalkPath(%v)\n", path.Debug())
	if cave.IsEnd() {
		// fmt.Println(path.Debug())
		return paths
	} else if cave.IsSmall() && !path.CanVisit(maxVisits, cave) {
		return paths // we quit that path as we can go no further
	}
	// fmt.Printf("WalkPath(cave=%v, depth=%v, path=%v)\n", cave.Name, depth, DebugPath(path))

	cave.VisitCount++

	for _, c := range cave.Caves {
		if path.CanVisit(maxVisits, c) {
			new_path := path.Copy()
			new_path.Add(c)
			paths = append(paths, new_path)
			d := depth + 1
			paths = cs.WalkPath(d, maxVisits, new_path, paths)
		}
	}
	return paths

}

func NewCaveSystem(data string) *CaveSystem {
	lines := strings.Split(data, "\n")
	caves := make(map[string]*Cave)
	var start *Cave
	var end *Cave
	all := make([]*Cave, 0)

	for _, line := range lines {
		splits := strings.Split(line, "-")
		c1name := splits[0]
		c2name := splits[1]
		cave1 := caves[c1name]
		cave2 := caves[c2name]
		if cave1 == nil {
			cave1 = &Cave{Name: c1name, Caves: make([]*Cave, 0)}
			all = append(all, cave1)
		}
		if cave2 == nil {
			cave2 = &Cave{Name: c2name, Caves: make([]*Cave, 0)}
			all = append(all, cave2)
		}
		caves[c1name] = cave1
		caves[c2name] = cave2
		cave1.Add(cave2)

		if cave1.IsStart() {
			start = cave1
		}

		if cave1.IsEnd() {
			end = cave1
		}

		if cave2.IsStart() {
			start = cave2
		}

		if cave2.IsEnd() {
			end = cave2
		}
	}
	cs := &CaveSystem{Caves: caves, Start: start, End: end, All: all}
	return cs
}

// rename this to the year and day in question
func (app *Application) Y2021D12P1() {
	day12Part1("test1", 1, DAY_2021_12_TEST_DATA_7)
	day12Part1("test2", 1, DAY_2021_12_TEST_DATA_19)
	day12Part1("testX", 1, DAY_2021_12_TEST_DATA_X)
	day12Part1("real", 1, DAY_2021_12_DATA)
}

func (app *Application) Y2021D12P2() {
	day12Part1("test1", 2, DAY_2021_12_TEST_DATA_7)
	day12Part1("test2", 2, DAY_2021_12_TEST_DATA_19)
	day12Part1("testX", 2, DAY_2021_12_TEST_DATA_X)
	day12Part1("real", 2, DAY_2021_12_DATA)
}

func day12Part1(title string, maxVisits int, data string) {
	cs := NewCaveSystem(data)
	paths := cs.PathFind(maxVisits)

	// cs2 := NewCaveSystem(DAY_2021_12_TEST_DATA_19)
	// paths2 := cs.PathFind()

	count := 0
	for _, path := range paths {
		if path.IsEnd() {
			count++
			fmt.Println(path.Debug())
		}
	}
	fmt.Printf("%v : There are %v paths.\n", title, count)
}

// // rename this to the year and day in question
// func (app *Application) Y2021D12P2() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
// func (app *Application) Y2021D12() {
// 	app.Y2021D12P1()
// 	app.Y2021D12P2()
// }
