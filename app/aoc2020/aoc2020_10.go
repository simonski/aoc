package aoc2020

/*

https://adventofcode.com/2020/day/10
--- Day 10: Adapter Array ---

Patched into the aircraft's data port, you discover weather forecasts of a massive tropical storm. Before you can figure out whether it will impact your vacation plans, however, your device suddenly turns off!

Its battery is dead.

You'll need to plug it in. There's only one problem: the charging outlet near your seat produces the wrong number of jolts. Always prepared, you make a list of all of the joltage adapters in your bag.

Each of your joltage adapters is rated for a specific output joltage (your puzzle input). Any given adapter can take an input 1, 2, or 3 jolts lower than its rating and still produce its rated output joltage.

In addition, your device has a built-in joltage adapter rated for 3 jolts higher than the highest-rated adapter in your bag. (If your adapter list were 3, 9, and 6, your device's built-in adapter would be rated for 12 jolts.)

Treat the charging outlet near your seat as having an effective joltage rating of 0.

Since you have some time to kill, you might as well test all of your adapters. Wouldn't want to get to your resort and realize you can't even charge your device!

If you use every adapter in your bag at once, what is the distribution of joltage differences between the charging outlet, the adapters, and your device?

For example, suppose that in your bag, you have adapters with the following joltage ratings:

16
10
15
5
1
11
7
19
6
12
4
With these adapters, your device's built-in joltage adapter would be rated for 19 + 3 = 22 jolts, 3 higher than the highest-rated adapter.

Because adapters can only connect to a source 1-3 jolts lower than its rating, in order to use every adapter, you'd need to choose them like this:

The charging outlet has an effective rating of 0 jolts, so the only adapters that could connect to it directly would need to have a joltage rating of 1, 2, or 3 jolts. Of these, only one you have is an adapter rated 1 jolt (difference of 1).
From your 1-jolt rated adapter, the only choice is your 4-jolt rated adapter (difference of 3).
From the 4-jolt rated adapter, the adapters rated 5, 6, or 7 are valid choices. However, in order to not skip any adapters, you have to pick the adapter rated 5 jolts (difference of 1).
Similarly, the next choices would need to be the adapter rated 6 and then the adapter rated 7 (with difference of 1 and 1).
The only adapter that works with the 7-jolt rated adapter is the one rated 10 jolts (difference of 3).
From 10, the choices are 11 or 12; choose 11 (difference of 1) and then 12 (difference of 1).
After 12, only valid adapter has a rating of 15 (difference of 3), then 16 (difference of 1), then 19 (difference of 3).
Finally, your device's built-in adapter is always 3 higher than the highest adapter, so its rating is 22 jolts (always a difference of 3).
In this example, when using every adapter, there are 7 differences of 1 jolt and 5 differences of 3 jolts.

Here is a larger example:

28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
In this larger example, in a chain that uses all of the adapters, there are 22 differences of 1 jolt and 10 differences of 3 jolts.

Find a chain that uses all of your adapters to connect the charging outlet to your device's built-in adapter and count the joltage differences between the charging outlet, the adapters, and your device. What is the number of 1-jolt differences multiplied by the number of 3-jolt differences?

*/
import (
	"fmt"
	"sort"

	"github.com/simonski/aoc/utils"
	goutils "github.com/simonski/goutils"
)

func (app *Application) Y2020D10_Summary() *utils.Summary {
	s := utils.NewSummary(2020, 10)
	s.Name = "Adapter Array"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

// AOC_2020_10 is the entrypoint
func (app *Application) Y2020D10P1() {
	// AOC_2020_10_part1_attempt1_from_cli(cli)
	// AOC_2020_10_part2_attempt1(cli)
}

func (app *Application) Y2020D10P2() {
	AOC_2020_10_part2_attempt2(app)
}

func AOC_2020_10_part1_attempt1_from_cli(app *Application) {
	cli := app.CLI
	filename := cli.GetFileExistsOrDie("-input")
	ints := goutils.Load_file_to_ints(filename)
	results := AOC_2020_10_part1_attempt1(ints)
	fmt.Printf("%v\n", results)

	value1 := results[1]
	value3 := results[3]

	fmt.Printf("%v * %v = %v\n", value1, value3, value1*value3)

}

func AOC_2020_10_part1_attempt1(ints []int) map[int]int {
	counts := make(map[int]int)
	sort.Ints(ints)
	current := 0
	for index1, value1 := range ints {
		diff := value1 - current
		current = value1
		_, exists := counts[diff]
		if exists {
			counts[diff]++
		} else {
			counts[diff] = 1
		}
		if index1+1 == len(ints) {
			break
		}
	}
	counts[3]++
	return counts
}

// type InfoMap struct {
// 	path      []int
// 	TickCount int
// }

// func (m *InfoMap) Tick() {
// 	m.TickCount++
// }
// func (m *InfoMap) Debug() {
// 	fmt.Printf("Info(tick=%v)\n", m.TickCount)
// 	fmt.Printf("Path: ")
// 	for _, value := range m.path {
// 		fmt.Printf("%v -> ", value)
// 	}
// 	fmt.Printf("\n")
// 	// fmt.Printf("%v\n", m.path)
// }
// func (m *InfoMap) Push(value int) {
// 	m.path = append(m.path, value)
// }

// func (m *InfoMap) Pop() int {
// 	value := m.path[len(m.path)-1]
// 	m.path = m.path[0 : len(m.path)-1]
// 	return value
// }

// func AOC_2020_10_part2_attempt1(cli *goutils.CLI) {
// 	filename := cli.GetFileExistsOrDie("-input")
// 	data := load_file_to_ints(filename)
// 	tolerance := cli.GetIntOrDie("-tolerance")

// 	// walk down the tree
// 	sort.Ints(data)
// 	info := &InfoMap{}
// 	total := walk(0, tolerance, data, info)
// 	fmt.Printf("There are %v paths\n", total)
// 	fmt.Printf("%v\n", data)
// }

func AOC_2020_10_part2_attempt2(app *Application) {
	cli := app.CLI
	filename := cli.GetFileExistsOrDie("-input")
	data := goutils.Load_file_to_ints(filename)
	fmt.Printf("data is   %v\n", data)
	sort.Ints(data)
	fmt.Printf("sorted is %v\n", data)

	// FROM JT as my head exploded:
	// # The number of ways to get to current joltage = ways to get to 1 jolt less than this joltage
	// #                                                + ways to get to 2 jolts less than this joltage
	// #                                                + ways to get to 3 jolts less than this joltage
	// ways[a] = ways.get(a - 1, 0) + ways.get(a - 2, 0) + ways.get(a - 3, 0)

	m := goutils.NewIntMap()
	m.Put(0, 1)
	max := 0
	for _, value := range data {
		total := m.Get(value-1, 0) + m.Get(value-2, 0) + m.Get(value-3, 0)
		m.Put(value, total)
		max = goutils.Max(max, total)
	}

	fmt.Printf("%v\n", m)
	fmt.Printf("Max is %v\n", max)

}

// func waysToGet(index int, tolerance int, data []int) []int {
// 	results := make([]int, 0)
// 	value := data[index]

// 	startPos := index - tolerance
// 	if startPos < 0 {
// 		startPos = 0
// 	}
// 	for position := startPos; position < index; position++ {
// 		candidate := data[position]
// 		if candidate+tolerance >= value {
// 			fmt.Printf("    [%v] candidate %v is in range.\n", position, candidate)
// 			results = append(results, candidate+tolerance)
// 		} else {
// 			fmt.Printf("    [%v] candidate %v is not in range.\n", position, candidate)

// 		}
// 	}
// 	return results
// }

// func walk(initial_index int, tolerance int, data []int, info *InfoMap) int {
// 	// then we are at the end and we don't need to walk the tree
// 	// else we aren't and we can attempt to walk it
// 	initial_value := data[initial_index]
// 	fmt.Printf("walk(index=%v value=%v)\n", initial_index, initial_value)
// 	info.Tick()
// 	if initial_index == len(data)-1 {
// 		// have now walked the whole thing
// 		fmt.Printf("\nwalk: we have now walked the whole length\n")
// 		return 1
// 	}

// 	if info.TickCount > 12 {
// 		os.Exit(0)
// 	}

// 	info.Push(initial_value)
// 	info.Debug()
// 	fmt.Printf("walk: initial index %v, initial value %v, tolerance %v, data %v\n", initial_index, initial_value, tolerance, data)
// 	total := 0
// 	for index := initial_index + 1; index < len(data); index++ {
// 		value := data[index]
// 		if value > initial_value+tolerance {
// 			// then we have exceeded the acceptable value and should break
// 			fmt.Printf("walk() walk children index %v, value %v\n", index, value)
// 			info.Pop()

// 		} else {
// 			// otherwise we are "inside" our tolerance and this is a child
// 			total += walk(index, tolerance, data, info)
// 		}
// 	}
// 	return total

// }

/*
--- Part Two ---
To completely determine whether you have enough adapters, you'll need to figure out how many different ways they can be arranged. Every arrangement needs to connect the charging outlet to your device. The previous rules about when adapters can successfully connect still apply.

The first example above (the one that starts with 16, 10, 15) supports the following arrangements:

(0), 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, (22)
(0), 1, 4, 5, 6, 7, 10, 12, 15, 16, 19, (22)
(0), 1, 4, 5, 7, 10, 11, 12, 15, 16, 19, (22)
(0), 1, 4, 5, 7, 10, 12, 15, 16, 19, (22)
(0), 1, 4, 6, 7, 10, 11, 12, 15, 16, 19, (22)
(0), 1, 4, 6, 7, 10, 12, 15, 16, 19, (22)
(0), 1, 4, 7, 10, 11, 12, 15, 16, 19, (22)
(0), 1, 4, 7, 10, 12, 15, 16, 19, (22)
(The charging outlet and your device's built-in adapter are shown in parentheses.) Given the adapters from the first example, the total number of arrangements that connect the charging outlet to your device is 8.

The second example above (the one that starts with 28, 33, 18) has many arrangements. Here are a few:

(0), 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31,
32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49, (52)

(0), 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31,
32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 49, (52)

(0), 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31,
32, 33, 34, 35, 38, 39, 42, 45, 46, 48, 49, (52)

(0), 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31,
32, 33, 34, 35, 38, 39, 42, 45, 46, 49, (52)

(0), 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31,
32, 33, 34, 35, 38, 39, 42, 45, 47, 48, 49, (52)

(0), 3, 4, 7, 10, 11, 14, 17, 20, 23, 25, 28, 31, 34, 35, 38, 39, 42, 45,
46, 48, 49, (52)

(0), 3, 4, 7, 10, 11, 14, 17, 20, 23, 25, 28, 31, 34, 35, 38, 39, 42, 45,
46, 49, (52)

(0), 3, 4, 7, 10, 11, 14, 17, 20, 23, 25, 28, 31, 34, 35, 38, 39, 42, 45,
47, 48, 49, (52)

(0), 3, 4, 7, 10, 11, 14, 17, 20, 23, 25, 28, 31, 34, 35, 38, 39, 42, 45,
47, 49, (52)

(0), 3, 4, 7, 10, 11, 14, 17, 20, 23, 25, 28, 31, 34, 35, 38, 39, 42, 45,
48, 49, (52)
In total, this set of adapters can connect the charging outlet to your device in 19208 distinct arrangements.

You glance back down at your bag and try to remember why you brought so many adapters; there must be more than a trillion valid ways to arrange them! Surely, there must be an efficient way to count the arrangements.

What is the total number of distinct ways you can arrange the adapters to connect the charging outlet to your device?


*/
