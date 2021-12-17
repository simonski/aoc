package aoc2021

func Pairs(value string) []string {
	pairs := make([]string, 0)
	for index := 0; index+1 < len(value); index++ {
		pair := value[index : index+2]
		pairs = append(pairs, pair)
	}
	return pairs
}
