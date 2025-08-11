package d12

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/simonski/cli"
)

var COUNTER int

func descend(context *Context, position int, current_line string, original_line string, rules []int) int {

	context.depth += 1
	letter := original_line[position : position+1]
	// spacesX := strings.Repeat(" ", len(original_line)-len(current_line))
	if letter == "?" {
		// fmt.Printf("descend   '%v%v', try '.' then '#'\n", current_line, spacesX)
		candidateL := current_line + "."
		candidateR := current_line + "#"
		resultL := evaluate(context, ".", position, candidateL, original_line, rules)
		// if context.cacheEnabled {
		// 	context.putCache(candidateL, position, resultL, nil)
		// }
		// fmt.Printf("L done, position=%v, result=%v\n", position, resultL)
		resultR := evaluate(context, "#", position, candidateR, original_line, rules)
		// if context.cacheEnabled {
		// 	context.putCache(candidateR, position, resultR, nil)
		// }
		// fmt.Printf("R done, position=%v, result=%v\n", position, resultR)
		context.depth -= 1
		return resultL + resultR
	} else {
		// fmt.Printf("descend '%v', use '%v'\n", current_line, spacesX)
		candidate := current_line + letter
		result := evaluate(context, letter, position, candidate, original_line, rules)

		// if context.cacheEnabled {
		// 	context.putCache(candidate, position, result, nil)
		// }

		context.depth -= 1
		return result
	}
}

func evaluate(context *Context, letter string, position int, current_line string, original_line string, rules []int) int {

	position += 1

	// what we need to cache is the work AHEAD, not hte work we've done.  This way you build a library
	// of work as we "know" the other way round
	// it's the work on the right somehow

	if context.cacheEnabled {
		value, exists := context.checkCache(current_line, position+1, nil)
		if exists {
			return value
		}
	}

	spacesX := strings.Repeat(" ", len(original_line)-len(current_line))
	if position > len(original_line) {
		fmt.Printf("R1 evaluate: '%v%v' position %v (len=%v) RETURN 0\n", current_line, spacesX, position, len(original_line))
		if context.cacheEnabled {
			context.putCache(current_line, position, 0, nil)
		}
		return 0
	}

	actuals := buildBlocks(current_line, rules)
	evalCode, _, evalInvalid := evaluateBlocks(position, current_line, original_line, actuals, rules)
	// spacesX := strings.Repeat(" ", len(original_line)-len(current_line))

	if evalInvalid {
		if context.cacheEnabled {
			context.putCache(current_line, position, 0, nil)
		}
		return 0
	}

	if len(current_line) == len(original_line) {
		if evalCode == EVAL_RULES_ALL_MATCH {
			if context.cacheEnabled {
				context.putCache(current_line, position, 1, nil)
			}
			return 1
		} else {
			if context.cacheEnabled {
				context.putCache(current_line, position, 0, nil)
			}
			return 0
		}
	}

	return descend(context, position, current_line, original_line, rules)

}

// buildActuals makes an []int of the "actual" contiguous blocks on a given line
// the "last" block might not pass a rule
func buildBlocks(current_line string, rules []int) []int {
	count := 0
	results := make([]int, 0)
	lastEntry := ""
	for _, e := range current_line {
		entry := string(e)
		if entry == "." {
			// then close off the last contiguous damaged block (if there is one)
			if count > 0 {
				results = append(results, count)
				count = 0
			}
		} else if entry == "#" {
			// then increase the block size
			count++
		}
		lastEntry = entry
	}
	if lastEntry == "#" {
		// then store this as the last block
		// unless we are still processing (the length check), drop it as it's not - yet a contiguous block
		results = append(results, count)
	}

	return results
}

const EVAL_RULES_ALL_MATCH = 0
const EVAL_TOO_MANY_BLOCKS = 1
const EVAL_BLOCK_TOO_BIG = 2
const EVAL_BLOCK_TOO_SMALL = 3
const EVAL_IN_PROGRESS = 4

func evaluateBlocks(position int, current_line string, original_line string, results []int, rules []int) (int, string, bool) {

	// invalid := false

	if len(results) > len(rules) {
		// too many blocks
		return EVAL_TOO_MANY_BLOCKS, "EVAL_TOO_MANY_BLOCKS", true
	}

	for index := range results {
		if results[index] > rules[index] {
			// definately invalid - block too large
			return EVAL_BLOCK_TOO_BIG, "EVAL_BLOCK_TOO_BIG  ", true
		}

		if results[index] < rules[index] {
			// it *might* be invalid - if there are more blocks after this, then it is invalid
			if index+1 < len(results) {
				return EVAL_BLOCK_TOO_SMALL, "EVAL_BLOCK_TOO_SMALL", true
			}
			// check the lastLetter - if it is a "." then the block is closed and it is too small
			l := len(current_line)
			lastLetter := current_line[l-1 : l]
			if lastLetter == "." {
				return EVAL_BLOCK_TOO_SMALL, "EVAL_BLOCK_TOO_SMALL", true
			}
		}
	}

	// so far we know
	// 1. there are not too many blocks
	// 2. no individual blocks are larger than their rule
	// 3. no individual blocks are smaller than their rule (except maybe the last one)
	if reflect.DeepEqual(results, rules) {
		// then all rules pass perfectly
		return EVAL_RULES_ALL_MATCH, "EVAL_RULES_ALL_MATCH", false
	} else {
		// then it's not wrong but it's not perfect
		return EVAL_IN_PROGRESS, "EVAL_IN_PROGRESS    ", false
	}

	// else it's not definately invalid, so we say, so far, it's valid.

	// isLast := index == len(results)-1

	// if !isLast && results[index] != rules[index] {
	// 	// then we know this is invalid as they don't match AND there is another rule
	// 	invalid = true
	// 	break
	// }

	// if isLast && results[index] != rules[index] && lastEntry == "." {
	// 	// then we know this is invalid as they don't match AND there is another rule
	// 	invalid = true
	// 	break
	// }

}

// return 1, invalid
// }

type Context struct {
	cache        map[string]int
	hits         int
	misses       int
	cli          *cli.CLI
	verbose      bool
	cacheEnabled bool
	depth        int
}

func NewContext(cli *cli.CLI) *Context {
	c := Context{}
	c.cache = make(map[string]int)
	c.cli = cli
	c.hits = 0
	c.misses = 0
	return &c
}

func (c *Context) Debug() string {
	line := fmt.Sprintf("Size=%v, hits=%v, misses=%v\n", len(c.cache), c.hits, c.misses)
	for key, _ := range c.cache {
		line = fmt.Sprintf("%v\n  '%v'", line, key)
	}
	return line
}

func (c *Context) putCache(line string, position int, result int, actuals []int) {
	// c.cache[line] = true
	key := c.key(line, position, actuals)
	c.cache[key] = result
}

func (c *Context) key(line string, position int, actuals []int) string {
	return line
	// return fmt.Sprintf("%v_%v", position, actuals)
}

func (c *Context) clearCache() {
	c.cache = make(map[string]int)
}

func (c *Context) checkCache(current_line string, position int, actuals []int) (int, bool) {

	// key := current_line
	// _, exists := cache[key]
	// if exists {
	// 	return true
	// }
	// return false
	key := c.key(current_line, position, actuals)
	value, exists := c.cache[key]
	if exists {
		c.hits += 1
	} else {
		c.misses += 1
	}
	return value, exists

	// for index := 0; index < len(current_line); index++ {
	// 	key := current_line[0 : len(current_line)-index]
	// 	_, exists := c.cache[key]
	// 	// fmt.Printf("   %v  [key=%v] %v\n", current_line, key, exists)
	// 	if exists {
	// 		c.hits += 1
	// 		return true
	// 	}
	// }
	// c.misses += 1
	// return false
}
