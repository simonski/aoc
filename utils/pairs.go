package utils

import "fmt"

// returns combinations of pairs
// input [ 0 1 2 3 4 5 6 ]
// output
// [0 1] [0 2] [0 3] [0 4] [0 5] [0 6]
// [1 2] [1 3] [1 4] [1 5] [1 6]
// [2 3] [2 4] [2 5] [2 6]
// [3 4] [3 5] [3 6]
// [4 5] [4 6]
// [5 6]

func pairs(values []int) []string {
	pairs := make([]string, 0)
	for index, value_a := range values {
		for _, value_b := range values[index:] {
			pair := fmt.Sprintf("%v.%v", value_a, value_b)
			pairs = append(pairs, pair)
		}
	}
	return pairs
}
