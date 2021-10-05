package app

/*
--- Day 4: The Ideal Stocking Stuffer ---
Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for all the economically forward-thinking little girls and boys.

To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

For example:

If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.
If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....
Your puzzle input is bgvyzdsv.
*/

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// AOC_2015_03 is the entrypoint
func (app *Application) Y2015D04() {
	app.Y2015D04P1()
}

func (app *Application) Y2015D04P1() {
	prefix := "bgvyzdsv"
	// prefix = "abcdef"
	counter := 0

	// word := fmt.Sprintf("%v%v", prefix, 609043)
	// bytes := []byte(word)
	// hash := fmt.Sprintf("%x", md5.Sum(bytes))
	// fmt.Printf("%v %v : %v\n", prefix, 609043, hash)
	// fmt.Printf("\n\n")

	for {
		word := fmt.Sprintf("%v%v", prefix, counter)
		bytes := []byte(word)
		hash := fmt.Sprintf("%x", md5.Sum(bytes))
		if strings.Index(hash, "000000") == 0 {
			fmt.Printf("%v %v\n", hash, counter)
			break
		}
		counter++
	}
}
