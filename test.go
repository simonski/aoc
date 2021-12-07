package main

// import (
// 	"fmt"
// )

// func build_prefix(depth int) string {
// 	result := ""
// 	for index := 0; index < depth; index++ {
// 		result += " "
// 	}
// 	return result
// }

// func main() {

// 	// 18 = 26
// 	// 80 = 5934
// 	// 256 = ?

// 	// 26 for 18 rounds
// 	// smallfish := []int{3, 4, 3, 1, 2}
// 	// smallfish := []int{1}
// 	// smallfish := []int{4}

// 	bigfish := []int{1, 1, 3, 5, 3, 1, 1, 4, 1, 1, 5, 2, 4, 3, 1, 1, 3, 1, 1, 5, 5, 1, 3, 2, 5, 4, 1, 1, 5, 1, 4, 2, 1, 4, 2, 1, 4, 4, 1, 5, 1, 4, 4, 1, 1, 5, 1, 5, 1, 5, 1, 1, 1, 5, 1, 2, 5, 1, 1, 3, 2, 2, 2, 1, 4, 1, 1, 2, 4, 1, 3, 1, 2, 1, 3, 5, 2, 3, 5, 1, 1, 4, 3, 3, 5, 1, 5, 3, 1, 2, 3, 4, 1, 1, 5, 4, 1, 3, 4, 4, 1, 2, 4, 4, 1, 1, 3, 5, 3, 1, 2, 2, 5, 1, 4, 1, 3, 3, 3, 3, 1, 1, 2, 1, 5, 3, 4, 5, 1, 5, 2, 5, 3, 2, 1, 4, 2, 1, 1, 1, 4, 1, 2, 1, 2, 2, 4, 5, 5, 5, 4, 1, 4, 1, 4, 2, 3, 2, 3, 1, 1, 2, 3, 1, 1, 1, 5, 2, 2, 5, 3, 1, 4, 1, 2, 1, 1, 5, 3, 1, 4, 5, 1, 4, 2, 1, 1, 5, 1, 5, 4, 1, 5, 5, 2, 3, 1, 3, 5, 1, 1, 1, 1, 3, 1, 1, 4, 1, 5, 2, 1, 1, 3, 5, 1, 1, 4, 2, 1, 2, 5, 2, 5, 1, 1, 1, 2, 3, 5, 5, 1, 4, 3, 2, 2, 3, 2, 1, 1, 4, 1, 3, 5, 2, 3, 1, 1, 5, 1, 3, 5, 1, 1, 5, 5, 3, 1, 3, 3, 1, 2, 3, 1, 5, 1, 3, 2, 1, 3, 1, 1, 2, 3, 5, 3, 5, 5, 4, 3, 1, 5, 1, 1, 2, 3, 2, 2, 1, 1, 2, 1, 4, 1, 2, 3, 3, 3, 1, 3, 5}

// 	// brute(80, smallfish)
// 	// brute(bigfish)

// 	// algo(18, smallfish)
// 	// algo(80, smallfish)
// 	// algo(256, smallfish)
// 	algo(256, bigfish)
// 	// algo(80, bigfish)

// }

// func algo(days int, data []int) {
// 	// total := len(data)
// 	total := 0
// 	depth := 0

// 	cache := make(map[int]int)

// 	for _, value := range data {
// 		count := count_children(cache, depth, days-(value+1)) // 1 is this fish itself
// 		// fmt.Printf("algo[%v], depth=%v, days=%v, count=%v\n", value, depth, days, count)
// 		total += count
// 	}
// 	// total
// 	total += len(data)
// 	fmt.Printf("algo(days=%v, data=%v), total=%v\n", days, data, total)
// }

// // this day days is the creation day
// func count_children(cache map[int]int, depth int, days int) int {
// 	if cache[days] != 0 {
// 		return cache[days]
// 	}
// 	if days < 0 {
// 		// fmt.Printf("count_children: depth=%v, days=%v, returning 0\n", depth, days)
// 		return 0
// 	}
// 	if days < 7 {
// 		// we KNOW we can't create a child, so we can safely return only this fish
// 		// fmt.Printf("count_children: depth=%v, days=%v, returning 1\n", depth, days)
// 		return 1
// 	}

// 	// so today is a count of 1 because it has created 1 fish
// 	// now we look at where we are and work out how many can we create from here

// 	// direct_spawns := (days / 7) // the +1 is "this" fish
// 	// c := direct_spawns
// 	// fmt.Printf("count_children: depth=%v, days=%v, direct_spawns=%v\n", depth, days, direct_spawns)

// 	// spawning_day := days
// 	total := 0
// 	for test_day := days; test_day >= 0; test_day -= 7 {
// 		// we spawn on this day - can we use this day + 9 as a spawning day?
// 		total += 1
// 		if test_day >= 0 {
// 			total += count_children(cache, depth+1, test_day-9)
// 		}
// 	}

// 	fmt.Printf("[depth=%v, days=%v, returning %v]\n", depth, days, total)
// 	cache[days] = total
// 	return total

// }

// func brute(days int, data []int) {
// 	fish := data
// 	fmt.Printf("hi, fish=%v\n", fish)
// 	for day := 0; day < days; day++ {
// 		new_fish := 0
// 		for index := 0; index < len(fish); index++ {
// 			f := fish[index]
// 			f -= 1
// 			fish[index] = f
// 			if f < 0 {
// 				new_fish += 1
// 				fish[index] = 6
// 			}

// 		}
// 		for index := 0; index < new_fish; index++ {
// 			fish = append(fish, 8)
// 		}
// 		fmt.Printf("Day[%v] = %v\n", day, len(fish))
// 	}
// }
